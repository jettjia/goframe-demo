package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"my-app/library/response"
)

var MiddlewareParam = new(middlewareTest)

type middlewareTest struct {}

// 这里是模拟一个中间件
// 这里获取 参数里没有 abc = 123
// 就提示不通过
func (m *middlewareTest) Check(r *ghttp.Request) {
	paramStr := r.GetInt("abc")
	if paramStr != 123 {
		response.JsonExit(r, 1, "err", "请求的参数里必须包含有 abc=123")
	}

	r.Middleware.Next()
}