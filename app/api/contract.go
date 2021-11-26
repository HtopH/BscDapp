package api

var ContractApi = contractApi{}

type contractApi struct{}

// @summary 注册
// @tags   合约接口
// @produce json
// @param   ref_addr formData string true "推荐人地址,无则传空字符串 "
// @router  /register [POST]
// @success 200 {object} service.JsonResponse "执行结果"
func (a *contractApi) Register() {}

// @summary 购买门票
// @tags   合约接口
// @produce json
// @param   _value formData int true "支付金额,big.int,需乘精度"
// @router  /payForTickets [POST]
// @success 200 {object} service.JsonResponse "执行结果"
func (a *contractApi) PayForTickets() {}

// @summary 投资
// @tags   合约接口
// @produce json
// @param   _value formData int true "支付金额,big.int,需乘精度"
// @router  /joinGame [POST]
// @success 200 {object} service.JsonResponse "执行结果"
func (a *contractApi) JoinGame() {}

// @summary 会员提现
// @tags   合约接口
// @produce json
// @param   _value formData int true "提现金额,big.int,需乘精度,不需用户填,直接传会员余额"
// @router  /userWithdraw [POST]
// @success 200 {object} service.JsonResponse "执行结果"
func (a *contractApi) UserWithdraw() {}
