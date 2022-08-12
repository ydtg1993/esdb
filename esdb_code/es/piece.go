package es

import (
	"time"
)

//片单对象
type IndexPiece struct {
	Id          int       `json:"id"`
	Uid         int       `json:"uid"`           //片单创建者id
	Name        string    `json:"name"`          //片单名称
	Cover       string    `json:"cover"`         //封面
	MovieSum    int       `json:"movie_sum"`     //影片数量
	LikeSum     int       `json:"like_sum"`      //收藏数量
	PvBrowseSum int       `json:"pv_browse_sum"` //浏览次数
	Intro       string    `json:"intro"`         //片单简介
	Status      int       `json:"status"`        //状态，1=正常
	IsHot       int       `json:"is_hot"`        //是否热门，2=热门
	Authority   int       `json:"authority"`     //是否公开
	Type        int       `json:"type"`          //创建类型，1=用户创建；2=管理员创建；3=用户默认
	Audit       int       `json:"audit"`         //审核状态，1=审核通过；0=待审核；3=审核不通过
	Remarks     string    `json:"remarks"`       //备注
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	KeyWord    string `json:"keyword"`  //搜索字段，需要凭借收录的影片名称+番号
	UserName   string `json:"username"` //创建者用户名
	UserAvatar string `json:"avatar"`   //用户头像
}

/**
* 标签对应的index结构
"number_of_shards": 1,  //每个索引的主分片数，默认值是 5 。这个配置在索引创建后不能修改,3w条数据/100M
"number_of_replicas": 0  //每个主分片的副本数，默认值是 1 。对于活动的索引库，这个配置可以随时修改。
*/
const MappingPiece = `
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
			"uid":{
				"type":"integer"
			},
			"name":{
				"type":"text",
				"analyzer":"ik_max_word"
			},
			"cover":{
				"type":"text"
			},
			"movie_sum":{
				"type":"integer"
			},
			"like_sum":{
				"type":"integer"
			},
			"pv_browse_sum":{
				"type":"integer"
			},
			"intro":{
				"type":"text"
			},
			"status":{
				"type":"integer"
			},
			"is_hot":{
				"type":"integer"
			},
			"authority":{
				"type":"integer"
			},
			"type":{
				"type":"integer"
			},
			"audit":{
				"type":"integer"
			},
			"remarks":{
				"type":"text"
			},
			"created_at":{
				"type":"date",
				"format": "yyyy-MM-dd HH:mm:ss"
			},
			"updated_at":{
				"type":"date",
				"format": "yyyy-MM-dd HH:mm:ss"
			},
			"keyword":{
				"type":"text",
				"analyzer":"ik_max_word"
			},
			"username":{
				"type":"text"
			},
			"avatar":{
				"type":"text"
			}
		}
	}
}`
