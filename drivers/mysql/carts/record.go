package carts

import (
	"gofilm/bussinesses/carts"
	"time"
)

type Carts struct {
	Id     int `gorm:"primaryKey"`
	Total  int
	Is_pay bool
	Films  []Film `gorm:"many2many:cart_films;foreignKey:Id"`
	UserId int
}

type Film struct {
	Id            int `gorm:"unique;primaryKey"`
	Title         string
	Description   string
	ReleaseDate   string
	Rating        float32
	Price         int
	Adult         bool
	LanguagesKode string
}

type User struct {
	Id        int
	Address   string
	FirstName string
	LastName  string
	Age       int
	NoHp      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (rec *Carts) toDomain() carts.Cart {
	var filmToDomain []carts.Film
	for i := 0; i < len(rec.Films); i++ {
		filmToDomain = append(filmToDomain, carts.Film(rec.Films[i]))
	}
	return carts.Cart{
		Id:     rec.Id,
		Total:  rec.Total,
		Is_pay: rec.Is_pay,
		Films:  filmToDomain,
		UserID: rec.UserId,
	}
}

func fromDomain(cartDomain carts.Cart) *Carts {
	var filmFromDomain []Film
	for i := 0; i < len(cartDomain.Films); i++ {
		filmFromDomain = append(filmFromDomain, Film(cartDomain.Films[i]))
	}
	return &Carts{
		Id:     cartDomain.Id,
		Total:  cartDomain.Total,
		Is_pay: cartDomain.Is_pay,
		Films:  filmFromDomain,
		UserId: cartDomain.UserID,
	}
}
