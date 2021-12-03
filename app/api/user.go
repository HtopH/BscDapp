package api

import (
	"BscDapp/app/common"
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"BscDapp/app/service"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

var User = user{}

type user struct{}

// @summary 会员信息
// @tags    会员模块
// @produce json
// @router  /api/user/get-user-info  [POST]
// @success 200 {object} model.UserInfo "执行结果"
func (a *user) GetUserInfo(r *ghttp.Request) {
	var data = model.UserInfo{}
	userInfo := service.User.GetUser(r)
	data.FaBscUser = userInfo
	data.ReferReward, _ = dao.FaBscCredit.Where("uid=? and type=?", userInfo.Id, model.CreditRefReward).Sum(dao.FaBscCredit.Columns.Num)
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: data, Message: common.SuccessMsg})
}

// @summary 转账记录
// @tags    会员模块
// @produce json
// @param   page formData int false "页码"
// @param   size formData int false "每页数量"
// @router  /api/user/get-transfer-list  [POST]
// @success 200 {object} model.UserTransferList "执行结果"
func (a *user) GetTransferList(r *ghttp.Request) {
	var (
		req *model.UserInfoRep
	)
	if err := r.Parse(&req); err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	userInfo := service.User.GetUser(r)
	req.UserAddr = userInfo.Address
	data, err := service.User.TransferList(r.Context(), req)
	if err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: data, Message: common.SuccessMsg})
}

// @summary 兑换记录
// @tags    会员模块
// @produce json
// @param   page formData int false "页码"
// @param   size formData int false "每页数量"
// @router  /api/user/get-ticket-list  [POST]
// @success 200 {object} model.UserTickerList "执行结果"
func (a *user) GetTicketList(r *ghttp.Request) {
	var (
		req *model.UserInfoRep
	)
	if err := r.Parse(&req); err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	userInfo := service.User.GetUser(r)
	req.UserId = userInfo.Id
	data, err := service.User.TicketList(r.Context(), req)
	if err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: data, Message: common.SuccessMsg})
}

// @summary 一代直推会员
// @tags    会员模块
// @produce json
// @param   page formData int false "页码"
// @param   size formData int false "每页数量"
// @router  /api/user/get-first-team  [POST]
// @success 200 {object} model.UserTickerList "执行结果"
func (a *user) GetFirstTeam(r *ghttp.Request) {
	var (
		req *model.UserInfoRep
	)
	if err := r.Parse(&req); err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	userInfo := service.User.GetUser(r)
	req.UserId = userInfo.Id
	data, err := service.User.FirstLevelTeam(r.Context(), req)
	if err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: data, Message: common.SuccessMsg})
}

// @summary 校验参与活动资格(调用合约Join前调用校验)
// @tags    会员模块
// @produce json
// @router  /api/user/check-user-join  [GET]
// @success 200 {object} service.JsonResponse "执行结果"
func (a *user) CheckUserJoin(r *ghttp.Request) {
	userInfo := service.User.GetUser(r)
	count, err := dao.FaBscUserGame.Where("uid=? and status=1", userInfo.Id).Count()
	if err != nil || count > 0 {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: common.Failure})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Message: common.SuccessMsg})
}
