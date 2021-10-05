package response

import "gofilm/bussinesses/genres"

type Genres struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ResponseGenres(domain genres.Genre) Genres {
	return Genres{
		Id: domain.Id,
		Name: domain.Name,
	}
}

func NewResponseArray(domainGenres []genres.Genre) []Genres {
	var resp []Genres

	for _, value := range domainGenres {
		resp = append(resp, ResponseGenres(value))
	}

	return resp
}