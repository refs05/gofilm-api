package genres

import (
	"gofilm/bussinesses/genres"
	"gofilm/controllers/genres/response"
	"net/http"
	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceGenre genres.GenreUseCase
}

func NewHandler(genreServ genres.GenreUseCase) *Presenter {
	return &Presenter{
		serviceGenre: genreServ,
	}
}

func (handler *Presenter) GetGenres(c echo.Context) error {
	genres, err := handler.serviceGenre.GetGenres()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return c.JSON(http.StatusOK, response.NewResponseArray(*genres))
}