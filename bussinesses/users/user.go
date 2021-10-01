package users

import (
	//"context"
	//"gofilm/drivers/databases/users"
	"time"
)

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

type UserUseCase interface {
	CreateToken(email, password string) (string, error)
	GetByID(id int) (*User, error)
	Update(id int, user *User) (*User, error)
	Store(user *User) (*User, error)
	Delete(id int) error
}

type UserRepository interface {
	GetValidUser(email, password string) (*User, error)
	GetByID(id int) (*User, error)
	Update(id int, user *User) (*User, error)
	Store(user *User) (*User, error)
	Delete(id int) error
}
