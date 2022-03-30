package pool

import (
	"context"
	"esdb/es"
	"esdb/rd"
	"fmt"
)

//删除索引
func DelIndex(name string) {

	//删除redis
	redisKey := IndexMap[name]
	rd.DelKey(redisKey)

	//删除索引
	es.DelIndex(name)
}

//测试索引情况
func TestIndex() {
	ctx := context.Background()

	indexName := "movie"
	//检测index是否存在
	ct, err := es.ESClient.IndexExists(indexName).Do(ctx)
	if err != nil {
		fmt.Println("es check index error->", err.Error())
		return
	}

	fmt.Println("测试通过", ct)
}
