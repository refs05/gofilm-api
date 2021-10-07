package users

import (
	"gofilm/bussinesses/users"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Address   string
	FirstName string
	LastName  string
	Age       int
	NoHp      string
	Email     string `gorm:"unique"`
	Password  string 
}

func (rec *Users) toDomain() users.User {
	return users.User{
		Id:        int(rec.ID),
		Address:   rec.Address,
		FirstName: rec.FirstName,
		LastName:  rec.LastName,
		Age:       rec.Age,
		NoHp:      rec.NoHp,
		Email:     rec.Email,
		Password:  rec.Password,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(userDomain users.User) *Users {
	return &Users{
		Address:   userDomain.Address,
		FirstName: userDomain.FirstName,
		LastName:  userDomain.LastName,
		Age:       userDomain.Age,
		NoHp:      userDomain.NoHp,
		Email:     userDomain.Email,
		Password:  userDomain.Password,
	}
}
