package api

import (
	"BscDapp/app/common"
	"BscDapp/app/service"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

var Index = indexApi{}

type indexApi struct{}

func (a *indexApi) Test(r *ghttp.Request) {
	//baseInfo, _ := dao.FaBscBaseInfo.Where("theKey=?", model.BaseReadKey).One()
	//num, err := strconv.Atoi(baseInfo.TheValue)
	//if err != nil {
	//	num = 1
	//	g.Log().Debug("Service ListenTask strconv Err:", err)
	//}
	//service.NewGame.ReadBlockLog(int64(num))
	//service.ListenTask.DealGameStatus(r.Context())
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Message: common.SuccessMsg})
}
