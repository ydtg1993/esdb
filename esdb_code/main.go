package main

import (
	"esdb/pool"
	"fmt"
	"runtime"
	"strings"

	_ "github.com/CodyGuo/godaemon"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func init() {
	//日志天数
	var logDay = ""
	//日志路径
	var logpath = ""
	//日志级别
	var loglevel = ""
	logErr := "error"

	if logDay = os.Getenv("logday"); logDay == "" {
		logDay = "7"
	}
	if logpath = os.Getenv("logpath"); logpath == "." {
		logpath = ""
	}
	//日志级别
	loglevel = os.Getenv("loglevel")

	logName := fmt.Sprintf("%shdb-import.log", logpath)

	if len(loglevel) > 1 {
		sLevel := strings.Split(loglevel, ",")
		logErr = `"` + strings.Join(sLevel, `","`) + `"`
	}

	level := 2
	if strings.ContainsAny(loglevel, "debug") == true {
		level = 7
	}

	logCfg := fmt.Sprintf(`{"filename":"%s","level":%d,"maxdays":%s,"separate":[%s]}`, logName, level, logDay, logErr)
	//记录日志
	err := logs.SetLogger(logs.AdapterMultiFile, logCfg)
	if err != nil {
		panic(err)
	}

	// 开始前的线程数
	logs.Debug("线程数量 starting: %d\n", runtime.NumGoroutine())
}

func main() {

	//导入演员的线程
	go pool.TaskActor()

	//导入标签的线程
	go pool.TaskLabel()

	//导入系列的线程
	go pool.TaskSeries()

	//导入片商的线程
	go pool.TaskFilm()

	//导入分类的线程
	go pool.TaskCategory()

	//导入导演的线程
	go pool.TaskDirector()

	//导入片单的线程
	go pool.TaskPiece()

	//导入影片数据的线程
	pool.TaskMovie()

}
