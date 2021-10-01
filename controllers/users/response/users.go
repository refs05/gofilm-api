package response

import (
	"gofilm/bussinesses/users"
	"time"
)

type Users struct {
	Id        int       `json:"id"`
	Address   string    `json:"address"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Age       int       `json:"age"`
	NoHp      string    `json:"nohp"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func FromDomain(domain users.User) Users {
	return Users{
		Id:        domain.Id,
		Address:   domain.Address,
		FirstName: domain.FirstName,
		LastName:  domain.LastName,
		Age:       domain.Age,
		NoHp:      domain.NoHp,
		Email:     domain.Email,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
