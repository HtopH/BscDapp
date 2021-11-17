package api

import (
	"BscDapp/app/common"
	"BscDapp/app/service"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

var Index = indexApi{}

type indexApi struct{}

func (a *indexApi) Index(r *ghttp.Request) {
	//service.NewGame.GetPlayCount()
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Message: common.SuccessMsg})
}

func (a *indexApi) PayToGame(r *ghttp.Request) {
	//service.NewGame.Pay()
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Message: common.SuccessMsg})
}

func (a *indexApi) AdminWithdraw(r *ghttp.Request) {
	//service.NewGame.AdminWithdraw()
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Message: common.SuccessMsg})
}
