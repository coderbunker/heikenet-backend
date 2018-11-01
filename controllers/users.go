package controllers

import (
	"github.com/coderbunker/heikenet-backend/models"
	"github.com/labstack/echo"
	"net/http"
)

func CreateUser(c echo.Context) error {
	new_user := new(models.NewUser)
	if err := c.Bind(new_user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "invalid json",
		})
	}
	user, err := models.CreateUser(c, new_user)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "can't create user",
		})
	}
	return c.JSON(http.StatusCreated, user)
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
