package response

import "gofilm/bussinesses/actors"

type Actors struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Birthday string	`json:"birthday"`
}

func ResponseActors(domain actors.Actor) Actors {
	return Actors{
		Id: domain.Id,
		Name: domain.Name,
		Birthday: domain.Birthday,
	}
}

func NewResponseArray(domainActors []actors.Actor) []Actors {
	var resp []Actors

	for _, value := range domainActors {
		resp = append(resp, ResponseActors(value))
	}

	return resp
}