package models

import (
	"fmt"
	"math"

	"github.com/beego/beego/v2/core/logs"
)

//片单影片关系
type MoviePieceList struct {
	Plid        int    //片单id
	Mid         int    //影片id
	MovieName   string //影片名称
	MovieNumber string //影片番号
}

/**
* 计算数量
 */
func (d *MoviePieceList) Total(plid string) int {
	total := 0

	//先计算数量
	q := `SELECT count(0) as count from piece_list_movie where plid in (%s) and status=1;`

	q = fmt.Sprintf(q, plid)

	row, err := DB.Query(q)
	defer row.Close()

	if err != nil {
		logs.Error("sql error->", q, err.Error())
		return total
	}

	if row.Next() == true {
		row.Scan(&total)
	}

	return total
}

/**
* 获取片单对应的影片列表，分页
* param 	plid 	string 	片单id
* param 	lastId	string 	传递参数，用来分页
 */
func (d *MoviePieceList) Lists(plid string, lastId int) (map[string][]string, int) {
	res := map[string][]string{}

	q := `SELECT A.id,A.mid,A.plid,B.name,B.number 
		FROM piece_list_movie A
		LEFT JOIN movie B
		ON A.mid=B.id 
		WHERE A.plid in (%s) and A.id>? and A.status=1 and B.status=1 order by id asc limit 50;`

	q = fmt.Sprintf(q, plid)

	rows, err := DB.Query(q, lastId)
	if err != nil {
		logs.Error("sql error->", q, err.Error())
		return res, 0
	}

	defer rows.Close()

	//扫描数据
	for rows.Next() {
		var id, mid int
		var pid, name, number string
		er := rows.Scan(&id, &mid, &pid, &name, &number)

		if er != nil {
			logs.Error("scan row error->", er.Error())
		}

		lastId = id

		//将名称和番号列表，对应到影片id
		tmp := []string{}
		if len(res[pid]) > 0 {
			tmp = res[pid]
		}
		tmp = append(tmp, name+" "+number)

		res[pid] = tmp
	}

	return res, lastId
}

/**
* 获取片单对应的影片列表，所有
* param 	plid 	string 	片单id
* param 	lastId	string 	传递参数，用来分页
 */
func (d *MoviePieceList) GetAll(plid string) map[string][]string {
	//符合条件的总记录数
	total := d.Total(plid)
	limit := 50

	//计算需要读取的数据数量
	all := math.Ceil(float64(total) / float64(limit))
	max := int(all)

	lastId := 0
	res := map[string][]string{}
	for i := 0; i < max; i++ {
		li := map[string][]string{}
		li, lastId = d.Lists(plid, lastId)

		//组合数据
		for k, v := range li {
			if len(res[k]) > 0 {
				res[k] = append(res[k], v...)
			} else {
				res[k] = v
			}
		}

		//得不到最后id，跳出循环
		if lastId < 1 {
			break
		}
	}

	return res
}
