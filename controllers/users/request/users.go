package request

import (
	"gofilm/bussinesses/users"
	//"github.com/golang-jwt/jwt/request"
)

type AddUser struct {
	Address   string `json:"address"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	NoHp      string `json:"nohp"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateUser struct {
}

func ToDomain(request AddUser) *users.User {
	return &users.User{
		Address:   request.Address,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Age:       request.Age,
		NoHp:      request.NoHp,
		Email:     request.Email,
		Password:  request.Password,
	}
}
