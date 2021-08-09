package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"my-app/app/api"
	"my-app/app/middleware"
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

// 额外的中间件判断
func MiddlewareABCTest(r *ghttp.Request) {
	middleware.MiddlewareParam.Check(r)
	r.Middleware.Next()
}

// 全局错误处理
func MiddlewareErrHandler(r *ghttp.Request) {
	middleware.MiddlewareErr.ErrHandle(r)
	r.Middleware.Next()
}


func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		// 全局错误处理
		group.Middleware(MiddlewareErrHandler)

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
			//group.Middleware(MiddlewareAuth)
			// 额外的中间件判断
			//group.Middleware(MiddlewareABCTest)

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
