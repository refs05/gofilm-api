package users

import (

)

type serviceUsers struct {
	repository UserRepository
}

func NewService(repoUser UserRepository) UserUseCase {
	return &serviceUsers {
		repository: repoUser,
	}
}

func (caseUser *serviceUsers) GetByID(id int) (*User, error) {
	return &User{}, nil
}

func (caseUser *serviceUsers) Update(user *User, id int) (*User, error) {
	return &User{}, nil
}

func (caseUser *serviceUsers) Store(user *User) (*User, error) {
	result, err := caseUser.repository.Store(user)
	if err != nil {
		return &User{}, err
	}
	return result, nil
}

func (caseUser *serviceUsers) Delete(id int, user *User) (*User, error) {
	return &User{}, nil
}