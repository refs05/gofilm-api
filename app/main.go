package main

import (
	"fmt"
	"gofilm/app/constant"
	"log"

	_userServ "gofilm/bussinesses/users"
	_userHandler "gofilm/controllers/users"
	_userRepo "gofilm/drivers/databases/users"

	_middleware "gofilm/app/middleware"
	_routes "gofilm/app/routes"

	"github.com/labstack/echo/v4"
	//"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(status string) *gorm.DB {
	db := "gofilm"
	// if status == "testing" {
	// 	db = "gofilm-test"
	// }
	connectionString := fmt.Sprintf("root:BavarianRestu2002@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", db)

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&_userRepo.Users{},
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

	routesInit := _routes.HandlerList{
		JWTMiddleware: configJWT.Init(),
		UserHandler:   *usersHandler,
	}

	routesInit.RouteUser(e)

	log.Fatal(e.Start(":8080"))
}


//redis belum
//config
//pisah config db
//categoryusecase
//unittest