package models

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
)

//系列对象
type MovieFilm struct {
	Id        int
	Name      string //名称
	Status    int    //状态
	MovieSum  int    //影片数量
	LikeSum   int    //收藏数量
	CreatedAt string //创建时间
	UpdatedAt string //更新

	CategoryId   int    //分类id
	CategoryName string //分类名称
}

/**
* 获取列表
* param 	fid 	系列id
* param 	where   查询条件,例如 id in(1,2,3)
* param 	updatedAt 最后更新时间
 */
func (d *MovieFilm) Lists(fid int, where, updatedAt string, limit int) (lastid int, res []*MovieFilm) {
	q := `SELECT id,name,status,movie_sum,like_sum,
		created_at,ifnull(updated_at,created_at) 
		FROM movie_film_companies
		where id>? and ifnull(updated_at,created_at)>? %s
		order by id asc limit %d;`

	if len(where) > 1 {
		q = fmt.Sprintf(q, "and "+where, limit)
	} else {
		q = fmt.Sprintf(q, "", limit)
	}

	rows, err := DB.Query(q, fid, updatedAt)
	if err != nil {
		logs.Error("sql error->", q, fid, updatedAt, err.Error())
		return lastid, res
	}
	defer rows.Close()

	//扫描数据
	for rows.Next() {
		dd := new(MovieFilm)
		er := rows.Scan(&dd.Id, &dd.Name, &dd.Status, &dd.MovieSum, &dd.LikeSum, &dd.CreatedAt, &dd.UpdatedAt)

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
func (d *MovieFilm) Total(updatedAt string) int {
	res := 0

	q := `SELECT count(0) as nums FROM movie_film_companies where ifnull(updated_at,created_at)>? `

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

/**
* 获取分类关系表
* param 	where   查询条件,例如 id in(1,2,3)
 */
func (d *MovieFilm) GetMovieFilmCategoryAssociate(where string) map[string]map[string]string {
	res := map[string]map[string]string{}
	q := `SELECT A.cid,A.film_companies_id,B.name 
		FROM movie_film_companies_category_associate A
		LEFT JOIN movie_film_companies_category B
		ON A.cid=B.id 
		WHERE A.status=1 and B.status=1 %s`
	if len(where) > 1 {
		q = fmt.Sprintf(q, "and "+where)
	}

	rows, err := DB.Query(q)
	if err != nil {
		logs.Error("sql error->", q, err.Error())
		return res
	}

	//扫描数据
	for rows.Next() {
		var fid, cid, name string
		er := rows.Scan(&cid, &fid, &name)

		if er != nil {
			logs.Error("scan row error->", er.Error())
		}

		tmp := map[string]string{}
		tmp[cid] = name

		res[fid] = tmp

	}

	return res
}

/**
* 获取影片关系表
* param 	where   查询条件,例如 mid in(1,2,3)
 */
func (d *MovieFilm) GetMovieAssociate(where string) map[string]map[string]string {
	res := map[string]map[string]string{}
	q := `SELECT A.film_companies_id,A.mid,B.name 
		FROM movie_film_companies_associate A
		LEFT JOIN movie_film_companies B
		ON A.film_companies_id=B.id 
		WHERE A.status=1 and B.status=1 %s`
	if len(where) > 1 {
		q = fmt.Sprintf(q, "and "+where)
	}

	rows, err := DB.Query(q)
	if err != nil {
		logs.Error("sql error->", q, err.Error())
		return res
	}

	//扫描数据
	for rows.Next() {
		var fid, mid, name string
		er := rows.Scan(&fid, &mid, &name)

		if er != nil {
			logs.Error("scan row error->", er.Error())
		}

		tmp := map[string]string{}
		tmp[fid] = name

		res[mid] = tmp
	}
	return res
}
