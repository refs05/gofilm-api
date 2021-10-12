package carts

import (
	"gofilm/bussinesses/carts"
	"gofilm/controllers/carts/request"
	"gofilm/controllers/carts/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceCart carts.CartUseCase
}

func NewHandler(cartServ carts.CartUseCase) *Presenter {
	return &Presenter{
		serviceCart: cartServ,
	}
}

func (handler *Presenter) Insert(e echo.Context) error {
	var req request.AddFilm

	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, "Error Noob")
	}

	domain := request.ToDomain(req)
	resp, err := handler.serviceCart.CreateCartUser(domain)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Server Internal Error")
	}

	return e.JSON(http.StatusOK, response.FromDomain(*resp))
}