package data

import "time"

// Movie - структура обхекта для фильмов
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // не будет показано
	Title     string    `json:"title"`
	Year      int       `json:"year, omitempty"` // если значение не задано то мы просто вывелём 0
	Runtime   int32     `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,o,omitempty"`
	Version   int32     `json:"version"`
}
