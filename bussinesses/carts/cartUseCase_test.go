package carts_test

import (
	"errors"
	"gofilm/bussinesses"
	"gofilm/bussinesses/carts"
	cartMock "gofilm/bussinesses/carts/mocks"
	"gofilm/bussinesses/films"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	cartRepository cartMock.CartRepository
	cartUseCase carts.CartUseCase
	cartFilm films.FilmUseCase
)

func setup() {
	cartUseCase = carts.NewService(&cartRepository, cartFilm)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestUpdateCartUser(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		domain := carts.Cart{
			Id: 1,
			Total: 100000,
			Is_pay: true,
			Films: []carts.Film{},
			UserID: 23,
		}
		cartRepository.On("UpdateCart", mock.AnythingOfType("int"), mock.AnythingOfType("*carts.Cart")).Return(&domain, nil).Once()

		result, err := cartUseCase.UpdateCartUser(1, &domain)

		assert.Nil(t, err)
		assert.Equal(t, &domain, result)
	})

	t.Run("repository error | test case 2", func(t *testing.T) {
		errNotFound := errors.New("Repo ID Not Found")
		cartRepository.On("UpdateCart", mock.AnythingOfType("int"), mock.AnythingOfType("*carts.Cart")).Return(&carts.Cart{}, errNotFound).Once()
		result, err := cartUseCase.UpdateCartUser(1, &carts.Cart{})

		assert.Equal(t, result, &carts.Cart{})
		assert.Equal(t, err, errNotFound)
	})
}

func TestGetCartByUser(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		domain := carts.Cart{
			Id: 1,
			Total: 100000,
			Is_pay: true,
			Films: []carts.Film{},
			UserID: 23,
		}

		cartRepository.On("GetCartByUser", mock.AnythingOfType("int")).Return(&domain, nil).Once()

		result, err := cartUseCase.GetCartByUser(1)

		assert.Nil(t, err)
		assert.Equal(t, &domain, result)
	})

	t.Run("repository error | test case 2", func(t *testing.T) {
		errNotFound := errors.New("Repo ID Not Found")
		cartRepository.On("GetCartByUser", mock.AnythingOfType("int")).Return(&carts.Cart{}, errNotFound).Once()
		result, err := cartUseCase.GetCartByUser(1)

		assert.Equal(t, result, &carts.Cart{})
		assert.Equal(t, err, errNotFound)
	})

	t.Run("invalid id | test case 3", func(t *testing.T) {
		result, err := cartUseCase.GetCartByUser(-1)

		assert.Equal(t, result, &carts.Cart{})
		assert.Equal(t, err, bussinesses.ErrIDNotFound)
	})
}


// func TestDeleteCartUser(t *testing.T) {
// 	t.Run("case 1", func(t *testing.T) {
// 		assert()
// 	})
// }

func TestGetCartByID(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		domain := carts.Cart{
			Id: 1,
			Total: 100000,
			Is_pay: true,
			Films: []carts.Film{},
			UserID: 23,
		}

		cartRepository.On("GetCartByUser", mock.AnythingOfType("int")).Return(&domain, nil).Once()

		result, err := cartUseCase.GetCartByID(1)

		assert.Nil(t, err)
		assert.Equal(t, &domain, result)
	})

	t.Run("repository error | test case 2", func(t *testing.T) {
		errNotFound := errors.New("Repo ID Not Found")
		cartRepository.On("GetCartByUser", mock.AnythingOfType("int")).Return(&carts.Cart{}, errNotFound).Once()
		result, err := cartUseCase.GetCartByID(1)

		assert.Equal(t, result, &carts.Cart{})
		assert.Equal(t, err, errNotFound)
	})

	t.Run("invalid id | test case 3", func(t *testing.T) {
		result, err := cartUseCase.GetCartByID(-1)

		assert.Equal(t, result, &carts.Cart{})
		assert.Equal(t, err, bussinesses.ErrIDNotFound)
	})
}

func TestChangeStatus(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		cartRepository.On("ChangeStatus", mock.AnythingOfType("int")).Return(nil).Once()

		err := cartUseCase.ChangeStatus(1)

		assert.Nil(t, err)
		//assert.Equal(t, &domain, result)
	})

	t.Run("repository error | test case 2", func(t *testing.T) {
		errNotFound := errors.New("Repo ID Not Found")
		cartRepository.On("ChangeStatus", mock.AnythingOfType("int")).Return(errNotFound).Once()
		err := cartUseCase.ChangeStatus(1)

		//assert.Equal(t, result, &carts.Cart{})
		assert.Equal(t, err, errNotFound)
	})

	t.Run("invalid id | test case 3", func(t *testing.T) {
		err := cartUseCase.ChangeStatus(-1)

		//assert.Equal(t, result, &carts.Cart{})
		assert.Equal(t, err, bussinesses.ErrIDNotFound)
	})
}

func TestGetFilmUser(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		domain := carts.Cart{
			Id: 1,
			Total: 100000,
			Is_pay: true,
			Films: []carts.Film{},
			UserID: 23,
		}

		cartRepository.On("GetFilmUser", mock.AnythingOfType("int")).Return(&domain, nil).Once()

		result, err := cartUseCase.GetFilmUser(1)

		assert.Nil(t, err)
		assert.Equal(t, &domain, result)
	})

	t.Run("repository error | test case 2", func(t *testing.T) {
		errNotFound := errors.New("Repo ID Not Found")
		cartRepository.On("GetFilmUser", mock.AnythingOfType("int")).Return(&carts.Cart{}, errNotFound).Once()
		result, err := cartUseCase.GetFilmUser(1)

		assert.Equal(t, result, &carts.Cart{})
		assert.Equal(t, err, errNotFound)
	})

	t.Run("invalid id | test case 3", func(t *testing.T) {
		result, err := cartUseCase.GetFilmUser(-1)

		assert.Equal(t, result, &carts.Cart{})
		assert.Equal(t, err, bussinesses.ErrIDNotFound)
	})
}