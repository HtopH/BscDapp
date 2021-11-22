package service

import (
	"BscDapp/app/chainService"
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"math"
	"math/big"
	"strings"
	"time"
)

var NewGame = newGame{}

type newGame struct {
	Conn         *chainService.BscGame
	WsConn       *chainService.BscGame
	Client       *ethclient.Client
	WsClient     *ethclient.Client
	TransactOpts *bind.TransactOpts
	Ctx          context.Context
}

var (
	rpcHttpUrl           = "https://data-seed-prebsc-1-s1.binance.org:8545/"                           //合约http连接
	rpcWsUrl             = "wss://speedy-nodes-nyc.moralis.io/783b783a3e310fa8f97290a5/bsc/testnet/ws" //合约ws连接
	contractAddr         = common.HexToAddress("0x10aaa4112Cd0B4e27941e3beE286773E784fE317")           //合约地址
	fromAddr             = common.HexToAddress("0x125a0daEE26BD73B37A3c2a86c84426c68743750")           //钱包账户地址
	privateKey           = "841da76418e1314614ed7d88ba3f29067f5d532304c70499f331fb3aab9b7fd8"          //钱包账户
	precision    float64 = 18
)

func (s *newGame) Init() {
	var err error
	s.Ctx = context.Background()
	s.Client, err = ethclient.Dial(rpcHttpUrl)
	if err != nil {
		g.Log().Debug("Service NewGame Init HttpClient Err :", err)
		fmt.Println("连接网络失败", err)
		return
	}

	s.Conn, err = chainService.NewBscGame(contractAddr, s.Client)
	if err != nil {
		g.Log().Debug("Service NewGame Init HttpNewBscGame Err :", err)
		return
	}

	s.WsClient, err = ethclient.Dial(rpcWsUrl)
	if err != nil {
		g.Log().Debug("Service NewGame Init WsClient Err :", err)
		return
	}
	s.WsConn, err = chainService.NewBscGame(contractAddr, s.WsClient)
	if err != nil {
		g.Log().Debug("Service NewGame Init WsNewBscGame Err :", err)
	}
}

//func (s *newGame) Pay() {
//	auth, err := s.GetTransactOpts()
//	if err != nil {
//		g.Log().Debug("Service NewGame GetTransactOpts Err :", err)
//		return
//	}
//	res, err := s.Conn.Pay(auth, s.GetBigInt(0.5), 2)
//	if err != nil {
//		g.Log().Debug("Service NewGame Pay Err :", err)
//		return
//	}
//	g.Log().Debug("Service NewGame Pay res:", res)
//}

//通过ID获取地址
func (s *newGame) IdToAddr(id uint64) common.Address {
	res, err := s.Conn.IdToAddr(nil, id)
	if err != nil {
		g.Log().Debug("Service NewGame Init WsNewBscGame Err :", err)
	}
	return res
}

