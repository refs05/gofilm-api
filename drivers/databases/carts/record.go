package carts

import (
	"gofilm/bussinesses/carts"
	"time"
)

type Carts struct {
	Id     int
	Total  int
	Is_pay bool
	Films  []Film
	UserId int
	User   User
}

type Film struct {
	Id          int
	Title       string
	Description string
	ReleaseDate string
	Rating      float32
	Price       int
	Adult       bool
	Genres      []Genre
	Languages   string
	CartID      int
}

type Genre struct {
	Id   int
	Name string
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
	return carts.Cart{
		Id: rec.Id,
		Total: rec.Total,
		Is_pay: rec.Is_pay,
		//film[]
		//userId
	}
}

func fromDomain(cartDomain carts.Cart) *Carts {
	return &Carts{
		Id: cartDomain.Id,
		Total: cartDomain.Total,
		Is_pay: cartDomain.Is_pay,
		//film[]
		//userId
	}
}