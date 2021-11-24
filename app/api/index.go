package api

import (
	"BscDapp/app/common"
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"BscDapp/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
	"strconv"
)

var Index = indexApi{}

type indexApi struct{}

func (a *indexApi) Test(r *ghttp.Request) {
	baseInfo, _ := dao.FaBscBaseInfo.Where("theKey=?", model.BaseReadKey).One()
	num, err := strconv.Atoi(baseInfo.TheValue)
	if err != nil {
		num = 1
		g.Log().Debug("Service ListenTask strconv Err:", err)
	}
	service.NewGame.ReadBlockLog(int64(num))
	//service.TimeTask.FinishBscTask()
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Message: common.SuccessMsg})
}
