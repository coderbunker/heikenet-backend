package main

import (
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/coderbunker/heikenet-backend/handlers"
	mid "github.com/coderbunker/heikenet-backend/middleware"
	"github.com/coderbunker/heikenet-backend/models"
)

func main() {
	config := models.AppConfig{}
	err := env.Parse(&config)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("postgres", config.DatabaseURL)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

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

	api := echo.New()
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{config.Hostname},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
	}))
	api.Use(mid.SetConfig(config))
	api.Use(mid.SetDB(db))
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.Secret),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/api/v1/login" {
				return true
			}
			return false
		},
	}))

	api_v1 := api.Group("/api/v1")
	api_v1.POST("/login", handlers.Login)

	api_v1.POST("/profiles", handlers.CreateProfile)
	api_v1.GET("/profiles", handlers.GetProfile)
	api_v1.PUT("/profiles", handlers.UpdateProfile)
	api_v1.DELETE("/profiles", handlers.DeleteProfile)

	api_v1.POST("/contracts/approve", handlers.Approve)
	api_v1.POST("/contracts/fund", handlers.Fund)
	api_v1.POST("/contracts/withdraw", handlers.Withdraw)

	api.Logger.Fatal(api.Start(":" + config.Port))
}
