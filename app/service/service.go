package service

type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码
	Data    interface{} `json:"data"`    // 返回数据
	Message string      `json:"message"` // 提示信息
}
