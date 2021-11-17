package router

import (
	"BscDapp/app/api"
	"BscDapp/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/swagger"
)

func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	//初始化合约
	service.NewGame.Init()

	//定时任务循环处理Bsc
	service.TimeTask.ListenTask()
	//需验证控制台登陆状态
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.ALL("/index", &api.Index)
	})
}
