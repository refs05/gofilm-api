package main

import (
	"log"

	_userServ "gofilm/bussinesses/users"
	_userHandler "gofilm/controllers/users"
	_userRepo "gofilm/drivers/mysql/users"

	_cartServ "gofilm/bussinesses/carts"
	_cartHandler "gofilm/controllers/carts"
	_cartRepo "gofilm/drivers/mysql/carts"

	_genreServ "gofilm/bussinesses/genres"
	_genreHandler "gofilm/controllers/genres"
	_genreRepo "gofilm/drivers/mysql/genres"
	_genreThird "gofilm/drivers/thirdparties/genres"

	_languageServ "gofilm/bussinesses/languages"
	_languageHandler "gofilm/controllers/languages"
	_languageRepo "gofilm/drivers/mysql/languages"
	_languageThird "gofilm/drivers/thirdparties/languages"

	_filmsServ "gofilm/bussinesses/films"
	_filmsHandler "gofilm/controllers/films"
	_filmsRepo "gofilm/drivers/mysql/films"
	_filmsThird "gofilm/drivers/thirdparties/films"

	_mysqlDriver "gofilm/drivers/mysql"

	_middleware "gofilm/app/middleware"
	_routes "gofilm/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Init() {
	viper.SetConfigFile("app/config/config.json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userRepo.Users{},
		&_genreRepo.Genres{},
		&_languageRepo.Languages{},
		&_filmsRepo.Films{},
		&_cartRepo.Carts{},
	)
}

func main() {
	Init()
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDB.InitialDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	
	e := echo.New()

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
	filmsService := _filmsServ.NewService(filmsRepo, filmsThird, genresService)
	filmsHandler := _filmsHandler.NewHandler(filmsService)

	cartsRepo := _cartRepo.NewMySQLRepo(db)
	cartsService := _cartServ.NewService(cartsRepo, filmsService)
	cartsHandler := _cartHandler.NewHandler(cartsService) 

	routesInit := _routes.HandlerList{
		JWTMiddleware: configJWT.Init(),
		UserHandler:   *usersHandler,
		GenreHandler:  *genresHandler,
		LanguageHandler: *langsHandler,
		FilmHandler: *filmsHandler,
		CartHandler: *cartsHandler,
	}

	routesInit.RouteUser(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}

































//carts
//unittest
//films casenya di genres
//payment
//koleksi
//CI/CD
//docker