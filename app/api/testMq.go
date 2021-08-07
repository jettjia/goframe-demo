package api

import (
	"github.com/gogf/gf/net/ghttp"
	"my-app/app/amqp/producer"
	"my-app/library/response"
)

// 注册控制器
var TestMq = testMqApi{}

type testMqApi struct{}


// rabbitmq send
func (*testMqApi) TestMqSend(r *ghttp.Request) {
	poolName := "default"

	msg := r.GetString("msg")

	producer.DemoTest(poolName, msg)

	response.JsonExit(r, 0, "ok", msg)
}
