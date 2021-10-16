package collections

import "gofilm/bussinesses/collections"

type Collections struct {
	Id int
	UserID int
	Films []Film `gorm:"foreignKey:id"`
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

func (rec *Collections) toDomain() collections.Collection {
	var filmToDomain []collections.Film
	for i := 0; i < len(rec.Films); i++ {
		filmToDomain = append(filmToDomain, collections.Film(rec.Films[i]))
	}
	return collections.Collection{
		Id: rec.Id,
		UserID: rec.UserID,
		Films: filmToDomain,
	}
}

func fromDomain(collectionDomain collections.Collection) *Collections {
	var filmFromDomain []Film
	for i := 0; i < len(collectionDomain.Films); i++ {
		filmFromDomain = append(filmFromDomain, Film(collectionDomain.Films[i]))
	}
	return &Collections{
		Id: collectionDomain.Id,
		UserID: collectionDomain.UserID,
		Films: filmFromDomain,
	}
}

