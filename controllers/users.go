package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"

	mid "github.com/coderbunker/heikenet-backend/middleware"
	"github.com/coderbunker/heikenet-backend/models"
)

func CreateUser(c echo.Context) error {
	u := new(models.NewUser)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "invalid json",
		})
	}

	if u.Name == "" || u.Email == "" || len(u.Password) == 0 {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "invalid account details",
		})
	}

	db, err := mid.GetDB(c)
	if err != nil {
		return err
	}

	var user models.User
	if db.First(&user, "email = ?", u.Email).RecordNotFound() {
		hpassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user = models.User{
			Name:     u.Name,
			Email:    u.Email,
			Password: hpassword,
		}
		db.Create(&user)

		return c.JSON(http.StatusCreated, user)
	} else {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "user exists",
		})
	}
}

func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "get")
}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "update")
}

func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "delete")
}
