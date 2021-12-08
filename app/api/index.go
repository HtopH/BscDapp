package api

import (
	"BscDapp/app/common"
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"BscDapp/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
	"time"
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
	data, err := service.Index.GetGameBaseInfo()
	if err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: data, Message: common.SuccessMsg})
}

// @summary 收益排行
// @tags    系统信息
// @produce json
// @param   page formData int false "页码"
// @param   size formData int false "每页数量"
// @router  /api/index/get-reward-top  [POST]
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

//初始化合约数据表
//func (a *indexApi) SystemInit(r *ghttp.Request) {
//	dao.FaBscCredit.Where("id>0").Delete()
//	dao.FaBscGameInfo.Where("id>0").Delete()
//	dao.FaBscListenLog.Where("id>0").Delete()
//	dao.FaBscTask.Where("id>0").Delete()
//	dao.FaBscTransfer.Where("id>0").Delete()
//	dao.FaBscUserGame.Where("id>0").Delete()
//	dao.FaBscUserTicket.Where("id>0").Delete()
//
//	dao.FaBscUser.Where("id>0").Delete()
//	dao.FaBscUser.OmitEmpty().Save(g.Map{"id": 1, "address": model.OwnAddr, "ref_id": 0})
//	dao.FaBscBaseInfo.Where("theKey=?", model.BaseSpendKey).Update(g.Map{"theValue": 0})
//
//}

// @summary 登录
// @tags    系统信息
// @produce json
// @param   address formData string true "钱包地址"
// @router  /api/index/login  [POST]
// @success 200 {object} model.UserRewardTop "执行结果"
func (a *indexApi) Login(r *ghttp.Request) {
	address := r.GetString("address")
	r.Cookie.Set("address", address)
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Message: common.SuccessMsg})
}

// @summary 公告
// @tags    系统信息
// @produce json
// @router  /api/index/notice  [GET]
// @success 200 {object} service.JsonResponse "执行结果"
func (a *indexApi) Notice(r *ghttp.Request) {
	res := service.GetConfig("notice")
	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: res, Message: common.SuccessMsg})
}

// @summary 修改粮草消耗
// @tags    系统信息
// @produce json
// @param   spend formData float64 true "数量"
// @router  /api/index/set-spend  [POST]
// @success 200 {object} service.JsonResponse "执行结果"
func (a *indexApi) SetSpend(r *ghttp.Request) {
	spend := r.GetFloat64("spend")
	if spend <= 0 {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: common.Failure})
	}
	res, err := service.NewGame.SetSpend(spend)
	if err != nil {
		_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	dao.FaBscBaseInfo.Where("theKey=?", model.BaseSpendKey).Update(g.Map{"theValue": spend, "updated": time.Now().Unix()})

	_ = r.Response.WriteJsonExit(service.JsonResponse{Code: http.StatusOK, Data: res, Message: common.SuccessMsg})
}
