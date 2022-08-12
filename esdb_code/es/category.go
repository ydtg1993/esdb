package es

import (
	"time"
)

//分类对象
type IndexMovieCategory struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

/**
* 分类对应的index结构
"number_of_shards": 1,  //每个索引的主分片数，默认值是 5 。这个配置在索引创建后不能修改,3w条数据/100M
"number_of_replicas": 0  //每个主分片的副本数，默认值是 1 。对于活动的索引库，这个配置可以随时修改。
*/
const MappingMovieCategory = `
{
	"settings":{
		"analysis":{
			"analyzer":{
				"ik":{
					"tokenizer":"ik_max_word"
				}
			}
		},
		"number_of_shards": 5,
		"number_of_replicas": 0
	},
	"mappings":{
		"properties":{
			"id":{
				"type":"integer"
			},
			"name":{
				"type":"text",
				"analyzer":"ik_max_word"
			},
			"status":{
				"type":"integer"
			},
			"created_at":{
				"type":"date",
				"format": "yyyy-MM-dd HH:mm:ss"
			},
			"updated_at":{
				"type":"date",
				"format": "yyyy-MM-dd HH:mm:ss"
			}
		}
	}
}`
