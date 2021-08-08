package router

import (
	"my-app/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/hello", api.Hello)

		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/user/registry", api.User.Register)
			group.ALL("/user/profile", api.User.Profile)
			group.ALL("/user/update", api.User.UpdateProfile)
		})

		group.ALL("/test/redis", api.TestRedis)
		group.ALL("/test/mq", api.TestMq)
		group.ALL("/test/es", api.TestEs)
		group.ALL("/test/csv", api.TestCsv)
	})


}
