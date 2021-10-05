package actors

import (
	"gofilm/bussinesses/actors"
	"gofilm/controllers/actors/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceActor actors.ActorUserCase
}

func NewHandler(actorServ actors.ActorUserCase) *Presenter {
	return &Presenter{
		serviceActor: actorServ,
	}
}

func (handler *Presenter) GetActors(c echo.Context) error {
	actors, err := handler.serviceActor.GetActor()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "err",
			"messages": err, 
		})
	}
	return c.JSON(http.StatusOK, response.NewResponseArray(*actors))
}