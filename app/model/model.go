package model

const (
	//监听类型
	ListenRegister   = "register"
	ListenBuy        = "buyTicket"
	ListenJoin       = "join"
	ListenWithdrawal = "withdrawal"

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
	CreditReward    = "returnReward"   //投资回报
	CreditPool      = "poolReward"     //奖池奖励
	CreditRefReward = "referrerReward" //推荐奖励
	CreditBuyTicket = "buyTicket"      //兑换门票
	CreditWithdraw  = "withdraw"       //兑换门票

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
