package service

import (
	"BscDapp/app/dao"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

// Middleware 中间件管理服务
var Middleware = middleware{}

type middleware struct{}

func (s *middleware) Auth(r *ghttp.Request) {
	addr := r.Header.Get("address")
	if addr == "" {
		addr = r.GetString("address")
	}
	if addr == "" {
		_ = r.Response.WriteJsonExit(g.Map{"code": http.StatusUnauthorized, "data": nil, "message": "请注册"})
	}
	userInfo, err := dao.FaBscUser.Where("address=?", addr).One()
	if err != nil || userInfo == nil {
		g.Log().Debug("Service Middleware User Find Err:", err)
		_ = r.Response.WriteJsonExit(g.Map{"code": http.StatusUnauthorized, "data": nil, "message": "请注册"})
	}
	r.SetCtxVar("user", userInfo)
	r.Middleware.Next()
}

//设置跨域信任域名
func (s *middleware) SetCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"localhost:8080", "localhost:8081", "183.3.158.22:8597"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

func (s *middleware) MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
