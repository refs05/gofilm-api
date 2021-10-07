package films

import (
	"gofilm/bussinesses/films"
)

type Response struct {
	Films []Film `json:"results"`
}

type Film struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"overview"`
	ReleaseDate string  `json:"release_date"`
	Rating      float32 `json:"vote_average"`
	Adult       bool    `json:"adult"`
	LanguagesKode   string  `json:"original_language"`
	Genres      []int `json:"genre_ids"`
}

type Genre struct {
	Id int 
	Name string
}

func ToDomain(film Film) films.Convert {
	return films.Convert{
		Id:          film.Id,
		Title:       film.Title,
		Description: film.Description,
		ReleaseDate: film.ReleaseDate,
		Rating:      film.Rating,
		Adult:       film.Adult,
		LanguagesKode:   film.LanguagesKode,
		Genres:      film.Genres,
	}
}
