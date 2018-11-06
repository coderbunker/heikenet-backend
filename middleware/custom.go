package middleware

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/coderbunker/heikenet-backend/models"
)

func GetDB(c echo.Context) (*gorm.DB, error) {
	db, ok := c.Get("db").(*gorm.DB)
	if !ok {
		return nil, errors.New("no db in context")
	}
	return db, nil
}

func SetDB(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			next(c)
			return nil
		}
	}
}

func GetConfig(c echo.Context) (*models.AppConfig, error) {
	config, ok := c.Get("config").(models.AppConfig)
	if !ok {
		return nil, errors.New("no config in context")
	}
	return &config, nil
}

func SetConfig(config models.AppConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("config", config)
			next(c)
			return nil
		}
	}
}
