package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "welcome to heike-network")
}

func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
