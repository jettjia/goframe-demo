package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"my-app/app/api"
)

// authHook is the HOOK function implements JWT logistics.
func MiddlewareAuth(r *ghttp.Request) {
	api.Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}

// MiddlewareCORS 跨域
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}


func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/hello", api.Hello)


		// jwt相关
		group.Group("/", func(g *ghttp.RouterGroup) {
			g.ALL("/login", api.Auth.LoginHandler)
			g.ALL("/refresh_token", api.Auth.RefreshHandler)
			g.ALL("/logout", api.Auth.LogoutHandler)
		})

		group.Group("/user", func(group *ghttp.RouterGroup) {
			//允许跨域访问
			group.Middleware(MiddlewareCORS)
			//JWT认证中间件
			group.Middleware(MiddlewareAuth)

			group.ALL("/registry", api.User.Register)
			group.ALL("/profile", api.User.Profile)
			group.ALL("/update", api.User.UpdateProfile)
		})

		group.Group("/test", func(group *ghttp.RouterGroup) {
			//允许跨域访问
			group.Middleware(MiddlewareCORS)

			group.ALL("/redis", api.TestRedis)
			group.ALL("/mq", api.TestMq)
			group.ALL("/es", api.TestEs)
			group.ALL("/csv", api.TestCsv)
		})
	})

}
