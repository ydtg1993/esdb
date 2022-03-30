package es

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/olivere/elastic"
)

var ESClient *elastic.Client

var esUrl string

//初始化
func init() {
	//读取配置文件
	esHost, _ := beego.AppConfig.String("eshost")
	esPort, _ := beego.AppConfig.String("esport")
	esUrl = fmt.Sprintf("http://%s:%s", esHost, esPort)

	var er error
	ESClient, er = elastic.NewClient(elastic.SetURL(esUrl, esUrl))
	if er != nil {
		// Handle error
		panic(er.Error())
	}

	version, err := ESClient.ElasticsearchVersion(esUrl)
	if err != nil {
		// Handle error
		panic(err.Error())
	}

	fmt.Printf("Elasticsearch version %s\n", version)
}
