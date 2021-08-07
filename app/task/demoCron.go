package task

import (
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
)

func init()  {
	cronT()
}

func cronT() {
	gcron.AddSingleton("* * * * * *", func() {
		glog.Println("doing")
	})
}