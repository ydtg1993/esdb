package models

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
)

//片单对象
type Piece struct {
	Id          int
	Uid         int    //创建者id
	Name        string //名称
	Cover       string //封面
	MovieSum    int    //影片数量
	LikeSum     int    //收藏的数量
	PvBrowseSum int    //pv数量
	Intro       string //片单介绍
	Status      int    //状态
	IsHot       int    //热门
	Authority   int    //是否公开
	Type        int    //类型
	Audit       int    //审核
	Remarks     string //备注
	CreatedAt   string //创建时间
	UpdatedAt   string //更新

	KeyWord    string //搜索条件组合，影片列表
	UserName   string //导演名称
	UserAvatar string //系列id
}

/**
* 从数据库中读取赛事列表,每次读取10条
* mid			id
* updatedAt		最后更新时间
* limit 		每次读取多少条
 */
func (d *Piece) Lists(mid int, updatedAt string, limit int) (lastid int, res []*Piece) {

	q := `SELECT id,uid,name,cover,movie_sum,like_sum,
			pv_browse_sum,ifnull(intro,'') intro,status,is_hot,authority,
			type,audit,ifnull(remarks,'') remarks,created_at,ifnull(updated_at,created_at) updated_at 
		FROM movie_piece_list 
		where id>? and ifnull(updated_at,created_at)>=? order by id asc limit %d;`
	q = fmt.Sprintf(q, limit)

	rows, err := DB.Query(q, mid, updatedAt)
	if err != nil {
		logs.Error("sql error->", q, mid, updatedAt, err.Error())
		return lastid, res
	}
	defer rows.Close()

	//扫描数据
	for rows.Next() {
		dd := new(Piece)
		er := rows.Scan(&dd.Id, &dd.Uid, &dd.Name, &dd.Cover, &dd.MovieSum, &dd.LikeSum,
			&dd.PvBrowseSum, &dd.Intro, &dd.Status, &dd.IsHot, &dd.Authority,
			&dd.Type, &dd.Audit, &dd.Remarks, &dd.CreatedAt, &dd.UpdatedAt)

		if er != nil {
			logs.Error("scan row error->", er.Error())
		}

		lastid = dd.Id

		res = append(res, dd)
	}

	return lastid, res
}

/**
* 总记录数
 */
func (d *Piece) Total(updatedAt string) int {
	res := 0

	q := `SELECT count(0) as nums FROM movie_piece_list where ifnull(updated_at,created_at)>? `

	row, err := DB.Query(q, updatedAt)
	if err != nil {
		logs.Error("sql error->", q, updatedAt, err.Error())
		return res
	}
	defer row.Close()

	if row.Next() == true {
		row.Scan(&res)
	}
	return res
}
