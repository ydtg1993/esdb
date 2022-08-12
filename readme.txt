2021-11-10更新说明

1，修改配置文件,线程数，根据服务器的压力情况（主要是tcp连接数），数字越大，并发越多，不要超过1000
vim  conf/app.conf

maxthreads = 300

2，系统引入了log方案（不要复制备注进去）
vim  conf/app.conf

logdays = 1        		//日志保存时间，不设置，默认7天（备注）
logpath = .   			//日志路径，需要/结尾，不设置默认当前目录（备注）
loglevel = error,debug	         //error表示记录错误，debug表示记录调试输出（备注）

3，删除重建movie索引
./esdb  del movie