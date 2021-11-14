package service

import (
	"BscDapp/app/chianService"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/frame/g"
	"math"
	"math/big"
)

var NewGame = newGame{}

type newGame struct {
	Conn         *chianService.BscGame
	WsConn       *chianService.BscGame
	Client       *ethclient.Client
	WsClient     *ethclient.Client
	TransactOpts *bind.TransactOpts
	Ctx          context.Context
}

var (
	rpcHttpUrl           = "https://data-seed-prebsc-1-s1.binance.org:8545/"                           //合约http连接
	rpcWsUrl             = "wss://speedy-nodes-nyc.moralis.io/783b783a3e310fa8f97290a5/bsc/testnet/ws" //合约ws连接
	contractAddr         = common.HexToAddress("0xbDeD695be08b81dDB0Ca3540F65d88A02562F147")           //合约地址
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

	s.Conn, err = chianService.NewBscGame(contractAddr, s.Client)
	if err != nil {
		g.Log().Debug("Service NewGame Init HttpNewBscGame Err :", err)
		return
	}

	s.WsClient, err = ethclient.Dial(rpcWsUrl)
	if err != nil {
		g.Log().Debug("Service NewGame Init WsClient Err :", err)
		return
	}
	s.WsConn, err = chianService.NewBscGame(contractAddr, s.WsClient)
	if err != nil {
		g.Log().Debug("Service NewGame Init WsNewBscGame Err :", err)
	}
}

func (s *newGame) GetPlayCount() {
	res, err := s.Conn.GetPlayerCount(nil)
	if err != nil {
		g.Log().Debug("Service NewGame Init WsNewBscGame Err :", err)
		return
	}
	g.Log().Debug("Service NewGame GetPlayCount res:", res)
}

func (s *newGame) Pay() {
	auth, err := s.GetTransactOpts()
	if err != nil {
		g.Log().Debug("Service NewGame GetTransactOpts Err :", err)
		return
	}
	res, err := s.Conn.Pay(auth, s.GetBigInt(0.5), 2)
	if err != nil {
		g.Log().Debug("Service NewGame Pay Err :", err)
		return
	}
	g.Log().Debug("Service NewGame Pay res:", res)

}

func (s *newGame) AdminWithdraw() {
	auth, err := s.GetTransactOpts()
	if err != nil {
		g.Log().Debug("Service NewGame GetTransactOpts Err :", err)
		return
	}
	res, err := s.Conn.WithdrawAdmin(auth, s.GetBigInt(0.4999999999999999), 1)
	if err != nil {
		g.Log().Debug("Service NewGame AdminWithdraw Err :", err)
		return
	}
	g.Log().Debug("Service NewGame AdminWithdraw res:", string(res.Data()))
}

func (s *newGame) ListenNewGame() {
	joinCh := make(chan *chianService.BscGameJoin)
	joinSub, err := s.WsConn.WatchJoin(&bind.WatchOpts{}, joinCh, nil)
	if err != nil {
		g.Log().Debug("WatchJoin监听合约特定事件失败", err)
		return
	}

	payLogCh := make(chan *chianService.BscGamePaylog)
	payLogSub, err := s.WsConn.WatchPaylog(&bind.WatchOpts{}, payLogCh)
	if err != nil {
		g.Log().Debug("WatchPayLog监听合约特定事件失败", err)
		return
	}
	g.Log().Debug("监听事件堵塞等待：")
	for {
		select {
		case res := <-joinCh:
			g.Log().Debug("joinCh监听返回一个结果：", res.Id, res.Addr) //该结果已解析
		case err = <-joinSub.Err():
			g.Log().Debug("joinSub监听事件结果错误", err)
		case res := <-payLogCh:
			g.Log().Debug("payLogCh监听返回一个结果：", res.Id, res.Type, res.Value) //该结果已解析
		case err = <-payLogSub.Err():
			g.Log().Debug("payLogSub监听事件结果错误", err)
		}
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
