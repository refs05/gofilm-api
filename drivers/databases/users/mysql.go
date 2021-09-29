package users

import (
	//"context"
	"gofilm/bussinesses/users"

	"gorm.io/gorm"
)

type mysqlUsersRepository struct {
	DB *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) users.UserRepository {
	return &mysqlUsersRepository{
		DB: db,
	}
}

func (nr *mysqlUsersRepository) GetByID(Id int) (*users.User, error) {
	user := users.User{}
	err := nr.DB.Where("id = ?", Id).First(&user).Error
	if err != nil {
		return &users.User{}, err
	}
	return &user, nil
}
	
func (nr *mysqlUsersRepository) Update(id int, user *users.User) (*users.User, error) {
	
	err := nr.DB.Model(&user).Updates(&user).Error;
	if err != nil {
		return &users.User{}, err
	}
	return user, err
}

func (nr *mysqlUsersRepository) Store(user *users.User) (*users.User, error) {
	return user, nil
}

func (nr *mysqlUsersRepository) Delete(Id int, user *users.User) (*users.User, error) {
	return user, nil
}