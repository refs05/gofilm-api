package users

import (
	//"context"
	"time"
)

type User struct {
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

type UserUseCase interface {
	GetByID(id int) (*User, error)
	Update(user *User, id int) (*User, error)
	Store(user *User) (*User, error)
	Delete(id int, user *User) (*User, error)
}

type UserRepository interface {
	GetByID(id int) (*User, error)
	Update(id int, user *User) (*User, error)
	Store(user *User) (*User, error)
	Delete(id int, user *User) (*User, error)
}