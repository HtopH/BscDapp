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

func (s *newGame) Init() {
	g.Log().Debug("newGame Init Start")
	var err error
	s.Ctx = context.Background()
	s.Client, err = ethclient.Dial(model.RpcHttpUrl)
	if err != nil {
		g.Log().Debug("Service NewGame Init HttpClient Err :", err)
		fmt.Println("连接网络失败", err)
		return
	}

	s.Conn, err = chainService.NewBscGame(model.ContractAddr, s.Client)
	if err != nil {
		g.Log().Debug("Service NewGame Init HttpNewBscGame Err :", err)
		return
	}

	s.WsClient, err = ethclient.Dial(model.RpcWsUrl)
	if err != nil {
		g.Log().Debug("Service NewGame Init WsClient Err :", err)
		return
	}
	s.WsConn, err = chainService.NewBscGame(model.ContractAddr, s.WsClient)
	if err != nil {
		g.Log().Debug("Service NewGame Init WsNewBscGame Err :", err)
	}
}

//系统支付给用户
func (s *newGame) Pay(param *model.TaskAddUserBalance) (string, error) {
	auth, err := s.GetTransactOpts()
	if err != nil {
		return "", err
	}
	res, err := s.Conn.Pay(auth, param.UserId, GetBigInt(param.Value, model.TokenDecimals))
	if err != nil {
		g.Log().Debug("Service NewGame Pay Err :", err)
		return "", err
	}
	return res.Hash().String(), nil
}

//更新当前活动场次
func (s *newGame) SetRound() (string, error) {
	auth, err := s.GetTransactOpts()
	if err != nil {
		return "", err
	}
	res, err := s.Conn.SetRound(auth)
	if err != nil {
		g.Log().Debug("Service NewGame SetRound Err :", err)
		return "", err
	}
	return res.Hash().String(), nil
}

//会员出局
//func (s *newGame) UserOut(param *model.TaskUserOut) (string, error) {
//	auth, err := s.GetTransactOpts()
//	if err != nil {
//		return "", err
//	}
//	res, err := s.Conn.UserOut(auth, param.UserId, param.Round)
//	if err != nil {
//		g.Log().Debug("Service NewGame UserOut Err :", err)
//		return "", err
//	}
//	return res.Hash().String(), nil
//}

//获取兑换比例
func (s *newGame) GetPercent() (uint32, error) {
	res, err := s.Conn.GetPercent(nil)
	if err != nil {
		g.Log().Debug("Service NewGame GetPercent Err:", err)
	}
	return res, err
}

//通过ID获取地址
func (s *newGame) IdToAddr(id uint64) common.Address {
	res, err := s.Conn.IdToAddr(nil, id)
	if err != nil {
		g.Log().Debug("Service NewGame IdToAddr Err :", err)
	}
	return res
}

