# goframe

## gf-cli

```
wget -O gf https://github.com/gogf/gf-cli/releases/download/v1.16.3/gf_linux_amd64 && chmod +x gf && ./gf install
```

```shell
#下载
wget https://goframe.org/cli/linux_amd64/gf 

#安装gf工具到/usr/local/go/bin
./gf install 

#查看版本
gf -v 


# win
https://goframe.org/cli/windows_amd64/gf.exe
```



```shell
# 初始化一个项目
gf init my-app
cd my-app
gf run main.go

# 生成dao
gf gen dao

# 生产model，测试无用
gf gen model -path ./app/model -c config/config.yaml -g sys -t sys_user


# 生产接口文件
gf swagger --pack

# docker
gf docker main.go -p -t jettjia/gf-demos:test  
```



## 简单案例

### 目录结构说明

```shell
├── app
│   ├── amqp
│   ├── api
│   ├── dao
│   ├── define
│   ├── model
│   ├── service
│   ├── shared
│   └── task
├── bin
│   └── linux_amd64
├── boot
│   └── boot.go
├── config
│   └── config.toml
├── docker
├── Dockerfile
├── document
├── go.mod
├── go.sum
├── i18n
├── library
│   ├── middleware
│   └── response
├── main.go
├── packed
│   ├── packed.go
│   └── swagger.go
├── public
│   ├── html
│   ├── plugin
│   └── resource
├── README.MD
├── router
│   └── router.go
├── swagger
│   └── swagger.json
└── template

```



### 定时任务

https://goframe.org/pages/viewpage.action?pageId=1114363#gtimer(任务定时器)-单例任务

```go
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
```



### redis

https://goframe.org/display/gf/NoSQL+Redis

```go
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

```

```
127.0.0.1:8199/test/redis/test-redis
```



### amqp-rabbitmq

https://github.com/streadway/amqp



```go
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

```

```
127.0.0.1:8199/test/mq/test-mq-send
{
    "msg" : "666"
}
```



### ES

github.com/olivere/elastic

https://www.cnblogs.com/Yemilice/p/12887812.html

```go
package api

import (
	"context"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/olivere/elastic/v7"
	"my-app/library/middleware/elasticsearch"
	"my-app/library/response"
	"time"
)

// 注册控制器
var TestEs = testEsApi{}

type testEsApi struct{}

// es create index
func (*testEsApi) CreateIndex(r *ghttp.Request) {
	esUrl := elasticsearch.MidEs.EsConnStr("default")
	esUser := elasticsearch.MidEs.EsUser("default")
	esPassword := elasticsearch.MidEs.EsPassword("default")

	fmt.Println(esUrl)
	fmt.Println(esUser)
	fmt.Println(esPassword)
	fmt.Println("=================")
	client, err := elastic.NewClient(
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetURL(esUrl),
		elastic.SetBasicAuth(esUser, esPassword),
		elastic.SetSniff(false),
	)
	fmt.Println(err)
	if err != nil {
		response.JsonExit(r, 1, "err", err)
	}

	indexname := "test-index" // indexname 你可以想成表名
	client.CreateIndex(indexname).Do(context.Background())

	response.JsonExit(r, 0, "ok")
}
```

```
127.0.0.1:8199/test/es/create-index
```



### 链路追踪 todo



### 第三方请求封装 todo



### csv导出 todo



### oss上传 todo



### 腾讯上传 todo





## 案例运行

```
# 注册
127.0.0.1:8199/user/registry
{
    "passport" : "222",
    "password" : "23",
    "nickname" : "jett"
}

# 获取
127.0.0.1:8199/user/profile
{
    "userId" : 1
}

# 修改
127.0.0.1:8199/user/update
{
    "id" : 1,
    "passport" : "222",
    "nickname" : "jett2"
}
```

