package api

import (
	"BscDapp/app/common"
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"BscDapp/app/service"
	"github.com/gogf/gf/frame/g"
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

	//n2 := service.GetNo()
	//n3 := service.GetBase(n2)
	//n1 := service.GetPercent()
	//g.Log().Debug(n2, n3, n1)
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
// @success 200 {object} model.GameIndexInfo "执行结果"
func (a *indexApi) GameInfo(r *ghttp.Request) {
	data := model.GameIndexInfo{}
	res, err := dao.FaBscGameInfo.Where("status=1").Order("round desc").One()
	if err != nil {
		g.Log().Debug("Api Index GameInfo GameInfo One Err:", err)
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	data.FaBscGameInfo = res
	data.TotalInvest, err = dao.FaBscUserGame.Where("game_round=?", res.Round).Sum(dao.FaBscUserGame.Columns.InvestNum)
	if err != nil {
		g.Log().Debug("Api Index GameInfo UserGame InvestNum Err:", err)
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: data, Message: common.SuccessMsg})
}

// @summary 收益排行
// @tags    系统信息
// @produce json
// @param   page formData int false "页码"
// @param   size formData int false "每页数量"
// @router  /api/index/get-reward-top  [GET]
// @success 200 {object} model.UserRewardTop "执行结果"
func (a *indexApi) GetRewardTop(r *ghttp.Request) {
	var (
		req *model.PageReq
	)
	if err := r.Parse(&req); err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	data, err := service.User.RewardTop(r.Context(), req)
	if err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: data, Message: common.SuccessMsg})
}
