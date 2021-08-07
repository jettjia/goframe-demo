package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"my-app/library/response"
)

// 注册控制器
var TestRedis = testRedisApi{}

type testRedisApi struct{}

// redis
func (*testRedisApi) TestRedis(r *ghttp.Request) {
	g.Redis().Do("SET", "k", "v")
	v, _ := g.Redis().DoVar("GET", "k")
	response.JsonExit(r, 0, "ok", v.String())
}
