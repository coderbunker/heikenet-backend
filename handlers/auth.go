package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	mid "github.com/coderbunker/heikenet-backend/middleware"
	"github.com/coderbunker/heikenet-backend/models"
)

const EXP = 72

func Login(c echo.Context) error {
	db, err := mid.GetDB(c)
	if err != nil {
		log.Fatal(err)
	}

	config, err := mid.GetConfig(c)
	if err != nil {
		log.Fatal(err)
	}

	googleOauthConfig := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, c.FormValue("code"))
	if err != nil {
		log.Println("code exchange failed")
		return err
	}

	if !token.Valid() {
		log.Println("Retreived invalid token")
		return err
	}

	// https://www.googleapis.com/oauth2/v3/tokeninfo?access_token=
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Println("failed getting user info")
		return err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("failed reading response body")
		return err
	}

	var user *models.GoogleUser
	err = json.Unmarshal(content, &user)
	if err != nil {
		log.Printf("Error unmarshaling Google user")
		return err
	}

	log.Println("id ---->", user.ID)

	var profile models.Profile
	if db.First(&profile, "id = ?", user.ID).RecordNotFound() {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "not found",
		})
	} else {
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
