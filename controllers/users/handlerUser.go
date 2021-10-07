package users

import (
	//"fmt"
	//"gofilm/bussinesses/genres"
	"gofilm/bussinesses/users"
	"gofilm/controllers"
	"gofilm/controllers/users/request"
	"gofilm/controllers/users/response"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceUser users.UserUseCase
}

func NewHandler(userServ users.UserUseCase) *Presenter {
	return &Presenter{
		serviceUser: userServ,
	}
}

func (handler *Presenter) Insert(e echo.Context) error {
	var req request.AddUser

	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, "Error Noob")
	}

	domain := request.ToDomain(req)

	resp, err := handler.serviceUser.Store(domain)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Error Disini tau")
	}

	return e.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *Presenter) Update(e echo.Context) error {
	var req request.AddUser

	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, "Error Noob")
	}

	domain := request.ToDomain(req)
	id, _ := strconv.Atoi(e.QueryParam("id"))

	resp, err := handler.serviceUser.Update(id, domain)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Error Disini tau")
	}

	return e.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *Presenter) Delete(e echo.Context) error {
	id, _ := strconv.Atoi(e.QueryParam("id"))

	err := handler.serviceUser.Delete(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Error Handler")
	}

	return e.JSON(http.StatusOK, "Delete Succes")
}

func (handler *Presenter) GetByID(e echo.Context) error {
	id, _ := strconv.Atoi(e.QueryParam("id"))

	resp, err := handler.serviceUser.GetByID(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Error Handler Jancok")
	}

	return e.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *Presenter) CreateToken(c echo.Context) error {
	
	var userReq users.User
	c.Bind(&userReq)
	

	token, err := handler.serviceUser.CreateToken(userReq.Email, userReq.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := struct {
		Token string `json:"token"`
	}{Token: token}
	
	return controllers.NewSuccessResponse(c, response)
}

