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
	Weight          int             `json:"weight"`
	Cid             int             `json:"cid"` //分类id

}

/**
* 影片对应的index结构
"number_of_shards": 1,  //每个索引的主分片数，默认值是 5 。这个配置在索引创建后不能修改,3w条数据/100M
"number_of_replicas": 0  //每个主分片的副本数，默认值是 1 。对于活动的索引库，这个配置可以随时修改。
*/
const MappingMovie = `
{
    "settings": {
        "analysis": {
            "analyzer": {
                "st_ik": {
                    "type": "custom",
                    "tokenizer": "ik_smart",
                    "char_filter": [
                        "tsconvert",
                        "stconvert"
                    ]
                },
				"number_analyzer": {
          			"tokenizer": "number_tokenizer"
        		}
            },
			"tokenizer":{
				"number_tokenizer":{
					"type": "edge_ngram",
					"min_gram": 2,
					"max_gram": 12,
					"token_chars": [
            			"letter",
            			"digit"
          			],
					"symbol":["-","_","."]
				}
			},
            "char_filter": {
                "tsconvert": {
                    "type": "stconvert",
                    "convert_type": "t2s"
                },
                "stconvert": {
                    "type": "stconvert",
                    "convert_type": "s2t"
                }
            },
			"number_of_shards": 5,
            "number_of_replicas": 0
        }
    },
    "mappings": {
        "properties": {
            "id": {
                "type": "integer"
            },
            "number": {
                "type": "text",
				"fields":{
					"nb":{
						"type": "text",
						"analyzer": "number_analyzer"
					}
				}
            },
			"name": {
                "type": "text",
                "fields": {
                    "spy": {
                        "type": "text",
                        "analyzer": "ik_smart"
                    },
                    "st": {
                        "type": "text",
                        "analyzer": "st_ik"
                    }
                }
            },
			"cid": {
                "type": "integer"
            },
            "release_time": {
                "type": "date",
                "format": "yyyy-MM-dd HH:mm:ss"
            },
            "small_cover": {
                "type": "text"
            },
            "big_cove": {
                "type": "text"
            },
            "trailer": {
                "type": "text"
            },
			"status": {
                "type": "long"
            },
            "is_download": {
                "type": "integer"
            },
            "is_subtitle": {
                "type": "integer"
            },
            "is_hot": {
                "type": "integer"
            },
            "is_short_comment": {
                "type": "integer"
            },
            "is_up": {
                "type": "integer"
            },
            "score": {
                "type": "float"
            },
            "score_people": {
                "type": "integer"
            },
            "wan_see": {
                "type": "long"
            },
            "seen": {
                "type": "long"
            },
            "new_comment_time": {
                "type": "date",
                "format": "yyyy-MM-dd HH:mm:ss"
            },
            "flux_linkage_time": {
                "type": "date",
                "format": "yyyy-MM-dd HH:mm:ss"
            },
			"updated_at": {
                "type": "date",
                "format": "yyyy-MM-dd HH:mm:ss"
            },
            "created_at": {
                "type": "date",
                "format": "yyyy-MM-dd HH:mm:ss"
            },
			"weight": {
                "type": "integer"
            }
        }
    }
}`
