package users

import (
	"gofilm/app/middleware"
	"gofilm/bussinesses"
	//"strings"
)

type serviceUsers struct {
	repository UserRepository
	jwtAuth *middleware.ConfigJWT
}

func NewService(repoUser UserRepository, jwtauth *middleware.ConfigJWT) UserUseCase {
	return &serviceUsers {
		repository: repoUser,
		jwtAuth: jwtauth,
	}
}

func (caseUser *serviceUsers) CreateToken(email, password string) (string, error) {
	
	userDomain, err := caseUser.repository.GetValidUser(email, password)
	if err != nil {
		return "", err
	}
	
	token := caseUser.jwtAuth.GenerateToken(userDomain.Id)
	return token, nil
}

func (caseUser *serviceUsers) GetByID(id int) (*User, error) {
	if id <= 0 {
		return &User{}, bussinesses.ErrIDNotFound
	}

	result, err :=  caseUser.repository.GetByID(id)

	if err != nil {
		return &User{}, err
	}

	return result, nil
}

func (caseUser *serviceUsers) Update(id int, user *User) (*User, error) {
	result, err := caseUser.repository.Update(id, user)
	if err != nil {
		return &User{}, err
	}
	return result, nil
}

func (caseUser *serviceUsers) Store(user *User) (*User, error) {
	result, err := caseUser.repository.Store(user)
	if err != nil {
		return &User{}, err
	}
	return result, nil
}

func (caseUser *serviceUsers) Delete(id int) error {
	
	err := caseUser.repository.Delete(id)
	if err != nil {
		return err
	}
	// if existedUser == (User{}) {
	// 	return bussinesses.ErrNotFound
	// }

	return nil
}