//设置粮草消耗值
func (s *newGame) SetSpend(num float64) (string, error) {
	auth, err := s.GetTransactOpts()
	if err != nil {
		return "", err
	}
	res, err := s.Conn.SetSpend(auth, GetBigInt(num, model.TokenDecimals))
	if err != nil {
		g.Log().Debug("Service NewGame SetSpend Err :", err)
		return "", err
	}
	return res.Hash().String(), nil
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
	buyTicketSub, err := s.WsConn.WatchBuyTicketLog(&bind.WatchOpts{}, buyTicketCh)
	if err != nil {
		g.Log().Debug("WatchBuyTicketLog监听合约特定事件失败", err)
		return
	}
	//参与活动
	joinCh := make(chan *chainService.BscGameJoinLog)
	joinSub, err := s.WsConn.WatchJoinLog(&bind.WatchOpts{}, joinCh)
	if err != nil {
		g.Log().Debug("WatchJoinLog监听合约特定事件失败", err)
		return
	}
	//会员提现
	userGetCh := make(chan *chainService.BscGameUserGetLog)
	userGetSub, err := s.WsConn.WatchUserGetLog(&bind.WatchOpts{}, userGetCh)
	if err != nil {
		g.Log().Debug("WatchUserGetLog监听合约特定事件失败", err)
		return
	}
	//门票转账
	transferCh := make(chan *chainService.BscGameTransferLog)
	transferSub, err := s.WsConn.WatchTransferLog(&bind.WatchOpts{}, transferCh)
	if err != nil {
		g.Log().Debug("WatchTransferLog监听合约特定事件失败", err)
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
			data.Uid = int(res.Id)
			data.Block = int64(res.Raw.BlockNumber)
			data.TxHash = res.Raw.TxHash.String()
			data.Data = gconv.String(res)

		case err = <-registerSub.Err():
			g.Log().Debug("joinSub监听事件结果错误", err)
			run = false
			break
		case res := <-buyTicketCh:
			g.Log().Debug("buyTicketCh监听返回一个结果：", res) //该结果已解析
			err = ListenTask.DealBuyTicket(context.Background(), res)
			if err != nil {
				data.Status = 2
				data.Remark = err.Error()
			} else {
				data.Status = 1
			}
			//插入监听记录
			data.Type = model.ListenBuy
			data.Uid = int(res.Id)
			data.Data = gconv.String(res)
			data.Block = int64(res.Raw.BlockNumber)
			data.TxHash = res.Raw.TxHash.String()

		case err = <-buyTicketSub.Err():
			g.Log().Debug("buyTicketSub监听事件结果错误", err)
			run = false
			break
		case res := <-joinCh:
			g.Log().Debug("joinCh监听返回一个结果：", res) //该结果已解析
			err = ListenTask.DealUserJoinGame(context.Background(), res)
			if err != nil {
				data.Status = 2
				data.Remark = err.Error()
			} else {
				data.Status = 1
			}
			//插入监听记录
			data.Type = model.ListenJoin
			data.Uid = int(res.Id)
			data.Data = gconv.String(res)
			data.Block = int64(res.Raw.BlockNumber)
			data.TxHash = res.Raw.TxHash.String()

		case err = <-joinSub.Err():
			g.Log().Debug("joinSub监听事件结果错误", err)
			run = false
			break
		case res := <-userGetCh:
			g.Log().Debug("userGetCh监听返回一个结果：", res) //该结果已解析
			err = ListenTask.DealUserWithdraw(context.Background(), res)
			if err != nil {
				data.Status = 2
				data.Remark = err.Error()
			} else {
				data.Status = 1
			}
			data.Type = model.ListenWithdrawal
			data.Uid = int(res.Id)
			data.Data = gconv.String(res)
			data.Block = int64(res.Raw.BlockNumber)
			data.TxHash = res.Raw.TxHash.String()

		case err = <-userGetSub.Err():
			g.Log().Debug("userGetSub监听事件结果错误", err)
			run = false
			break
		case res := <-transferCh:
			g.Log().Debug("transferCh监听返回一个结果：", res) //该结果已解析
			err = ListenTask.DealTransfer(context.Background(), res)
			if err != nil {
				data.Status = 2
				data.Remark = err.Error()
			} else {
				data.Status = 1
			}
			data.Type = model.ListenTransfer
			data.Data = gconv.String(res)
			data.Block = int64(res.Raw.BlockNumber)
			data.TxHash = res.Raw.TxHash.String()

		case err = <-transferSub.Err():
			g.Log().Debug("transferSub监听事件结果错误", err)
			run = false
			break
		}
		data.Created = int(time.Now().Unix())
		_, err = dao.FaBscListenLog.OmitEmpty().Save(data)
		if err != nil {
			g.Log().Debug("ListenNewGame ListenLog Save Err:", err)
		}
		//更新系统区块高度
		_, err = dao.FaBscBaseInfo.Where("theKey=?", model.BaseReadKey).Update(g.Map{"theValue": data.Block + 1})
		if err != nil {
			g.Log().Debug("ListenNewGame BaseInfo Update Err:", err)
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
			model.ContractAddr,
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
		count, err := dao.FaBscListenLog.Where("tx_hash=?", vLog.TxHash.String()).Count()
		//没有记录,说明没有监听到,需要补足
		if count > 0 {
			continue
		}
		logData := model.FaBscListenLog{
			TxHash:  vLog.TxHash.String(),
			Block:   int64(vLog.BlockNumber),
			LogType: 2,
			Created: int(time.Now().Unix()),
			Status:  0,
		}
		doType := struct {
			DoType uint8
		}{}
		err = contractAbi.UnpackIntoInterface(&doType, "registerLog", vLog.Data)
		g.Log().Debug("doType:", doType)
		switch doType.DoType {
		case model.DoTypeRegister:
			//TODO 尝试解析注册事件
			regEvent := struct {
				DoType uint8
				Id     uint64
				Addr   common.Address
				RefId  uint64
			}{}
			err = contractAbi.UnpackIntoInterface(&regEvent, "registerLog", vLog.Data)
			g.Log().Debug("registerLog 解析日志:", regEvent)
			//不报错说明是注册事件
			if err == nil && regEvent.DoType == 1 {
				//日志解析不到地址,直接通过ID读取会员地址
				regEvent.Addr = s.IdToAddr(regEvent.Id)
				regData := &chainService.BscGameRegisterLog{
					DoType: regEvent.DoType,
					Id:     regEvent.Id,
					Addr:   regEvent.Addr,
					RefId:  regEvent.RefId,
				}

				err = ListenTask.DealRegister(context.Background(), regData)
				if err != nil {
					logData.Status = 2
					logData.Remark = err.Error()
				} else {
					logData.Status = 1
				}
				logData.Uid = int(regEvent.Id)
				logData.Data = gconv.String(regData)
				logData.Type = model.ListenRegister
			}
		case model.DoTypeBuy:
			//TODO 尝试解析购买门票事件
			buyEvent := struct {
				DoType    uint8
				Id        uint64
				Value     *big.Int
				GetTicket *big.Int
				Percent   uint32
			}{}
			err = contractAbi.UnpackIntoInterface(&buyEvent, "buyTicketLog", vLog.Data)
			g.Log().Debug("buyTicketLog 解析日志:", buyEvent)
			//不报错说明是购买门票事件
			if err == nil && buyEvent.DoType == 2 {
				buyData := &chainService.BscGameBuyTicketLog{
					DoType:    buyEvent.DoType,
					Id:        buyEvent.Id,
					Value:     buyEvent.Value,
					GetTicket: buyEvent.GetTicket,
					Percent:   buyEvent.Percent,
				}
				logData.Data = gconv.String(buyData)
				err = ListenTask.DealBuyTicket(context.Background(), buyData)
				if err != nil {
					logData.Status = 2
					logData.Remark = err.Error()
				} else {
					logData.Status = 1
				}
				logData.Uid = int(buyEvent.Id)
				logData.Type = model.ListenBuy

			}
		case model.DoTypeJoin:
			//TODO 尝试解析参与活动事件
			joinEvent := struct {
				DoType uint8
				Id     uint64
				Value  *big.Int
				Round  uint32
			}{}
			err = contractAbi.UnpackIntoInterface(&joinEvent, "joinLog", vLog.Data)
			g.Log().Debug("joinLog 解析日志:", joinEvent)
			if err == nil && joinEvent.DoType == 3 {
				joinData := &chainService.BscGameJoinLog{
					DoType: joinEvent.DoType,
					Id:     joinEvent.Id,
					Value:  joinEvent.Value,
					Round:  joinEvent.Round,
				}
				logData.Data = gconv.String(joinData)
				err = ListenTask.DealUserJoinGame(context.Background(), joinData)
				if err != nil {
					logData.Status = 2
					logData.Remark = err.Error()
				} else {
					logData.Status = 1
				}
				logData.Uid = int(joinEvent.Id)
				logData.Type = model.ListenJoin
			}
		case model.DoTypeWithdrawal:
			//TODO 尝试解析提现请求
			getEvent := struct {
				DoType     uint8
				RewardType uint8
				Id         uint64
				Value      *big.Int
			}{}
			err = contractAbi.UnpackIntoInterface(&getEvent, "userGetLog", vLog.Data)
			g.Log().Debug("userGetLog 解析日志:", getEvent)
			if err == nil && getEvent.DoType == 4 {
				getData := &chainService.BscGameUserGetLog{
					DoType:     getEvent.DoType,
					RewardType: getEvent.RewardType,
					Id:         getEvent.Id,
					Value:      getEvent.Value,
				}
				logData.Data = gconv.String(getData)
				err = ListenTask.DealUserWithdraw(context.Background(), getData)
				if err != nil {
					logData.Status = 2
					logData.Remark = err.Error()
				} else {
					logData.Status = 1
				}
				logData.Uid = int(getEvent.Id)
				logData.Type = model.ListenWithdrawal
			}
		case model.DoTypeTransfer:
			//TODO 尝试解析转账请求
			transferEvent := struct {
				DoType   uint8
				FromAddr common.Address
				ToAddr   common.Address
				Value    *big.Int
			}{}
			err = contractAbi.UnpackIntoInterface(&transferEvent, "transferLog", vLog.Data)
			g.Log().Debug("transferLog 解析日志:", transferEvent)
			if err == nil && transferEvent.DoType == 5 {
				getData := &chainService.BscGameTransferLog{
					DoType:   transferEvent.DoType,
					FromAddr: transferEvent.FromAddr,
					ToAddr:   transferEvent.ToAddr,
					Value:    transferEvent.Value,
				}
				logData.Data = gconv.String(getData)
				err = ListenTask.DealTransfer(context.Background(), getData)
				if err != nil {
					logData.Status = 2
					logData.Remark = err.Error()
				} else {
					logData.Status = 1
				}
				logData.Type = model.ListenTransfer
			}
		}

		//将获取到的交易日志存储起来
		if err != nil {
			logData.Status = 2
			logData.Remark = err.Error()
		}
		_, err = dao.FaBscListenLog.OmitEmpty().Save(logData)
		if err != nil {
			g.Log().Debug("ReadBlockLog ListenLog Save err:", err)
		}
	}
	//更新拉取区块高度起始点
	_, _ = dao.FaBscBaseInfo.Where("theKey=?", model.BaseReadKey).Update(g.Map{"theValue": endBlock + 1, "updated": time.Now().Unix()})
	//结束的区块还不是最新的，继续拉取
	if maxBlockNum > uint64(endBlock+1) {
		s.ReadBlockLog(endBlock + 1)
	}
}

//组装写入合约的auth（需要消耗邮费）
func (s *newGame) GetTransactOpts() (*bind.TransactOpts, error) {
	nonce, err := s.Client.PendingNonceAt(s.Ctx, model.FromAddr)
	if err != nil {
		g.Log().Debug("GetTransactOpts 获取nonce错误", err)
		return nil, err
	}
	gasPrice, err := s.Client.SuggestGasPrice(s.Ctx)
	if err != nil {
		g.Log().Debug("GetTransactOpts Get gasPrice error", err)
		return nil, err
	}
	priKey, err := crypto.HexToECDSA(model.PrivateKey)
	auth := bind.NewKeyedTransactor(priKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth, nil
}
