package model

import "github.com/ethereum/go-ethereum/common"

//列表输出字段
type ListOutPut struct {
	Page  int
	Size  int
	Total int
}

//页码
type PageReq struct {
	Page int `d:"1"`  //页码
	Size int `d:"10"` //每页条数
}

var (
	//合约地址
	ContractAddr = common.HexToAddress("0x5625Cf2D8e8E0492A5B190B897973ebC8e57a7c1")
	//操作员地址
	FromAddr = common.HexToAddress("0x125a0daEE26BD73B37A3c2a86c84426c68743750")
)

//币安链DApp基础信息
const (
	//根节点地址
	OwnAddr = "0x125a0daEE26BD73B37A3c2a86c84426c68743750"
	//币安http连接
	RpcHttpUrl = "https://data-seed-prebsc-1-s1.binance.org:8545/"
	//币安ws连接
	RpcWsUrl = "wss://speedy-nodes-nyc.moralis.io/783b783a3e310fa8f97290a5/bsc/testnet/ws"
	//钱包账户私钥
	PrivateKey = "841da76418e1314614ed7d88ba3f29067f5d532304c70499f331fb3aab9b7fd8"
)

const (
	//监听类型
	ListenRegister   = "register"
	ListenBuy        = "buyTicket"
	ListenJoin       = "join"
	ListenWithdrawal = "withdrawal"
	ListenTransfer   = "transfer"

	DoTypeRegister   = 1
	DoTypeBuy        = 2
	DoTypeJoin       = 3
	DoTypeWithdrawal = 4
	DoTypeTransfer   = 5

	//系统基础信息key
	BaseReadKey  = "readBlock"
	BaseSpendKey = "spendTicket"

	//精度
	TokenDecimals   = "1000000000000000000" //代币精度
	PercentDecimals = "1000000"             //比例精度

	//发送交易事件
	SendPay      = "pay"
	SendSetRound = "setRound"
	SendUserOut  = "userOut"

	//余额变动事件
	CreditReward       = "returnReward"   //投资回报
	CreditPool         = "poolReward"     //奖池奖励
	CreditRefReward    = "referrerReward" //推荐奖励
	CreditBuyTicket    = "buyTicket"      //兑换门票
	CreditWithdraw     = "withdraw"       //提现投资回报
	CreditPoolWithdraw = "poolWithdraw"   //提现奖池奖励
	CreditRefWithdraw  = "refWithdraw"    //提现奖池奖励

	//提现类型
	WithdrawRefType    = 3 //推荐奖提现
	WithdrawPoolType   = 2 //奖池提现
	WithdrawCreditType = 1 //投资回报提现

	//余额字段类型
	UserCreditOne   = "credit"    //投资回报
	UserCreditTwo   = "ticket"    //门票
	UserCreditThree = "pool"      //奖池奖励
	UserCreditFour  = "refCredit" //推荐奖

	//数值比例设置
	PercentBase       = 100 //基础数值
	PercentJoinTicket = 10  //投资消耗门票比例
	PercentJoinOut    = 70  //投资给出比例
	PercentJoinReturn = 140 //投资回报比例
	PercentJoinToJack = 15  //投资抽取到奖池比例
	PercentJoinToSeed = 15  //投资抽取到种子池比例

	//奖池金额与对应的倒计时
	GameLevelOneTime   = 60 * 60 * 24
	GameLevelTwoPool   = 100000
	GameLevelTwoTime   = 60 * 60 * 12
	GameLevelThreePool = 200000
	GameLevelThreeTime = 60 * 60 * 6
	GameLevelFourPool  = 300000
	GameLevelFourTime  = 60 * 60 * 3
	GameLevelFivePool  = 400000
	GameLevelFiveTime  = 60 * 90
	GameBaseTime       = 60 * 10

	TaskLateTime = 60 * 15 //任务延迟等待时间
)
