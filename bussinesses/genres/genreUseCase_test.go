package genres_test

import (
	"errors"
	"gofilm/bussinesses/genres"
	genreMock "gofilm/bussinesses/genres/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	genreRepository genreMock.GenreRepository
	genreUseCase genres.GenreUseCase
	genreThird genreMock.GenreFromAPI
)

func setup() {
	genreUseCase = genres.NewService(&genreRepository, &genreThird)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetGenreById(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		domain := genres.Genre{
			Id: 1,
			Name: "Horror",
		}
		genreRepository.On("GetGenreById", mock.AnythingOfType("int")).Return(&domain, nil).Once()

		result, err := genreUseCase.GetGenreById(12)

		assert.Nil(t, err)
		assert.Equal(t, &domain, result)
	})

	t.Run("repository error | test case 2", func(t *testing.T) {
		errNotFound := errors.New("Repo ID Not Found")
		genreRepository.On("GetGenreById", mock.AnythingOfType("int")).Return(&genres.Genre{}, errNotFound).Once()
		result, err := genreUseCase.GetGenreById(23)

		assert.Equal(t, result, &genres.Genre{})
		assert.Equal(t, err, errNotFound)
	})
}

func TestGetGenres(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		domain := genres.Genre{
			Id: 1,
			Name: "Horror",
		}
		fromApi := []genres.Genre{}
		genreThird.On("GetGenreFromAPI").Return(&fromApi).Once()
		//for i := 0; i < len(fromApi); i++ {
			genreRepository.On("StoreGenre", mock.AnythingOfType("*genres.Genre")).Return(&domain, nil).Once()
		//}
		genreRepository.On("GetGenres").Return(&fromApi, nil)
		result, err := genreUseCase.GetGenres()

		assert.Nil(t, err)
		assert.Equal(t, &fromApi, result)
	})
}