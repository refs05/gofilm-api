package response

import (
	"gofilm/bussinesses/films"
)

type Films struct {
	Id          int
	Title       string
	Description string
	ReleaseDate string
	Rating      float32
	Price       int
	Adult       bool
	Genres      []int
	Languages   string
}

func ResponseFilms(domain films.Film) Films {
	var res []int
	for i := 0; i < len(domain.Genres); i++ {
		res = append(res, domain.Genres[i].Id)
	}
	return Films{
		Id:          domain.Id,
		Title:       domain.Title,
		Description: domain.Description,
		ReleaseDate: domain.ReleaseDate,
		Rating:      domain.Rating,
		Price:       domain.Price,
		Adult:       domain.Adult,
		Genres:      res,
		Languages:   domain.LanguagesKode,
	}
}

func NewResponseArray(domainFilms []films.Film) []Films {
	var resp []Films

	for _, value := range domainFilms {
		resp = append(resp, ResponseFilms(value))
	}

	return resp
}
