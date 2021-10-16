package response

import "gofilm/bussinesses/collections"

type Collection struct {
	Id     int
	UserID int
	Films   []Film
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

func FromDomain(domain collections.Collection) Collection {
	var filmFromDomain []Film
	for i := 0; i < len(domain.Films); i++ {
		
		filmFromDomain = append(filmFromDomain, Film(domain.Films[i]))
	}
	return Collection{
		Id: domain.Id,
		UserID: domain.UserID,
		Films: filmFromDomain,
	}
}