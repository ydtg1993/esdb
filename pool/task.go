package pool

import (
	"time"
)

//每5分钟
func TaskActor() {
	d := time.Duration(time.Minute * 2)
	t := time.NewTicker(d)
	defer t.Stop()

	//启动更新数据到es
	ActorsDo()

	for {
		<-t.C
		ActorsDo()
	}
}

//每5分钟
func TaskLabel() {
	d := time.Duration(time.Minute * 2)
	t := time.NewTicker(d)
	defer t.Stop()

	//启动更新数据到es
	LabelDo()

	for {
		<-t.C
		LabelDo()
	}
}

//每5分钟
func TaskSeries() {
	d := time.Duration(time.Minute * 2)
	t := time.NewTicker(d)
	defer t.Stop()

	//启动更新数据到es
	SeriesDo()

	for {
		<-t.C
		SeriesDo()
	}
}

//每5分钟
func TaskFilm() {
	d := time.Duration(time.Minute * 2)
	t := time.NewTicker(d)
	defer t.Stop()

	//启动更新数据到es
	FilmDo()

	for {
		<-t.C
		FilmDo()
	}
}

//每5分钟
func TaskCategory() {
	d := time.Duration(time.Minute * 2)
	t := time.NewTicker(d)
	defer t.Stop()

	//启动更新数据到es
	CategoryDo()

	for {
		<-t.C
		CategoryDo()
	}
}

//每5分钟
func TaskDirector() {
	d := time.Duration(time.Minute * 2)
	t := time.NewTicker(d)
	defer t.Stop()

	//启动更新数据到es
	DirectorDo()

	for {
		<-t.C
		DirectorDo()
	}
}

//每5分钟
func TaskMovie() {
	d := time.Duration(time.Minute * 2)
	t := time.NewTicker(d)
	defer t.Stop()

	//启动更新数据到es
	MovieDo()

	for {
		<-t.C
		MovieDo()
	}
}

//每5分钟
func TaskPiece() {
	d := time.Duration(time.Minute * 2)
	t := time.NewTicker(d)
	defer t.Stop()

	//启动更新数据到es
	PieceDo()

	for {
		<-t.C
		PieceDo()
	}
}
