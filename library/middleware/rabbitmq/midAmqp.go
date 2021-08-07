package rabbitmq

import (
	"github.com/gogf/gf/frame/g"
	"strings"
)

var MidAmpq = midAmpq{}

type midAmpq struct {}

// mq 配置读取
func (*midAmpq)MqConnStr(pool string) string{

	address := g.Cfg().GetString("rabbitmq." + pool + ".0.Address")
	port := g.Cfg().GetString("rabbitmq." + pool + ".0.Port")
	user := g.Cfg().GetString("rabbitmq." + pool + ".0.User")
	password := g.Cfg().GetString("rabbitmq." + pool + ".0.Password")
	vhost := g.Cfg().GetString("rabbitmq." + pool + ".0.Vhost")

	var build strings.Builder
	build.WriteString("amqp://")
	build.WriteString(user)
	build.WriteString(":")
	build.WriteString(password)
	build.WriteString("@")
	build.WriteString(address)
	build.WriteString(":")
	build.WriteString(port)
	build.WriteString(vhost)
	connStr := build.String()

	return connStr
}
