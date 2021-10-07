package users_test

import (
	"errors"
	"gofilm/app/middleware"
	"gofilm/bussinesses"
	"gofilm/bussinesses/users"
	userMock "gofilm/bussinesses/users/mocks"
	"os"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepository userMock.UserRepository
	userUseCase users.UserUseCase

)

func setup() {
	userUseCase = users.NewService(&userRepository, &middleware.ConfigJWT{})
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetByID(t *testing.T) {
	t.Run("valid test | test case 1", func(t *testing.T) {
		domain := users.User{
			Id: 1,
			Address: "Jakarta",   
			FirstName: "Restu",
			LastName: "Fajar",
			Age: 34,
			NoHp: "089654456543",  
			Email: "restudsdsd#gmail.com",
			Password: "hahaha",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: time.Time{},
		}
		userRepository.On("GetByID", mock.AnythingOfType("int")).Return(&domain, nil).Once()

		result, err := userUseCase.GetByID(1)

		assert.Nil(t, err)
		assert.Equal(t, domain.FirstName, result.FirstName)
	})

	t.Run("invalid id | test case 2", func(t *testing.T) {
		result, err := userUseCase.GetByID(-1)

		assert.Equal(t, result, &users.User{})
		assert.Equal(t, err, bussinesses.ErrIDNotFound)
	})

	t.Run("repository error | test case 3", func(t *testing.T) {
		errNotFound := errors.New("Repo ID Not Found")
		userRepository.On("GetByID", mock.AnythingOfType("int")).Return(&users.User{}, errNotFound).Once()
		result, err := userUseCase.GetByID(10)

		assert.Equal(t, result, &users.User{})
		assert.Equal(t, err, errNotFound)
	})
}

//