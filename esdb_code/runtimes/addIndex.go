package runtimes

import (
	"esdb/es"
	"runtime"
	"strconv"
	"sync"

	"github.com/beego/beego/v2/core/logs"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

//定义一个并发的线程池通道，用来约束发送请求的并发数量
var chans chan string

func init() {
	maxThreads, _ := strconv.Atoi(os.Getenv("MAX_THREADS"))
	chans = make(chan string, maxThreads)
}

type DoRun struct {
	Count     int        //线程完成的计数器
	Err       bool       //错误状态
	countLock sync.Mutex //并发锁
}

func (this *DoRun) process(sId, indexName, txt string) {

	//触发异步来更新数据
	indexID := es.AddIndex(indexName, sId, txt)

	//如果写入失败了，这里再写入一次
	if len(indexID) < 1 {
		for i := 0; i < 5; i++ {
			indexID = es.AddIndex(indexName, sId, txt)
			logs.Error("初次写入失败，第", i+1, "重试", indexName, indexID)

			//写入成功，跳出
			if len(indexID) > 0 {
				break
			}
		}
	}

	//异常情况处理,将线程状态变成异常
	if len(indexID) < 1 {
		this.Err = true
		logs.Error("5次尝试后，最终写入失败", indexName, sId)
	}

	this.countLock.Lock()
	defer this.countLock.Unlock()
	//更新计数器,防止并发计数，这里使用线程锁
	this.Count = this.Count + 1
}

func (this *DoRun) handle(sId, indexName, txt string) {

	this.process(sId, indexName, txt)
	// 信号完成：开始启用下一个请求

	// 将缓冲区释放一个容量
	<-chans
}

func (this *DoRun) Work(indexName, sId, txt string) {
	// 当通道已满的时候将被阻塞
	// 所以停在这里等待，直到有容量（被释放），才能继续去处理请求

	//开启线程，占用一个缓冲区容量
	chans <- sId

	//对象赋值
	go this.handle(sId, indexName, txt)

	logs.Debug(("线程数量 runing: %d, 索引 %s id %s 进入线程池 \n"), runtime.NumGoroutine(), indexName, sId)
}
