package genres

import "gofilm/bussinesses/genres"

//import "gofilm/bussinesses/genres"

type Genres struct {
	Id int `gorm:"unique"`
	Name string
}

func fromDomain(genreDomain genres.Genre) *Genres {
	return &Genres{
		Id: genreDomain.Id,
		Name: genreDomain.Name,
	}
}

func (rec *Genres) toDomain() genres.Genre {
	return genres.Genre{
		Id: rec.Id,
		Name: rec.Name,
	}
}

