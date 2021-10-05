package main

import (
	"fmt"
	"gofilm/app/constant"
	"log"

	_userServ "gofilm/bussinesses/users"
	_userHandler "gofilm/controllers/users"
	_userRepo "gofilm/drivers/databases/users"

	_genreServ "gofilm/bussinesses/genres"
	_genreThird "gofilm/drivers/thirdparties/genres"
	_genreHandler "gofilm/controllers/genres"
	_genreRepo "gofilm/drivers/databases/genres"

	_actorServ "gofilm/bussinesses/actors"
	_actorHandler "gofilm/controllers/actors"
	_actorRepo "gofilm/drivers/databases/actors"

	_languageServ "gofilm/bussinesses/languages"
	_languageThird "gofilm/drivers/thirdparties/languages"
	_languageHandler "gofilm/controllers/languages"
	_languageRepo "gofilm/drivers/databases/languages"

	_filmsServ "gofilm/bussinesses/films"
	_filmsThird "gofilm/drivers/thirdparties/films"
	_filmsHandler "gofilm/controllers/films"
	_filmsRepo "gofilm/drivers/databases/films"

	_middleware "gofilm/app/middleware"
	_routes "gofilm/app/routes"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(status string) *gorm.DB {
	db := "gofilm"

	connectionString := fmt.Sprintf("root:BavarianRestu2002@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", db)

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&_userRepo.Users{},
		&_genreRepo.Genres{},
		&_actorRepo.Actors{},
		&_languageRepo.Languages{},
		&_filmsRepo.Films{},
	)

	return DB
}

func main() {
	fmt.Println("haloooo")
	db := InitDB("")
	e := echo.New()

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       constant.SecretJWT,
		ExpiresDuration: constant.ExpiresDuration,
	}

	usersRepo := _userRepo.NewMySQLRepo(db)
	usersService := _userServ.NewService(usersRepo, &configJWT)
	usersHandler := _userHandler.NewHandler(usersService)

	genresRepo := _genreRepo.NewMySQLRepo(db)
	genresThird := _genreThird.NewConsumeAPI()
	genresService := _genreServ.NewService(genresRepo, genresThird)
	genresHandler := _genreHandler.NewHandler(genresService)

	langsRepo := _languageRepo.NewMySQLRepo(db)
	langsThird := _languageThird.NewConsumeAPI()
	langsService := _languageServ.NewService(langsRepo, langsThird)
	langsHandler := _languageHandler.NewHandler(langsService) 

	filmsRepo := _filmsRepo.NewMySQLRepo(db)
	filmsThird := _filmsThird.NewConsumeAPI()
	filmsService := _filmsServ.NewService(filmsRepo, filmsThird)
	filmsHandler := _filmsHandler.NewHandler(filmsService)

	actorsRepo := _actorRepo.NewMySQLRepo(db)
	actorsService := _actorServ.NewService(actorsRepo)
	actorsHandler := _actorHandler.NewHandler(actorsService)

	routesInit := _routes.HandlerList{
		JWTMiddleware: configJWT.Init(),
		UserHandler:   *usersHandler,
		GenreHandler:  *genresHandler,
		ActorHandler: *actorsHandler,
		LanguageHandler: *langsHandler,
		FilmHandler: *filmsHandler,
	}

	routesInit.RouteUser(e)

	log.Fatal(e.Start(":8080"))
}



//pisah config db
//carts
//unittest
//films
//payment
//koleksi