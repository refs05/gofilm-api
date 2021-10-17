package collections_test

import (
	"errors"
	"gofilm/bussinesses"
	"gofilm/bussinesses/carts"
	"gofilm/bussinesses/collections"
	collecMock "gofilm/bussinesses/collections/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	collecRepository collecMock.CollectionRepository
	collecUseCase collections.CollectionUseCase
	collecCart carts.CartUseCase
)

func setup() {
	collecUseCase = collections.NewService(&collecRepository, collecCart)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetCollection(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		domain := collections.Collection{
			Id: 1,
			Films: []collections.Film{},
			UserID: 23,
		}

		collecRepository.On("GetCollection", mock.AnythingOfType("int")).Return(&domain, nil).Once()

		result, err := collecUseCase.GetCollection(1)

		assert.Nil(t, err)
		assert.Equal(t, &domain, result)
	})

	t.Run("repository error | test case 2", func(t *testing.T) {
		errNotFound := errors.New("Repo ID Not Found")
		collecRepository.On("GetCollection", mock.AnythingOfType("int")).Return(&collections.Collection{}, errNotFound).Once()
		result, err := collecUseCase.GetCollection(1)

		assert.Equal(t, result, &collections.Collection{})
		assert.Equal(t, err, errNotFound)
	})

	t.Run("invalid id | test case 3", func(t *testing.T) {
		result, err := collecUseCase.GetCollection(-1)

		assert.Equal(t, result, &collections.Collection{})
		assert.Equal(t, err, bussinesses.ErrIDNotFound)
	})
}

