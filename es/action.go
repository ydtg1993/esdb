package es

import (
	"context"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/core/logs"
)

/**
*  检测index是否存在
* param 	string 		indexName    索引名称
* param 	string 		mappingName  索引结构
 */
func CheckIndex(indexName, mappingName string) {
	ctx := context.Background()

	//检测index是否存在
	exists, err := ESClient.IndexExists(indexName).Do(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "503") == true {
			exists = false
		} else {
			logs.Error("es check index error->", indexName, err.Error())

			return
		}
	}

	//不存在就创建
	if !exists {
		// Create a new index.
		createIndex, err := ESClient.CreateIndex(indexName).BodyString(mappingName).Do(ctx)
		if err != nil {
			// Handle error
			logs.Error("es create index error->", indexName, err.Error())
			return
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			logs.Error("es Acknowledged index error->", indexName)
			return
		}
		fmt.Println("create index", indexName, "success")
	}
}

/**
* 将数据写入index
* param 	string 		indexName    索引名称
* param 	string 		mappingName  写入数据
 */
func AddIndex(indexName, indexId, data string) string {
	res := ""
	ctx := context.Background()

	ESClient.Ping(esUrl).Do(ctx)

	put, err := ESClient.Index().
		Index(indexName).
		Type("_doc").
		Id(indexId).
		BodyString(data).
		Do(ctx)

	if err != nil {
		// Handle error
		logs.Error("写入数据 error->", indexName, err.Error(), data, put)
		return ""
	}

	logs.Debug(fmt.Sprintf("调试 %s to index %s, 状态 %d,版本 %d,结果 %s,seq %d \n", put.Id, put.Index, put.Status, put.Version, put.Result, put.SeqNo))

	res = put.Id

	return res
}

/**
* 删除一个index
 */
func DelIndex(indexName string) {

	ctx := context.Background()

	ESClient.Ping(esUrl).Do(ctx)

	exists, err := ESClient.IndexExists(indexName).Do(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "503") == true {
			exists = false
		} else {
			logs.Error("es check index error->", indexName, err.Error())
			return
		}
	}

	if exists == false {
		fmt.Println("索引不存在", indexName)
		return
	}

	del, er := ESClient.DeleteIndex(indexName).Do(ctx)

	if er != nil {
		// Handle error
		logs.Error("es del index error->", indexName, err.Error())
		return
	}
	if !del.Acknowledged {
		// Not acknowledged
		logs.Error("es Acknowledged index error->", indexName)
		return
	}
}
