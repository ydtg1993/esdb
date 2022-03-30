package es

import (
	"time"
)

//压缩图
type ImgMapJson struct {
	Img    string `json:"img"`
	BigImg string `json:"big_img"`
}

//磁链
type FluxLinkJson struct {
	Time      time.Time `json:"time"`
	IsWarning int       `json:"is-warning"`
	Meta      string    `json:"meta"`
	IsSmall   int       `json:"is-small"`
	Tooltip   int       `json:"tooltip"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
}

//影片对象
type IndexMovie struct {
	Id              int             `json:"id"`
	Number          string          `json:"number"`
	Name            string          `json:"name"`
	Time            string          `json:"time"`
	ReleaseTime     string          `json:"release_time"`
	Issued          string          `json:"issued"`
	Sell            string          `json:"sell"`
	SmallCover      string          `json:"small_cover"`
	BigCove         string          `json:"big_cove"`
	Trailer         string          `json:"trailer"`
	Map             []*ImgMapJson   `json:"map"`
	Score           float64         `json:"score"`
	ScorePeople     int             `json:"score_people"`
	CommentNum      int             `json:"comment_num"`
	WanSee          int             `json:"wan_see"`
	Seen            int             `json:"seen"`
	FluxLinkageNum  int             `json:"flux_linkage_num"`
	FluxLinkage     []*FluxLinkJson `json:"flux_linkage"`
	Status          int             `json:"status"`
	IsDownload      int             `json:"is_download"`
	IsSubtitle      int             `json:"is_subtitle"`
	IsHot           int             `json:"is_hot"`
	IsShortComment  int             `json:"is_short_comment"`
	IsUp            int             `json:"is_up"`
	NewCommentTime  time.Time       `json:"new_comment_time"`
	FluxLinkageTime time.Time       `json:"flux_linkage_time"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`

	CategoryId   int    `json:"categoty_id"`   //分类id
	CategoryName string `json:"categoty_name"` //分类名称

	DirectorId   int    `json:"director_id"`   //导演id
	DirectorName string `json:"director_name"` //导演名称

	SeriesId   int    `json:"series_id"`   //系列id
	SeriesName string `json:"series_name"` //系列名称

	FilmId   int    `json:"film_id"`   //片商id
	FilmName string `json:"film_name"` //片商名称
}

/**
* 影片对应的index结构
"number_of_shards": 1,  //每个索引的主分片数，默认值是 5 。这个配置在索引创建后不能修改,3w条数据/100M
"number_of_replicas": 0  //每个主分片的副本数，默认值是 1 。对于活动的索引库，这个配置可以随时修改。
*/
const MappingMovie = `
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
			"number":{
				"type":"text"
			},
			"name":{
				"type":"text",
				"analyzer":"ik_max_word"
			},
			"time":{
				"type":"long"
			},
			"release_time":{
				"type":"date",
				"format": "yyyy-MM-dd HH:mm:ss"
			},
			"issued":{
				"type":"keyword"
			},
			"sell":{
				"type":"keyword"
			},
			"small_cover":{
				"type":"text"
			},
			"big_cove":{
				"type":"text"
			},
			"trailer":{
				"type":"text"
			},
			"score":{
				"type":"float"
			},
			"score_people":{
				"type":"integer"
			},
			"comment_num":{
				"type":"long"
			},
			"wan_see":{
				"type":"long"
			},
			"seen":{
				"type":"long"
			},
			"flux_linkage_num":{
				"type":"integer"
			},
			"flux_linkage" : {
          		"type" : "text"
        	},
			"status":{
				"type":"long"
			},
			"map" : {
          		"type" : "text"
        	},
			"is_download":{
				"type":"integer"
			},
			"is_subtitle":{
				"type":"integer"
			},
			"is_hot":{
				"type":"integer"
			},
			"is_short_comment":{
				"type":"integer"
			},
			"is_up":{
				"type":"integer"
			},
			"new_comment_time":{
				"type":"date",
				"format": "yyyy-MM-dd HH:mm:ss"
			},
			"flux_linkage_time":{
				"type":"date",
				"format": "yyyy-MM-dd HH:mm:ss"
			},
			"created_at":{
				"type":"date",
				"format": "yyyy-MM-dd HH:mm:ss"
			},
			"updated_at":{
				"type":"date",
				"format": "yyyy-MM-dd HH:mm:ss"
			},
			"categoty_id":{
				"type":"integer"
			},
			"categoty_name":{
				"type":"keyword"
			},
			"director_id":{
				"type":"integer"
			},
			"director_name":{
				"type":"keyword"
			},
			"series_id":{
				"type":"integer"
			},
			"series_name":{
				"type":"keyword"
			},
			"film_id":{
				"type":"integer"
			},
			"film_name":{
				"type":"keyword"
			},
			"actor":{
				"type":"text",
				"analyzer":"ik_max_word"
			}
		}
	}
}`
