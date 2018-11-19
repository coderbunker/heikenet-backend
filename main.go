package main

import (
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/coderbunker/heikenet-backend/controllers"
	mid "github.com/coderbunker/heikenet-backend/middleware"
	"github.com/coderbunker/heikenet-backend/models"
)

func main() {
	// get app config from env
	config := models.AppConfig{}
	err := env.Parse(&config)
	if err != nil {
		log.Fatal(err)
	}

	// connect to db
	db, err := gorm.Open("postgres", config.DatabaseURL)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// this is for dev
	// db.AutoMigrate(
	// 	&models.User{},
	// 	&models.Profile{},
	// )
	// db.DropTable(
	// 	&models.User{},
	// 	&models.Profile{},
	// )
	// db.CreateTable(
	// 	&models.User{},
	// 	&models.Profile{},
	// )

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
	}))
	e.Use(mid.SetConfig(config))
	e.Use(mid.SetDB(db))

	// routes: url/login, url/register
	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.CreateUser)

	// all routes in this group are protected with jwt middleware
	// routes: url/api/v1/users
	api_v1 := e.Group("/api/v1")
	api_v1.Use(middleware.JWT([]byte(config.Secret)))
	api_v1.GET("/users", controllers.GetUser)
	api_v1.PUT("/users", controllers.UpdateUser)
	api_v1.DELETE("/users", controllers.DeleteUser)

	// routes: url/api/v1/users/profiles
	api_v1.POST("/users/profiles", controllers.CreateProfile)
	api_v1.GET("/users/profiles", controllers.GetProfile)
	api_v1.PUT("/users/profiles", controllers.UpdateProfile)
	api_v1.DELETE("/users/profiles", controllers.DeleteProfile)

	// routes: url/api/v1/users/contracts/*
	api_v1.POST("/users/contracts/approve", controllers.Approve)
	api_v1.POST("/users/contracts/fund", controllers.Fund)
	api_v1.POST("/users/contracts/withdraw", controllers.Withdraw)

	// start server
	e.Logger.Fatal(e.Start(":" + config.Port))
}
