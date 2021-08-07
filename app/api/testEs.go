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