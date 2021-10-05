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
	
	err := nr.DB.Model(&user).Where("id = ?", id).Updates(&user).Error;
	if err != nil {
		return &users.User{}, err
	}
	return user, err
}

func (nr *mysqlUsersRepository) Store(user *users.User) (*users.User, error) {
	rec := fromDomain(*user)

	result := nr.DB.Create(rec);

	if result.Error != nil {
		return user, result.Error
	}

	res := rec.toDomain()

	return &res, nil
}

func (nr *mysqlUsersRepository) Delete(Id int) error {
	user := Users{}

	user.ID = uint(Id)
	

	if err := nr.DB.Unscoped().Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (nr *mysqlUsersRepository) GetValidUser(email, password string) (*users.User, error) {
	user := users.User{}

	if err := nr.DB.Unscoped().Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return &users.User{}, err
	}

	return &user, nil
}