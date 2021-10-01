package routes

import (
	//"gofilm/app/middleware"
	"gofilm/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	JWTMiddleware middleware.JWTConfig
	UserHandler users.Presenter
}

func (handler *HandlerList) RouteUser(e *echo.Echo) {
	users := e.Group("/users", middleware.JWTWithConfig(handler.JWTMiddleware))
	
	users.POST("/register", handler.UserHandler.Insert)
	users.POST("/update/:id", handler.UserHandler.Update)
	users.DELETE("/delete/:id", handler.UserHandler.Delete)
	users.GET("/:id", handler.UserHandler.GetByID)

	auth := e.Group("/")
	
	auth.POST("login", handler.UserHandler.CreateToken)
}

