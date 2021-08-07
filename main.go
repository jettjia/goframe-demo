package main

import (
	_ "my-app/boot"
	_ "my-app/router"

	_ "my-app/app/amqp/consumer"
	//_ "my-app/app/task"

	"github.com/gogf/gf/frame/g"

)

func main() {
	g.Server().Run()
}
