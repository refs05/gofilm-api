package languages

import (
	"gofilm/bussinesses/languages"
	"gofilm/controllers/languages/response"
	"net/http"
	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceLang languages.LangUseCase
}

func NewHandler(langServ languages.LangUseCase) *Presenter {
	return &Presenter{
		serviceLang: langServ,
	}
}

func (handler *Presenter) GetLanguages(c echo.Context) error {
	langs, err := handler.serviceLang.GetLangs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return c.JSON(http.StatusOK, response.NewResponseArray(*langs))
}