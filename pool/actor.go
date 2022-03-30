package pool

import (
	"encoding/json"
	"esdb/es"
	"esdb/models"
	"esdb/rd"
	"esdb/runtimes"
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

var actorLock sync.Mutex

//将演员数据写入index
func ActorsDo() {
	//线程锁
	actorLock.Lock()
	defer actorLock.Unlock()

	upTime := time.Now()

	//索引的名称
	indexName := "movie_actor"
	//redis的key
	redisKey := IndexMap[indexName]
	//索引的mapping
	mapping := EsMap[indexName]

	//最后更新时间
	lastTime := rd.StringGet(redisKey)
	if len(lastTime) < 1 {
		lastTime = "1970-01-01 00:00:00"
	}

	//创建movie的index
	es.CheckIndex(indexName, mapping)

	//计算总记录数
	ma := new(models.MovieActor)
	total := ma.Total(lastTime)
	limit := 50

	if total < 1 {
		//空数据，不处理
		logs.Debug(indexName, "执行完成时间:", 0, "秒")
		return
	}

	//计算需要读取的数据数量
	all := math.Ceil(float64(total) / float64(limit))
	max := int(all)

	//开启多线程处理数据
	multThread := new(runtimes.DoRun)
	multThread.Count = 0
	multThread.Err = false

	//用来分页
	lastid := 0
	for i := 0; i < max; i++ {
		var res []*models.MovieActor
		mRes := map[string]*models.MovieActor{}
		IDs := []string{}

		lastid, res = ma.Lists(lastid, "", lastTime, limit)

		//计算扩展分类
		for _, v := range res {
			//得到id数组
			IDs = append(IDs, strconv.Itoa(v.Id))
			//最终结果的数据格式
			mRes[strconv.Itoa(v.Id)] = v
		}

		//得到分类id，用分号切割
		aids := strings.Join(IDs, ",")
		if len(aids) > 0 {
			mCagegory := ma.GetMovieActorCategoryAssociate(fmt.Sprintf(" A.aid in (%s) ", aids))

			//拼接最终结果
			for key, val := range mCagegory {
				for cid, cname := range val {
					mRes[key].CategoryId, _ = strconv.Atoi(cid)
					mRes[key].CategoryName = cname
				}
			}
		}

		//将数据写入index
		for _, v := range mRes {
			mTmp := map[string]interface{}{}
			mTmp["id"] = v.Id
			mTmp["name"] = v.Name
			mTmp["photo"] = v.Photo
			mTmp["sex"] = v.Sex
			mTmp["social_accounts"] = v.SocialAccounts
			mTmp["movie_sum"] = v.MovieSum
			mTmp["like_sum"] = v.LikeSum
			mTmp["status"] = v.Status
			mTmp["created_at"] = v.CreatedAt
			mTmp["updated_at"] = v.UpdatedAt
			mTmp["categoty_id"] = v.CategoryId
			mTmp["categoty_name"] = v.CategoryName

			byteTxt, _ := json.Marshal(mTmp)
			txt := string(byteTxt)

			//进入多线程处理写入
			multThread.Work(indexName, strconv.Itoa(v.Id), txt)
			logs.Debug(indexName, v.Id, "总记录数=", total, "计数器=", multThread.Count)
		}

		//得不到最后id，跳出循环
		if lastid < 1 {
			break
		}
	}

	//线程等待
	for {
		//线程执行完成
		if multThread.Count == total {
			//导入成功了，标记最有更新时间
			rd.StringSet(redisKey, upTime.Add(-1*time.Minute).Format("2006-01-02 15:04:05"))
			logs.Debug(indexName, "本次导入成功", multThread.Count, "条,最后的时间节点", upTime.Add(-time.Minute*1).Format("2006-01-02 15:04:05"))
			break
		}

		//线程中存在错误
		if multThread.Err == true {
			logs.Debug(indexName, "数据导入失败，请查看错误日志")
			break
		}

		//多线程执行超时，设置一个超时时间
		endTime := time.Now()
		diff := endTime.Unix() - upTime.Unix()
		timeOut := int64(total) / 5

		//线程最低不能小于10秒
		if timeOut < 10 {
			timeOut = 10
		}

		if diff >= timeOut {
			logs.Debug(indexName, "线程执行超时", timeOut)
			break
		}
	}

	logs.Debug(("线程数量 finished: %d\n"), runtime.NumGoroutine())

	//计算结束时间
	endTime := time.Now()
	diff := endTime.Unix() - upTime.Unix()
	logs.Debug(indexName, "执行完成时间:", diff, "秒")
}
