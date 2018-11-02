package controllers

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/coderbunker/heikenet-backend/models"
)

func CreateUser(c echo.Context) error {
	new_user := new(models.NewUser)
	if err := c.Bind(new_user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "can't read json",
		})
	}

	user, err := models.CreateUser(c, new_user)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": err.Error(),
		})
	} else {
		return c.JSON(http.StatusCreated, user)
	}
}

func GetUser(c echo.Context) error {
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	return c.JSON(http.StatusConflict, map[string]string{
	// 		"error": "can't read id",
	// 	})
	// }
	//
	// user, err := models.GetUser(c, id)
	// if err != nil {
	// 	return c.JSON(http.StatusConflict, map[string]string{
	// 		"error": "can't find user",
	// 	})
	// } else {
	// 	return c.JSON(http.StatusCreated, user)
	// }
	return c.String(http.StatusOK, "id")
}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "update")
}

func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "delete")
}
