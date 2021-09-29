package request

import (
	"gofilm/bussinesses/users"

	//"github.com/golang-jwt/jwt/request"
)

type AddUser struct {
	Id int 
	Address string
	FirstName string
	LastName string
	Age int
	NoHp int
	Email string
	Password string
}

type UpdateUser struct{

}

func ToDomain(request AddUser) *users.User {
	return &users.User{
		Id: request.Id,
		Address: request.Address,
		FirstName: request.FirstName,
		LastName: request.LastName,
		Age: request.Age,
		NoHp: request.NoHp,
		Email: request.Email,
		Password: request.Password,
	}
}