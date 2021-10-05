package films

import "gofilm/bussinesses/films"

type Response struct {
	Films []Film `json:"results"`
}

type Film struct {
	Id          int     `json:"id"`
	Title       string  `json:"original_title"`
	Description string  `json:"overview"`
	ReleaseDate string  `json:"release_date"`
	Rating      float32 `json:"vote_average"`
	Adult       bool    `json:"adult"`
	Languages   string  `json:"original_language"`
	Genres      []Genre `json:"genre_ids"`
}

// type Language struct {
// 	Kode string `json:"original_language"`
// }

type Genre struct {
	Id int
	Name string
}

func ToDomain(film Film) films.Film {
	var genreToDomain []films.Genre
	for i := 0; i < len(film.Genres); i++ {
		genreToDomain = append(genreToDomain, films.Genre(film.Genres[i]))
	}
	return films.Film{
		Id:          film.Id,
		Title:       film.Title,
		Description: film.Description,
		ReleaseDate: film.ReleaseDate,
		Rating:      film.Rating,
		Adult:       film.Adult,
		Languages:   film.Languages,
		Genres:      genreToDomain,
	}
}
