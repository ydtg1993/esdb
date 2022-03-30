package main

import (
	"esdb/pool"
	_ "esdb/routers"
	"fmt"
	"os"
	"runtime"
	"strings"

	_ "github.com/CodyGuo/godaemon"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"

	"github.com/urfave/cli"
)

func main() {

	//如果后面跟了参数，直接进行命令行操作
	if len(os.Args) > 1 {
		cliDo()
		fmt.Println("执行完成")
		return
	}

	//日志天数
	logDay, _ := beego.AppConfig.Int("logday")
	//日志路径
	logpath, _ := beego.AppConfig.String("logpath")
	//日志级别
	loglevel, _ := beego.AppConfig.String("loglevel")

	if logDay < 1 {
		logDay = 7
	}
	if logpath == "." {
		logpath = ""
	}
	logName := fmt.Sprintf("%sesdb.log", logpath)

	logErr := "error"
	if len(loglevel) > 1 {
		sLevel := strings.Split(loglevel, ",")
		logErr = `"` + strings.Join(sLevel, `","`) + `"`
	}

	level := 2
	if strings.ContainsAny(loglevel, "debug") == true {
		level = 7
	}

	logCfg := fmt.Sprintf(`{"filename":"%s","level":%d,"maxdays":%d,"separate":[%s]}`, logName, level, logDay, logErr)

	//记录日志
	logs.SetLogger(logs.AdapterMultiFile, logCfg)

	// 开始前的线程数
	logs.Debug(("线程数量 starting: %d\n"), runtime.NumGoroutine())

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

/**
* 执行命令行操作
 */
func cliDo() {

	app := cli.NewApp()
	app.Version = "1.0.1"
	app.Name = "esdb"
	app.Usage = "参数"
	app.UsageText = "本程序专门为同步es使用"
	app.ArgsUsage = ``

	app.Email = "qqc88.abo@gmail.com"
	app.Author = "abo"

	//这个方法就是这个命令已启动会运行什么
	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c) //这个是打印app的help界面
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:    "set",
			Aliases: []string{"-set"},
			Usage: `{表名} {年-月-日} 
					--表名可选[movie, movie_actor,movie_director,movie_film_companies,movie_label,movie_series,piece]
					`,
			Action: func(c *cli.Context) error {
				//参数1
				do := c.Args().First()
				//参数2
				t := c.Args().Get(1)

				if len(t) < 8 {
					fmt.Println("请在表名后面正确输入年月日，比如2021-10-05")
				}

				if len(pool.IndexMap[do]) > 1 {
					pool.Settime(do, t)
					return nil
				}

				fmt.Println("请使用-h查看帮助")

				return nil
			},
		},
		{
			Name:    "del",
			Aliases: []string{"-del"},
			Usage: `{表名}
					--表名可选[movie, movie_actor,movie_director,movie_film_companies,movie_label,movie_series,piece]
					`,
			Action: func(c *cli.Context) error {
				//参数1
				do := c.Args().First()

				if len(pool.IndexMap[do]) > 1 {
					pool.DelIndex(do)
					return nil
				}

				fmt.Println("请使用-h查看帮助")

				return nil
			},
		},
		{
			Name:    "test",
			Aliases: []string{"-t"},
			Usage:   `测试elasticsearch的连接情况`,
			Action: func(c *cli.Context) error {
				pool.TestIndex()
				return nil
			},
		},
	}

	app.Run(os.Args)
}
