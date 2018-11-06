package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"

	mid "github.com/coderbunker/heikenet-backend/middleware"
	"github.com/coderbunker/heikenet-backend/models"
)

const EXP = 72

func Login(c echo.Context) error {
	login_user := new(models.LoginUser)
	if err := c.Bind(login_user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "invalid json",
		})
	}

	if login_user.Email == "" || login_user.Password == "" {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "invalid account details",
		})
	}

	db, err := mid.GetDB(c)
	if err != nil {
		log.Fatal(err)
	}

	config, err := mid.GetConfig(c)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if db.First(&user, "email = ?", login_user.Email).RecordNotFound() {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "not found",
		})
	} else {
		err = bcrypt.CompareHashAndPassword(user.Password, []byte(login_user.Password))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "unauthorized",
			})
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = user.ID
		claims["exp"] = time.Now().Add(time.Hour * EXP).Unix()

		t, err := token.SignedString([]byte(config.Secret))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
}
