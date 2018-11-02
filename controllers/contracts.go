package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Approve(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}

func Fund(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}

func Withdraw(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
