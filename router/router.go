package router

import (
	"BscDapp/app/api"
	"BscDapp/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/swagger"
	"net/http"
)

func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	//初始化合约
	service.NewGame.Init()
	//定时任务循环处理Bsc
	//service.TimeTask.ListenTask()
	//动态路由
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware.SetCORS)
		group.ALL("/index", &api.Index)
		group.Middleware(service.Middleware.Auth)
		group.ALL("/user", &api.User)
	})
	s.BindStatusHandler(http.StatusNotFound, func(r *ghttp.Request) {
		_ = r.Response.WriteJsonExit(g.Map{"code": http.StatusNotFound, "data": nil, "message": ""})
	})
	s.BindStatusHandler(http.StatusInternalServerError, func(r *ghttp.Request) {
		_ = r.Response.WriteJsonExit(g.Map{"code": http.StatusInternalServerError, "data": nil, "message": ""})
	})
}
