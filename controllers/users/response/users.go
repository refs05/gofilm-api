package response

import (
	"gofilm/bussinesses/users"
	"time"
)

type Users struct {
	Id        int
	Address   string
	FirstName string
	LastName  string
	Age       int
	NoHp      int
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDomain(domain users.User) Users {
	return Users{
		Id: domain.Id,
		Address: domain.Address,
		FirstName: domain.FirstName,
		LastName: domain.LastName,
		Age: domain.Age,
		NoHp: domain.NoHp,
		Email: domain.Email,
		Password: domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}