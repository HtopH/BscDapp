package api

import (
	"BscDapp/app/common"
	"BscDapp/app/service"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

var User = user{}

type user struct{}

// @summary 会员信息
// @tags    会员模块
// @produce json
// @router  /api/user/get-user-info  [GET]
// @success 200 {object} model.FaBscUser "执行结果"
func (a *user) GetUserInfo(r *ghttp.Request) {
	userInfo := service.User.GetUser(r)
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: userInfo, Message: common.SuccessMsg})
}
