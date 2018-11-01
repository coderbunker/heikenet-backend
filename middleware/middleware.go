package middleware

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
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

// func GetSecret(c echo.Context) (string, error) {
// 	secret, ok := c.Get("secret").(string)
// 	if !ok {
// 		return "", errors.New("no secret in context")
// 	}
// 	return secret, nil
// }
//
// func SetSecret(secret string) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			c.Set("secret", secret)
// 			next(c)
// 			return nil
// 		}
// 	}
// }
