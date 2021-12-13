package api

var ContractApi = contractApi{}

type contractApi struct{}

// @summary 注册（先调中心化接口校验会员是否存在）
// @tags   合约接口
// @produce json
// @param   ref_addr formData string true "推荐人地址,无则传系统根节点地址"
// @router  /register [POST]
// @success 200 {object} service.JsonResponse "执行结果"
func (a *contractApi) Register() {}

// @summary 购买门票(支付USDT)
// @tags   合约接口
// @produce json
// @param   _value formData int true "支付金额,big.int,需乘精度"
// @router  /payForTickets [POST]
// @success 200 {object} service.JsonResponse "执行结果"
func (a *contractApi) PayForTickets() {}

// @summary 投资（先调中心化接口校验投资资格）(支付USDT和粮草)
// @tags   合约接口
// @produce json
// @param   _value formData int true "支付金额,big.int,需乘精度"
// @router  /joinGame [POST]
// @success 200 {object} service.JsonResponse "执行结果"
func (a *contractApi) JoinGame() {}

// @summary 会员提现（先调中心化会员数据接口校验提现额度）
// @tags   合约接口
// @produce json
// @param   _type formData int true "提现类型:1-投资回报,2-奖池奖励,3-推荐奖励"
// @param   _value formData int true "提现金额,big.int,需乘精度,不需用户填,直接传会员余额"
// @router  /userWithdraw [POST]
// @success 200 {object} service.JsonResponse "执行结果"
func (a *contractApi) UserWithdraw() {}

// @summary 会员粮草转账(支付粮草)
// @tags   合约接口
// @produce json
// @param   _addr formData string true "接收人钱包地址"
// @param   _value formData int true "转账金额,big.int,需乘精度"
// @router  /transferToUser [POST]
// @success 200 {object} service.JsonResponse "执行结果"
func (a *contractApi) transferToUser() {}
