package response

import "gofilm/bussinesses/carts"

type Cart struct {
	Id int `json:"id"`
	FilmID []int `json:"film_id"`
	UserID int `json:"user_id"`
	Total int `json:"total"`
	Is_pay bool `json:"is_pay"`
}

type Film struct {
	Id          int
	Title       string
	Description string
	ReleaseDate string
	Rating      float32
	Price       int
	Adult       bool
	LanguagesKode   string
}

func FromDomain(domain carts.Cart) Cart {
	var filmFromDomain []int
	for i := 0; i < len(domain.Films); i++ {
		
		filmFromDomain = append(filmFromDomain, int(domain.Films[i].Id))
	}
	return Cart{
		Id: domain.Id,
		FilmID: filmFromDomain,
		UserID: domain.UserID,
		Total: domain.Total,
		Is_pay: domain.Is_pay,
	}
}