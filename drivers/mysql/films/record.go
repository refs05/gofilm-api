package films

import (
	"gofilm/bussinesses/films"

)

type Films struct {
	Id            int `gorm:"primaryKey"`
	Title         string
	Description   string
	ReleaseDate   string
	Rating        float32
	Price         int
	Adult         bool
	LanguagesKode string `gorm:"index"`
	Language      Languages `gorm:"foreignKey:LanguagesKode"`
	Genres        []Genres `gorm:"many2many:film_genres;foreignKey:Id"`
}

type Genres struct {
	Id  int `gorm:"unique;primaryKey"`
	Name string
}

type Languages struct {
	Kode string `gorm:"primaryKey"`
	Name string
}

func fromDomain(filmDomain films.Film) *Films {
	
	var genreFromDomain []Genres
	for i := 0; i < len(filmDomain.Genres); i++ {
		
		genreFromDomain = append(genreFromDomain, Genres(filmDomain.Genres[i]))
	}
	return &Films{
		Id:            filmDomain.Id,
		Title:         filmDomain.Title,
		Description:   filmDomain.Description,
		ReleaseDate:   filmDomain.ReleaseDate,
		Rating:        filmDomain.Rating,
		Price:         filmDomain.Price,
		Adult:         filmDomain.Adult,
		LanguagesKode: filmDomain.LanguagesKode,
		Genres:       genreFromDomain,
	}
}

func (rec *Films) toDomain() films.Film {
	var genreToDomain []films.Genre
	for i := 0; i < len(rec.Genres); i++ {
		
		genreToDomain = append(genreToDomain, films.Genre(rec.Genres[i]))
	}
	return films.Film{
		Id:          rec.Id,
		Title:       rec.Title,
		Description: rec.Description,
		ReleaseDate: rec.ReleaseDate,
		Rating:      rec.Rating,
		Price:       rec.Price,
		Adult:       rec.Adult,
		LanguagesKode:   rec.LanguagesKode,
		Genres:      genreToDomain,
	}
}
