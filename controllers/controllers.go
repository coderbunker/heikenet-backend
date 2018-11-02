package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "login")
}
