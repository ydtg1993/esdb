package pool

import (
	"esdb/es"
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
