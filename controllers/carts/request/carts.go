package request

import "gofilm/bussinesses/carts"

type AddFilm struct {
	FilmID []int `json:"film_id"`
	UserID int `json:"user_id"`
}

func ToDomain(request AddFilm) *carts.Convert {
	return &carts.Convert{
		UserID: request.UserID,
		Films: request.FilmID,
	}
}