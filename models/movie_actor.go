package models

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
)

//演员对象
type MovieActor struct {
	Id             int
	Name           string //演员名称
	Photo          string //演员照片
	Sex            string //演员性别
	SocialAccounts string //社交账户的json数组
	MovieSum       int    //影片数量
	LikeSum        int    //收藏数
	Status         int    //状态
	CreatedAt      string //创建时间
	UpdatedAt      string //更新

	CategoryId   int    //演员分类id
	CategoryName string //演员分类名称
}

/**
* 获取列表
* param 	aid 	演员id
* param 	where   查询条件,例如 id in(1,2,3)
* param 	updatedAt 最后更新时间
 */
func (d *MovieActor) Lists(aid int, where, updatedAt string, limit int) (lastid int, res []*MovieActor) {
	q := `SELECT id,name,ifnull(photo,''),sex,ifnull(social_accounts,''),
		movie_sum,like_sum,status,created_at,ifnull(updated_at,created_at) 
		FROM movie_actor
		where id>? and ifnull(updated_at,created_at)>? %s
		order by id asc limit %d;`

	if len(where) > 1 {
		q = fmt.Sprintf(q, "and "+where, limit)
	} else {
		q = fmt.Sprintf(q, "", limit)
	}

	rows, err := DB.Query(q, aid, updatedAt)
	if err != nil {
		logs.Error("sql error->", q, aid, updatedAt, err.Error())
		return lastid, res
	}
	defer rows.Close()

	//扫描数据
	for rows.Next() {
		dd := new(MovieActor)
		er := rows.Scan(&dd.Id, &dd.Name, &dd.Photo, &dd.Sex, &dd.SocialAccounts,
			&dd.MovieSum, &dd.LikeSum, &dd.Status, &dd.CreatedAt, &dd.UpdatedAt)

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
func (d *MovieActor) Total(updatedAt string) int {
	res := 0

	q := `SELECT count(0) as nums FROM movie_actor where ifnull(updated_at,created_at)>? `

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
* 获取演员分类关系表
* param 	where   查询条件,例如 id in(1,2,3)
 */
func (d *MovieActor) GetMovieActorCategoryAssociate(where string) map[string]map[string]string {
	res := map[string]map[string]string{}
	q := `SELECT A.cid,A.aid,B.name 
		FROM movie_actor_category_associate A
		LEFT JOIN movie_actor_category B
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
		var aid, cid, name string
		er := rows.Scan(&cid, &aid, &name)

		if er != nil {
			logs.Error("scan row error->", er.Error())
		}

		tmp := map[string]string{}
		tmp[cid] = name

		res[aid] = tmp
	}

	return res
}

/**
* 获取影片对应演员关系表
* param 	where   查询条件,例如 id in(1,2,3)
 */
func (d *MovieActor) GetMovieActorAssociate(where string) map[string][]string {
	res := map[string][]string{}
	q := `SELECT A.mid,A.aid,B.name 
		FROM movie_actor_associate A
		LEFT JOIN movie_actor B
		ON A.aid=B.id 
		WHERE %s A.status=1 and B.status=1`
	if len(where) > 1 {
		q = fmt.Sprintf(q, where+" and")
	}

	rows, err := DB.Query(q)
	if err != nil {
		logs.Error("sql error->", q, err.Error())
		return res
	}

	//扫描数据
	for rows.Next() {
		var mid, aid, name string
		er := rows.Scan(&mid, &aid, &name)

		if er != nil {
			logs.Error("scan row error->", er.Error())
		}
		//将演员名称列表，对应到影片id
		tmp := []string{}
		if len(res[mid]) > 0 {
			tmp = res[mid]
		}
		tmp = append(tmp, name)

		res[mid] = tmp
	}

	return res
}