//监听合约事件
func (s *newGame) ListenNewGame() {
	//注册
	registerCh := make(chan *chainService.BscGameRegisterLog)
	registerSub, err := s.WsConn.WatchRegisterLog(&bind.WatchOpts{}, registerCh, nil)
	if err != nil {
		g.Log().Debug("WatchRegisterLog监听合约特定事件失败", err)
		return
	}
	//购买门票
	buyTicketCh := make(chan *chainService.BscGameBuyTicketLog)
	buyTicketSub, err := s.WsConn.WatchBuyTicketLog(&bind.WatchOpts{}, buyTicketCh, nil)
	if err != nil {
		g.Log().Debug("WatchBuyTicketLog监听合约特定事件失败", err)
		return
	}
	//投资
	joinCh := make(chan *chainService.BscGameJoinLog)
	joinSub, err := s.WsConn.WatchJoinLog(&bind.WatchOpts{}, joinCh, nil)
	if err != nil {
		g.Log().Debug("WatchJoinLog监听合约特定事件失败", err)
		return
	}
	//提现
	userGetCh := make(chan *chainService.BscGameUserGetLog)
	userGetSub, err := s.WsConn.WatchUserGetLog(&bind.WatchOpts{}, userGetCh, nil)
	if err != nil {
		g.Log().Debug("WatchUserGetLog监听合约特定事件失败", err)
		return
	}
	g.Log().Debug("监听事件堵塞等待：")
	run := true
	for run {
		data := model.FaBscListenLog{}
		select {
		case res := <-registerCh:
			g.Log().Debug("registerCh监听返回一个结果：", res) //该结果已解析
			err = ListenTask.DealRegister(context.Background(), res)
			if err != nil {
				data.Status = 2
				data.Remark = err.Error()
			} else {
				data.Status = 1
			}
			//插入监听记录
			data.Type = model.ListenRegister
			data.Block = int64(res.Raw.BlockNumber)
			data.TxHash = res.Raw.TxHash.String()
			data.Data = gconv.String(res)
			data.Created = int(time.Now().Unix())

			_, err = dao.FaBscListenLog.OmitEmpty().Save(data)
			if err != nil {
				g.Log().Debug("ListenNewGame Err:", err)
			}
		case err = <-registerSub.Err():
			g.Log().Debug("joinSub监听事件结果错误", err)
			run = false
			break
		case res := <-buyTicketCh:
			g.Log().Debug("buyTicketCh监听返回一个结果：", res) //该结果已解析
			err = ListenTask.DelBuyTicket(context.Background(), res)
			if err != nil {
				data.Status = 2
				data.Remark = err.Error()
			} else {
				data.Status = 1
			}
			//插入监听记录
			data.Type = model.ListenBuy
			data.Data = gconv.String(res)
			data.TxHash = res.Raw.TxHash.String()
			data.Data = gconv.String(res)
			data.Created = int(time.Now().Unix())
			_, err = dao.FaBscListenLog.OmitEmpty().Save(data)
			if err != nil {
				g.Log().Debug("ListenNewGame Err:", err)
			}
		case err = <-buyTicketSub.Err():
			g.Log().Debug("buyTicketSub监听事件结果错误", err)
			run = false
			break
		case res := <-joinCh:
			g.Log().Debug("joinCh监听返回一个结果：", res) //该结果已解析
			data.Type = model.ListenJoin
			data.Data = gconv.String(res)
		case err = <-joinSub.Err():
			g.Log().Debug("joinSub监听事件结果错误", err)
			run = false
			break
		case res := <-userGetCh:
			g.Log().Debug("userGetCh监听返回一个结果：", res) //该结果已解析
			data.Type = model.ListenWithdrawal
			data.Data = gconv.String(res)
		case err = <-userGetSub.Err():
			g.Log().Debug("userGetSub监听事件结果错误", err)
			run = false
			break
		}

	}
}

