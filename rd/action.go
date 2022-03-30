package rd

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

/**
* 写入字符串
* param 	key 	主键
* param		val 	字符串
 */
func StringSet(key, val string) {
	RD := RdClient.Get()
	defer RD.Close()

	key = fmt.Sprintf("%s:%s", RdDB, key)

	//写入redis
	RD.Send("set", key, val)

}

/**
* 读取字符串
* param 	key 	主键
 */
func StringGet(key string) string {
	RD := RdClient.Get()
	defer RD.Close()

	key = fmt.Sprintf("%s:%s", RdDB, key)

	res, _ := redis.String(RD.Do("get", key))
	return res
}

/**
* 删除一个key
 */
func DelKey(key string) {
	RD := RdClient.Get()
	defer RD.Close()

	key = fmt.Sprintf("%s:%s", RdDB, key)
	RD.Do("del", key)
}
