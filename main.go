package main

import (
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	uuid "github.com/satori/go.uuid"

	"github.com/coderbunker/heikenet-backend/controllers"
	mid "github.com/coderbunker/heikenet-backend/middleware"
	"github.com/coderbunker/heikenet-backend/models"
)

func main() {
	config := models.AppConfig{}
	err := env.Parse(&config)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("postgres", config.Database_url)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// this is for dev
	db.AutoMigrate(
		&models.User{},
		&models.Profile{},
	)
	db.DropTable(
		&models.User{},
		&models.Profile{},
	)
	db.CreateTable(
		&models.User{},
		&models.Profile{},
	)

	// test users for dev ---------------------------------
	test_user, _ := models.CreateUser(db, &models.NewUser{Name: "denis", Email: "mail", Password: "42"})
	test_user_uuid, _ := uuid.FromString(test_user.ID)
	test_profile, _ := models.CreateProfile(db, test_user_uuid, &models.NewProfile{Rate: 4.2, Info: "a@b.c", Wallet: "0x00001"})
	log.Println(test_user.ID)
	log.Println(test_profile.ID)
	//-----------------------------------------------------

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(mid.SetConfig(config))
	e.Use(mid.SetDB(db))

	api_v1 := e.Group("/api/v1")
	api_v1.POST("/login", controllers.Login)
	api_v1.POST("/register", controllers.CreateUser)

	// routes: url/api/v1/users
	users := api_v1.Group("/users")
	// users.Use(middleware.JWT([]byte(config.Secret)))
	users.GET("/:id", controllers.GetUser)
	users.PUT("/:id", controllers.UpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)

	//--------------------------------------------------
	//TODO: get uuid from token!!! remove id from routes
	//--------------------------------------------------

	// routes: url/api/v1/users/profiles
	profiles := users.Group("/:id/profiles")
	profiles.POST("", controllers.CreateProfile)
	profiles.GET("/:pid", controllers.GetProfile)
	profiles.PUT("/:pid", controllers.UpdateProfile)
	profiles.DELETE("/:pid", controllers.DeleteProfile)

	// routes: url/api/v1/users/contracts
	contracts := users.Group("/contracts")
	contracts.POST("/approve", controllers.Approve)
	contracts.POST("/fund", controllers.Fund)
	contracts.POST("/withdraw", controllers.Withdraw)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
