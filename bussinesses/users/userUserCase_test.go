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

func TestUpdate(t *testing.T) {
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
		userRepository.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("*users.User")).Return(&domain, nil).Once()
		
		result, err := userUseCase.Update(1, &domain)

		assert.Nil(t, err)
		assert.Equal(t, &domain, result)
	})

	t.Run("repository error | test case 2", func(t *testing.T) {
		errNotFound := errors.New("Repo ID Not Found")
		userRepository.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("*users.User")).Return(&users.User{}, errNotFound).Once()
		result, err := userUseCase.Update(1, &users.User{})

		assert.Equal(t, result, &users.User{})
		assert.Equal(t, err, errNotFound)
	})
}

func TestStore(t *testing.T) {
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
		userRepository.On("Store", mock.AnythingOfType("*users.User")).Return(&domain, nil).Once()

		result, err := userUseCase.Store(&domain)

		assert.Nil(t, err)
		assert.Equal(t, &domain, result)
	})

	t.Run("repository error | test case 2", func(t *testing.T) {
		errNotFound := errors.New("Repo ID Not Found")
		userRepository.On("Store", mock.AnythingOfType("*users.User")).Return(&users.User{}, errNotFound).Once()
		result, err := userUseCase.Store(&users.User{})

		assert.Equal(t, result, &users.User{})
		assert.Equal(t, err, errNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("valid test | test case 1", func(t *testing.T) {
		userRepository.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()
		err := userUseCase.Delete(1)
		assert.Nil(t, err)
	})

	t.Run("repository error | test case 2", func(t *testing.T) {
		errNotFound := errors.New("Repo ID Not Found")
		userRepository.On("Delete", mock.AnythingOfType("int")).Return(errNotFound).Once()
		err := userUseCase.Delete(1)
		assert.Equal(t, err, errNotFound)
	})
}

func TestCreateToken(t *testing.T) {
	t.Run("valid test | test case 1", func(t *testing.T) {
		
		userRepository.On("GetValidUser", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&users.User{}, nil).Once()
		_, err := userUseCase.CreateToken("restu@gmail.com", "halo123")
	
		assert.Nil(t, err)
	})

	t.Run("repository error | test case 2", func(t *testing.T) {
		errNotFound := errors.New("Repo ID Not Found")
		userRepository.On("GetValidUser", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&users.User{}, errNotFound).Once()
		_, err := userUseCase.CreateToken("restu@gmail.com", "halo123")
		assert.Equal(t, "", "")
		assert.Equal(t, err, errNotFound)
	})

	
}

