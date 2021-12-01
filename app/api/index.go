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
	//service.ListenTask.DealGameStatus(r.Context())
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Message: common.SuccessMsg})
}

// @summary 基础配置
// @tags    系统信息
// @produce json
// @router  /api/index/base-info   [GET]
// @success 200 {object} model.BscBaseInfo "执行结果"
func (a *indexApi) BaseInfo(r *ghttp.Request) {
	res, err := service.Index.GetBaseInfo()
	if err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: res, Message: common.SuccessMsg})
}

// @summary 场次信息
// @tags    系统信息
// @produce json
// @router  /api/index/game-info   [GET]
// @success 200 {object} model.FaBscGameInfo "执行结果"
func (a *indexApi) GameInfo(r *ghttp.Request) {
	res, err := dao.FaBscGameInfo.Where("status=1").Order("round desc").One()
	if err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: res, Message: common.SuccessMsg})
}
