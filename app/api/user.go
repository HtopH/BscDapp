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
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: data, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), common.SuccessMsg)})
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
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), err.Error())})
	}
	userInfo := service.User.GetUser(r)
	req.UserAddr = userInfo.Address
	data, err := service.User.TransferList(r.Context(), req)
	if err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), err.Error())})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: data, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), common.SuccessMsg)})
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
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), err.Error())})
	}
	userInfo := service.User.GetUser(r)
	req.UserId = userInfo.Id
	data, err := service.User.TicketList(r.Context(), req)
	if err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), err.Error())})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: data, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), common.SuccessMsg)})
}

// @summary 一代直推会员
// @tags    会员模块
// @produce json
// @param   page formData int false "页码"
// @param   size formData int false "每页数量"
// @router  /api/user/get-first-team  [POST]
// @success 200 {object} model.TeamUserList "执行结果"
func (a *user) GetFirstTeam(r *ghttp.Request) {
	var (
		req *model.UserInfoRep
	)
	if err := r.Parse(&req); err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), err.Error())})
	}
	userInfo := service.User.GetUser(r)
	req.UserId = userInfo.Id
	data, err := service.User.FirstLevelTeam(r.Context(), req)
	if err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), err.Error())})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: data, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), common.SuccessMsg)})
}

// @summary 校验参与活动资格(调用合约Join前调用校验)
// @tags    会员模块
// @produce json
// @router  /api/user/check-user-join  [GET]
// @success 200 {object} service.JsonResponse "执行结果"
func (a *user) CheckUserJoin(r *ghttp.Request) {
	userInfo := service.User.GetUser(r)
	gameInfo, err := dao.FaBscUserGame.Where("uid=? and status<3", userInfo.Id).One()
	if err != nil || gameInfo != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), common.Failure)})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), common.SuccessMsg)})
}

// @summary 会员活动信息
// @tags    会员模块
// @produce json
// @router  /api/user/user-game-info  [GET]
// @success 200 {object} model.UserGameInfo "执行结果"
func (a *user) UserGameInfo(r *ghttp.Request) {
	userInfo := service.User.GetUser(r)
	data, err := service.User.GetUserGameInfo(r.Context(), userInfo.Id)
	if err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), err.Error())})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: data, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), common.SuccessMsg)})
}

// @summary 获得投资收益(转卖)
// @tags    会员模块
// @produce json
// @router  /api/user/get-game-reward  [GET]
// @success 200 {object} service.JsonResponse "执行结果"
func (a *user) GetGameReward(r *ghttp.Request) {
	userInfo := service.User.GetUser(r)
	err := service.User.GameReward(r.Context(), userInfo.Id)
	if err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), err.Error())})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), common.SuccessMsg)})
}

// @summary 获得推荐奖统计
// @tags    会员模块
// @produce json
// @router  /api/user/get-ref-reward-info  [GET]
// @success 200 {object} model.UserRefRewardInfo "执行结果"
func (a *user) GetRefRewardInfo(r *ghttp.Request) {
	userInfo := service.User.GetUser(r)
	data, err := service.User.UserRefReward(userInfo.Id, 5)
	if err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), err.Error())})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: data, Message: service.Multilingual(r.GetCtxVar(common.LanguageKey).String(), common.SuccessMsg)})
}
