package pool

import (
	"encoding/json"
	"esdb/es"
	"esdb/models"
	"esdb/runtimes"
	"github.com/beego/beego/v2/core/logs"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var movieLock sync.Mutex

//将数据写入index
func MovieDo() {
	//线程锁
	movieLock.Lock()
	defer movieLock.Unlock()

	upTime := time.Now()

	//索引的名称
	indexName := "movie"
	//redis的key

	//索引的mapping
	mapping := EsMap[indexName]

	//创建movie的index
	es.CheckIndex(indexName, mapping)

	//最后更新时间
	lastTime := es.SearchLast(indexName)
	if len(lastTime) < 1 {
		lastTime = "1970-01-01 00:00:00"
	}
	logs.Debug(indexName, "时间节点:", lastTime)
	//计算总记录数
	ma := new(models.Movie)
	total := ma.Total(lastTime)
	limit := 100

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
		var res []*models.Movie
		mRes := map[string]*models.Movie{}
		IDs := []string{}

		lastid, res = ma.Lists(lastid, lastTime, limit)

		//计算扩展分类
		for _, v := range res {
			//得到id数组
			IDs = append(IDs, strconv.Itoa(v.Id))
			//最终结果的数据格式
			mRes[strconv.Itoa(v.Id)] = v
		}

		//将数据写入index
		for _, v := range mRes {
			mTmp := map[string]interface{}{}
			mTmp["id"] = v.Id
			mTmp["number"] = strings.ToLower(v.Number)
			mTmp["name"] = v.Name
			mTmp["release_time"] = v.ReleaseTime
			mTmp["status"] = v.Status
			mTmp["small_cover"] = v.SmallCover
			mTmp["big_cove"] = v.BigCove
			mTmp["trailer"] = v.Trailer
			mTmp["is_download"] = v.IsDownload
			mTmp["is_subtitle"] = v.IsSubtitle
			mTmp["is_hot"] = v.IsHot
			mTmp["is_short_comment"] = v.IsShortComment
			mTmp["is_up"] = v.IsUp
			mTmp["score"] = v.Score
			mTmp["score_people"] = v.ScorePeople
			mTmp["wan_see"] = v.WanSee
			mTmp["seen"] = v.Seen
			mTmp["new_comment_time"] = v.NewCommentTime
			mTmp["flux_linkage_time"] = v.FluxLinkageTime
			mTmp["created_at"] = v.CreatedAt
			mTmp["updated_at"] = v.UpdatedAt
			mTmp["weight"] = v.Weight
			mTmp["cid"] = v.Cid

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
