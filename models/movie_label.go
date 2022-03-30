package models

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
)

//标签对象
type MovieLabel struct {
	Id        int
	Name      string //名称
	Status    int    //状态
	CreatedAt string //创建时间
	UpdatedAt string //更新

	CategoryId   int    //分类id
	CategoryName string //分类名称
}

/**
* 获取列表
* param 	lid 	标签id
* param 	where   查询条件,例如 id in(1,2,3)
* param 	updatedAt 最后更新时间
 */
func (d *MovieLabel) Lists(lid int, where, updatedAt string, limit int) (lastid int, res []*MovieLabel) {
	q := `SELECT id,name,status,created_at,ifnull(updated_at,created_at) 
		FROM movie_label
		where id>? and ifnull(updated_at,created_at)>? %s
		order by id asc limit %d;`

	if len(where) > 1 {
		q = fmt.Sprintf(q, "and "+where, limit)
	} else {
		q = fmt.Sprintf(q, "", limit)
	}

	rows, err := DB.Query(q, lid, updatedAt)
	if err != nil {
		logs.Error("sql error->", q, lid, updatedAt, err.Error())
		return lastid, res
	}
	defer rows.Close()

	//扫描数据
	for rows.Next() {
		dd := new(MovieLabel)
		er := rows.Scan(&dd.Id, &dd.Name, &dd.Status, &dd.CreatedAt, &dd.UpdatedAt)

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
func (d *MovieLabel) Total(updatedAt string) int {
	res := 0

	q := `SELECT count(0) as nums FROM movie_label where ifnull(updated_at,created_at)>? `

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
func (d *MovieLabel) GetMovieLabelCategoryAssociate(where string) map[string]map[string]string {
	res := map[string]map[string]string{}
	q := `SELECT A.cid,A.lid,B.name 
		FROM movie_label_category_associate A
		LEFT JOIN movie_label_category B
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
		var lid, cid, name string
		er := rows.Scan(&cid, &lid, &name)

		if er != nil {
			logs.Error("scan row error->", er.Error())
		}

		tmp := map[string]string{}
		tmp[cid] = name

		res[lid] = tmp

	}

	return res
}
