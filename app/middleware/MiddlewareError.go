package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

var MiddlewareErr = new(middlewareErr)

type middlewareErr struct {}


// 这里是对全局返回的错误进行统一处理
// 比如服务器500了，返回给前端的信息就是统一的一句话
func (m *middlewareErr) ErrHandle(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.Writeln("哎哟我去，服务器居然开小差了，程序员giegie要罚款了！")
	}
}