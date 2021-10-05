package openapi

import "gofilm/bussinesses/genres"

type Response struct {
	Genres []Genre `json:"genres"`
}

type Genre struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func ToDomain(id int, name string) genres.Genre {
	return genres.Genre{
		Id:   id,
		Name: name,
	}
}
