package films

import (
	"gofilm/bussinesses/films"
	"gofilm/controllers/films/response"
	"net/http"
	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceFilm films.FilmUseCase
}

func NewHandler(filmServ films.FilmUseCase) *Presenter {
	return &Presenter{
		serviceFilm: filmServ,
	}
}

func (handler *Presenter) GetFilms(c echo.Context) error {
	films, err := handler.serviceFilm.GetPopularFilms()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return c.JSON(http.StatusOK, response.NewResponseArray(*films))
}