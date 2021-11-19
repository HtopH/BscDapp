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
	TokenDecimals   = "1000000000000000000"
	PercentDecimals = "1000000"

	//发送交易事件
	SendAddBalance = "addBalance"

	//余额变动事件
	CreditReward = "reward" //投资回报
	CreditPool   = "pool"   //奖池奖励
)
