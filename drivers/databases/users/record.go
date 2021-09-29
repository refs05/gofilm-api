package users

import (
	"gofilm/bussinesses/users"
	"time"
)

type Users struct {
	Id int 
	Address string
	FirstName string
	LastName string
	Age int
	NoHp int
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Users) toDomain() users.User {
	return users.User{
		Id: rec.Id,
		Address: rec.Address,
		FirstName: rec.FirstName,
		LastName: rec.LastName,
		Age: rec.Age,
		NoHp: rec.NoHp,
		Email: rec.Email,
		Password: rec.Password,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(userDomain users.User) *Users {
	return &Users{
		Id: userDomain.Id,
		Address: userDomain.Address,
		FirstName: userDomain.FirstName,
		LastName: userDomain.LastName,
		Age: userDomain.Age,
		NoHp: userDomain.NoHp,
		Email: userDomain.Email,
		Password: userDomain.Password,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}