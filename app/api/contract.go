package api

var ContractApi = contractApi{}

type contractApi struct{}

// @summary PC网页支付
// @tags   合约接口
// @produce json
// @param   ref_addr formData string true "推荐人地址,无则传空字符串 "
// @param   orderNo formData string true "订单号"
// @router  register [POST]
// @success 200 {object} model.MpRes "执行结果"