//读取合约日志（在指定的区块高度范围）（日志即合约所触发的事件）
func (s *newGame) ReadBlockLog(startBlock int64) {
	maxBlockNum, _ := s.Client.BlockNumber(s.Ctx)
	if uint64(startBlock) > maxBlockNum {
		return
	}
	endBlock := startBlock + 5000
	if endBlock > int64(maxBlockNum) {
		endBlock = int64(maxBlockNum)
	}
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(startBlock),
		ToBlock:   big.NewInt(endBlock),
		Addresses: []common.Address{
			contractAddr,
		},
	}
	logs, err := s.Client.FilterLogs(s.Ctx, query)
	if err != nil {
		g.Log().Debug("查询日志失败", err)
		return
	}

	contractAbi, err := abi.JSON(strings.NewReader(chainService.BscGameABI))
	if err != nil {
		g.Log().Debug("解析合约abi错误", err)
		return
	}

	for _, vLog := range logs {
		g.Log().Debug("vLog TxHash:", vLog.TxHash.String())
		count, err := dao.FaBscListenLog.Where("tx_hash=?", vLog.TxHash.String()).Count()
		//没有记录,说明没有监听到,需要补足
		if count > 0 {
			continue
		}
		logData := model.FaBscListenLog{
			Data:    gconv.String(vLog),
			TxHash:  vLog.TxHash.String(),
			Block:   int64(vLog.BlockNumber),
			LogType: 2,
			Created: int(time.Now().Unix()),
			Status:  0,
		}

		//尝试解析注册事件
		regEvent := struct {
			Id    uint64
			Addr  common.Address
			RefId uint64
		}{}
		err = contractAbi.UnpackIntoInterface(&regEvent, "registerLog", vLog.Data)
		//不报错说明是注册事件
		if err == nil {
			//日志解析不到地址,直接通过ID读取会员地址
			regEvent.Addr = s.IdToAddr(regEvent.Id)
			g.Log().Debug("registerLog 读取未被监听的日志:", regEvent)
			regData := &chainService.BscGameRegisterLog{
				Id:    regEvent.Id,
				Addr:  regEvent.Addr,
				RefId: regEvent.RefId,
			}
			err = ListenTask.DealRegister(context.Background(), regData)
			if err != nil {
				logData.Status = 2
				logData.Remark = err.Error()
			}
			logData.Type = model.ListenRegister
			continue
		}

		//尝试解析购买门票事件
		buyEvent := struct {
			Addr    common.Address
			Fee     big.Int
			Num     big.Int
			Percent uint32
		}{}
		err = contractAbi.UnpackIntoInterface(buyEvent, "buyTicketLog", vLog.Data)
		//不报错说明是购买门票事件
		if err == nil {
			//日志解析不到地址,直接通过ID读取会员地址
			regEvent.Addr = s.IdToAddr(regEvent.Id)
			g.Log().Debug("registerLog 读取未被监听的日志:", regEvent)
			regData := &chainService.BscGameRegisterLog{
				Id:    regEvent.Id,
				Addr:  regEvent.Addr,
				RefId: regEvent.RefId,
			}
			err = ListenTask.DealRegister(context.Background(), regData)
			if err != nil {
				logData.Status = 2
				logData.Remark = err.Error()
			}
			logData.Type = model.ListenRegister
			continue
		}

		//res, err = contractAbi.Unpack("userGetLog", vLog.Data)
		//不报错说明是购买门票事件
		if err == nil {
			//g.Log().Debug("registerLog 日志:", res)
			continue
		}

		//将获取到的交易日志存储起来
		_, err = dao.FaBscListenLog.OmitEmpty().Save(logData)
		if err != nil {
			g.Log().Debug("ReadBlockLog registerLog save err:", err)
		}
	}
	//更新拉取区块高度起始点
	dao.FaBscBaseInfo.Where("theKey=?", model.BaseReadKey).Update(g.Map{"theValue": endBlock + 1})
	//结束的区块还不是最新的，继续拉取
	if maxBlockNum > uint64(endBlock+1) {
		s.ReadBlockLog(endBlock + 1)
	}
}

//组装写入合约的auth（需要消耗邮费）
func (s *newGame) GetTransactOpts() (*bind.TransactOpts, error) {
	nonce, err := s.Client.PendingNonceAt(s.Ctx, fromAddr)
	if err != nil {
		g.Log().Debug("获取nonce错误", err)
		return nil, err
	}
	gasPrice, err := s.Client.SuggestGasPrice(s.Ctx)
	if err != nil {
		g.Log().Debug("Get gasPrice error", err)
		return nil, err
	}
	priKey, err := crypto.HexToECDSA(privateKey)
	auth := bind.NewKeyedTransactor(priKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth, nil
}

func (s *newGame) GetBigInt(num float64) *big.Int {
	res := num * math.Pow(10, precision)
	return big.NewInt(int64(res))
}
