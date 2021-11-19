package api

import (
	"BscDapp/app/common"
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"BscDapp/app/service"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"net/http"
)

var Index = indexApi{}

type indexApi struct{}

func (a *indexApi) Test(r *ghttp.Request) {
	baseInfo, _ := dao.FaBscBaseInfo.Where("theKey=?", model.BaseReadKey).One()
	service.NewGame.ReadBlockLog(gconv.Int64(baseInfo.TheValue))
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Message: common.SuccessMsg})
}
