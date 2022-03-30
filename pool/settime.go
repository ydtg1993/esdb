package pool

import (
	"esdb/es"
	"esdb/rd"
	"fmt"
)

//对应的redis健
var IndexMap = map[string]string{
	"movie":                "synch_es_movie",
	"movie_actor":          "synch_es_actor",
	"movie_category":       "synch_es_category",
	"movie_director":       "synch_es_director",
	"movie_film_companies": "synch_es_film_companies",
	"movie_label":          "synch_es_label",
	"movie_series":         "synch_es_series",
	"piece":                "synch_es_piece",
}

//对应的es对象
var EsMap = map[string]string{
	"movie":                es.MappingMovie,
	"movie_actor":          es.MappingMovieActor,
	"movie_category":       es.MappingMovieCategory,
	"movie_director":       es.MappingMovieDirector,
	"movie_film_companies": es.MappingMovieFilm,
	"movie_label":          es.MappingMovieLabel,
	"movie_series":         es.MappingMovieSeries,
	"piece":                es.MappingPiece,
}

/**
* 重新设置消息时间
 */
func Settime(name, date string) {

	//抢锁
	switch name {
	case "movie":
		movieLock.Lock()
		defer movieLock.Unlock()
	case "movie_actor":
		actorLock.Lock()
		defer actorLock.Unlock()
	case "movie_category":
		categoryLock.Lock()
		defer categoryLock.Unlock()
	case "movie_director":
		directorLock.Lock()
		defer directorLock.Unlock()
	case "movie_film_companies":
		filmLock.Lock()
		defer filmLock.Unlock()
	case "movie_label":
		labelLock.Lock()
		defer labelLock.Unlock()
	case "movie_series":
		seriesLock.Lock()
		defer seriesLock.Unlock()
	case "piece":
		pieceLock.Lock()
		defer pieceLock.Unlock()
	default:
		return
	}

	redisKey := IndexMap[name]

	rd.StringSet(redisKey, fmt.Sprintf("%s 00:00:00", date))

	fmt.Println(name, "完成设置时间节点：", rd.StringGet(redisKey))
}
