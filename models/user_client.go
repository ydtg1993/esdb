package models

import (
	"fmt"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
)

//用户对象
type UserClient struct {
	Id       int
	UserName string //用户名
	Avatar   string //状态
}

/**
* 获取用户列表,通过uid
* param 	ids 	用户列表
 */
func (d *UserClient) Lists(ids string) map[string]*UserClient {
	q := `SELECT id,nickname as username,avatar
		FROM user_client
		where id in (%s) limit 100;`

	q = fmt.Sprintf(q, ids)

	res := map[string]*UserClient{}

	rows, err := DB.Query(q)
	if err != nil {
		logs.Error("sql error->", q)
		return res
	}
	defer rows.Close()

	//扫描数据
	for rows.Next() {
		dd := new(UserClient)
		er := rows.Scan(&dd.Id, &dd.UserName, &dd.Avatar)

		if er != nil {
			logs.Error("scan row error->", er.Error())
		}

		res[strconv.Itoa(dd.Id)] = dd
	}

	return res
}
