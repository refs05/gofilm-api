package collections

import (
	"gofilm/bussinesses/collections"
	"gofilm/controllers/collections/request"
	"gofilm/controllers/collections/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceCollection collections.CollectionUseCase
}

func NewHandler(collectionServ collections.CollectionUseCase) *Presenter {
	return &Presenter{
		serviceCollection: collectionServ,
	}
}

func (handler *Presenter) GetByUserID(e echo.Context) error {
	var req request.Collection

	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, "Error Noob")
	}

	resp, err := handler.serviceCollection.GetCollection(req.UserID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Error Handler")
	}

	return e.JSON(http.StatusOK, response.FromDomain(*resp))
}