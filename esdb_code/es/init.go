package es

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/olivere/elastic/v7"
	"os"
)

var ESClient *elastic.Client

var esUrl string

//初始化
func init() {
	//读取配置文件
	esHost := os.Getenv("eshost")
	esPort := os.Getenv("esport")
	esUrl = fmt.Sprintf("http://%s:%s", esHost, esPort)

	var er error
	ESClient, er = elastic.NewClient(elastic.SetURL(esUrl, esUrl), elastic.SetSniff(false))

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
