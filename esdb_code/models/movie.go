package models

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
)

//电影对象
type Movie struct {
	Id              int
	Number          string  //番号
	Name            string  //影片名称
	Time            string  //播放时长（秒）
	ReleaseTime     string  //发布时间
	Issued          string  //发行
	Sell            string  //卖家
	SmallCover      string  //小封面
	BigCove         string  //大封面
	Trailer         string  //预告片
	Map             string  //组图
	Score           float64 //积分
	ScorePeople     int     //评分人数
	CommentNum      int     //评论数
	WanSee          int     //想看数量
	Seen            int     //看过数量
	FluxLinkageNum  int     //磁链数量
	FluxLinkage     string  //磁链
	Status          int     //状态
	IsDownload      int     //是否可下载
	IsSubtitle      int     //是否包含字幕
	IsHot           int     //是否热门
	IsShortComment  int     //是否包含短评
	IsUp            int     //上架还是下架
	NewCommentTime  string  //最新评论时间
	FluxLinkageTime string  //磁链更新时间
	CreatedAt       string  //创建时间
	UpdatedAt       string  //更新
	Weight          int
	Cid             int
}

/**
* 从数据库中读取赛事列表,每次读取10条
* mid			id
* updatedAt		最后更新时间
* limit 		每次读取多少条
 */
func (d *Movie) Lists(mid int, updatedAt string, limit int) (lastid int, res []*Movie) {

	q := `
	SELECT id,number,ifnull(name,''),ifnull(time,0),ifnull(release_time,'2006-01-02 00:00:00'),
		ifnull(issued,''),ifnull(sell,''),ifnull(small_cover,''),ifnull(big_cove,''),ifnull(trailer,''),
		ifnull(map,''),score,score_people,comment_num,wan_see,
		seen,flux_linkage_num,ifnull(flux_linkage,''),status,is_download,
		is_subtitle,is_hot,is_short_comment,is_up,ifnull(new_comment_time,'2006-01-02 00:00:00'),
		ifnull(flux_linkage_time,'2006-01-02 00:00:00'),created_at,ifnull(updated_at,created_at),
		weight,cid
	FROM movie
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
		dd := new(Movie)
		er := rows.Scan(&dd.Id, &dd.Number, &dd.Name, &dd.Time, &dd.ReleaseTime,
			&dd.Issued, &dd.Sell, &dd.SmallCover, &dd.BigCove, &dd.Trailer,
			&dd.Map, &dd.Score, &dd.ScorePeople, &dd.CommentNum, &dd.WanSee,
			&dd.Seen, &dd.FluxLinkageNum, &dd.FluxLinkage, &dd.Status, &dd.IsDownload,
			&dd.IsSubtitle, &dd.IsHot, &dd.IsShortComment, &dd.IsUp, &dd.NewCommentTime,
			&dd.FluxLinkageTime, &dd.CreatedAt, &dd.UpdatedAt, &dd.Weight, &dd.Cid)

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
func (d *Movie) Total(updatedAt string) int {
	res := 0

	q := `SELECT count(0) as nums FROM movie where ifnull(updated_at,created_at)>? `

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
