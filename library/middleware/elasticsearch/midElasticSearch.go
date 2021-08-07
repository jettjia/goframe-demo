package elasticsearch

import (
	"github.com/gogf/gf/frame/g"
	"strings"
)

var MidEs = midEs{}

type midEs struct {}

// es 配置读取
func (*midEs)EsConnStr(pool string) string{

	address := g.Cfg().GetString("es." + pool + ".0.Address")
	port := g.Cfg().GetString("es." + pool + ".0.Port")

	var build strings.Builder
	build.WriteString("http://")
	build.WriteString(address)
	build.WriteString(":")
	build.WriteString(port)
	connStr := build.String()

	return connStr
}

func (*midEs) EsUser(pool string) string  {
	return g.Cfg().GetString("es." + pool + ".0.User")
}

func (*midEs) EsPassword(pool string) string  {
	return g.Cfg().GetString("es." + pool + ".0.Password")
}
