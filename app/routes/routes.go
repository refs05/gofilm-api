package routes

import (
	"gofilm/controllers/actors"
	"gofilm/controllers/films"
	"gofilm/controllers/genres"
	"gofilm/controllers/languages"
	"gofilm/controllers/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	JWTMiddleware middleware.JWTConfig
	UserHandler users.Presenter
	GenreHandler genres.Presenter
	ActorHandler actors.Presenter
	LanguageHandler languages.Presenter
	FilmHandler films.Presenter
}

func (handler *HandlerList) RouteUser(e *echo.Echo) {
	users := e.Group("/users", middleware.JWTWithConfig(handler.JWTMiddleware))
	users.POST("/register", handler.UserHandler.Insert)
	users.POST("/update/", handler.UserHandler.Update)
	users.DELETE("/delete/", handler.UserHandler.Delete)
	users.GET("/", handler.UserHandler.GetByID)

	auth := e.Group("/")
	auth.POST("login", handler.UserHandler.CreateToken)

	genres := e.Group("/genres")
	genres.GET("", handler.GenreHandler.GetGenres)

	langs := e.Group("/languages")
	langs.GET("", handler.LanguageHandler.GetLanguages)

	films := e.Group("/toprated")
	films.GET("", handler.FilmHandler.GetFilms)
}

