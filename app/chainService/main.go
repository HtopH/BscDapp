package chainService

//import (
//	contractService "BscDapp/contract"
//	"context"
//	"fmt"
//	"github.com/ethereum/go-ethereum"
//	"github.com/ethereum/go-ethereum/accounts/abi"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/ethereum/go-ethereum/ethclient"
//	"math"
//	"math/big"
//	"strings"
//	time2 "time"
//)
//
//var (
//	conn         *contractService.BscGame
//	wsConn       *contractService.BscGame
//	rpcHttpUrl   = "https://speedy-nodes-nyc.moralis.io/783b783a3e310fa8f97290a5/bsc/testnet"  //合约http连接
//	rpcWsUrl     = "wss://speedy-nodes-nyc.moralis.io/783b783a3e310fa8f97290a5/bsc/testnet/ws" //合约ws连接
//	contractAddr = common.HexToAddress("0xbDeD695be08b81dDB0Ca3540F65d88A02562F147")           //合约地址
//	fromAddr     = common.HexToAddress("0x125a0daEE26BD73B37A3c2a86c84426c68743750")           //钱包账户地址
//	privateKey   = "841da76418e1314614ed7d88ba3f29067f5d532304c70499f331fb3aab9b7fd8"          //钱包账户
//	client       *ethclient.Client
//	wsClient     *ethclient.Client
//	err          error
//	timeFormat   = "2006-01-02 15:04:05"
//	thisQuery    ethereum.FilterQuery
//	ctx          context.Context //123123123
//	precision    float64         = 18
//)
//
//func init() {
//	ctx = context.Background()
//
//	client, err = ethclient.Dial(rpcHttpUrl)
//	if err != nil {
//		fmt.Println("连接网络失败", err)
//		return
//	}
//	defer client.Close()
//
//	conn, err = contractService.NewBscGame(contractAddr, client)
//	if err != nil {
//		fmt.Println("连接合约失败", err)
//		return
//	}
//
//	wsClient, err = ethclient.Dial(rpcWsUrl)
//	if err != nil {
//		fmt.Println("连接wss网络失败", err)
//		return
//	}
//	wsConn, err = contractService.NewBscGame(contractAddr, wsClient)
//	thisQuery = ethereum.FilterQuery{
//		Addresses: []common.Address{contractAddr},
//	}
//}
//
//func main() {
//	defer wsClient.Close()
//	defer wsClient.Close()
//	ListenJoin()
//}
//
//func GetPlayCount() {
//	res, err := conn.GetPlayerCount(nil)
//	if err != nil {
//		fmt.Println("调用合约失败", err)
//	} else {
//		fmt.Println(res)
//	}
//}
//
//func PlayInfo() {
//	id := big.NewInt(1)
//	res, err := conn.Players(nil, id)
//	if err != nil {
//		fmt.Println("调用合约失败", err)
//	} else {
//		fmt.Println(res)
//	}
//}
//
////读取合约日志（在指定的区块高度范围）（日志即合约所触发的事件）
//func readJoinBlockLog() {
//	query := ethereum.FilterQuery{
//		FromBlock: big.NewInt(14038265),
//		ToBlock:   big.NewInt(14038265),
//		Addresses: []common.Address{
//			contractAddr,
//		},
//	}
//	logs, err := client.FilterLogs(ctx, query)
//	if err != nil {
//		fmt.Println("查询日志失败", err)
//	}
//
//	contractAbi, err := abi.JSON(strings.NewReader(string(contractService.BscGameABI)))
//	if err != nil {
//		fmt.Println("解析日志abi错误", err)
//	}
//
//	for _, vLog := range logs {
//		res, fErr := contractAbi.Unpack("join", vLog.Data) //TODO 这里需要传入事件名称作为参数，存在一个问题，如果有多个事件，该如何解析日志
//		if fErr != nil {
//			fmt.Println("解析日志错误", fErr)
//			continue
//		}
//		//结构对应上了合约里的event事件输出，所以查询的应该是emit触法日志
//		fmt.Printf("解析日志结果:%+v,%T\n", res, res[0])
//
//	}
//
//}
//
//func ListenJoin() {
//	joinCh := make(chan *contractService.BscGameJoin)
//	joinSub, err := wsConn.WatchJoin(&bind.WatchOpts{}, joinCh, nil)
//	if err != nil {
//		fmt.Println("WatchJoin监听合约特定事件失败", err)
//		return
//	}
//
//	payLogCh := make(chan *contractService.BscGamePaylog)
//	payLogSub, err := wsConn.WatchPaylog(&bind.WatchOpts{}, payLogCh)
//	if err != nil {
//		fmt.Println("WatchPaylog监听合约特定事件失败", err)
//		return
//	}
//	now := time2.Now().Format(timeFormat)
//	fmt.Println(now + " 监听事件堵塞等待：")
//	for {
//		select {
//		case res := <-joinCh:
//			now = time2.Now().Format(timeFormat)
//			fmt.Println(now+" 监听返回一个结果：", res.Id, res.Addr) //该结果已解析
//		case err = <-joinSub.Err():
//			now = time2.Now().Format(timeFormat)
//			fmt.Println(now+" 监听事件结果错误", err)
//		case res := <-payLogCh:
//			now = time2.Now().Format(timeFormat)
//			fmt.Println(now+" 监听返回一个结果：", res.Id, res.Type, res.Value) //该结果已解析
//		case err = <-payLogSub.Err():
//			now = time2.Now().Format(timeFormat)
//			fmt.Println(now+" 监听事件结果错误", err)
//		}
//	}
//}
//
//func GetBalance()  {
//
//}
//
//func Pay() {
//	auth, err := getTransactOpts()
//	if err != nil {
//		return
//	}
//	res, err := conn.Pay(auth, getBigInt(0.5), 1)
//	if err != nil {
//		fmt.Println("调用合约失败", err)
//	} else {
//		fmt.Println(res)
//	}
//}
//
////组装写入合约的auth（需要消耗邮费）
//func getTransactOpts() (*bind.TransactOpts, error) {
//	nonce, err := client.PendingNonceAt(ctx, fromAddr)
//	if err != nil {
//		fmt.Println("获取nonce错误", err)
//		return nil, err
//	}
//	gasPrice, err := client.SuggestGasPrice(ctx)
//	if err != nil {
//		fmt.Println("Get gasPrice error", err)
//		return nil, err
//	}
//	priKey, err := crypto.HexToECDSA(privateKey)
//	auth := bind.NewKeyedTransactor(priKey)
//	auth.Nonce = big.NewInt(int64(nonce))
//	auth.Value = big.NewInt(0)
//	auth.GasLimit = uint64(300000)
//	auth.GasPrice = gasPrice
//	return auth, nil
//}
//
//func getBigInt(num float64) *big.Int {
//	res := num * math.Pow(10, precision)
//	return big.NewInt(int64(res))
//}
