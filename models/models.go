package models

import (
	"errors"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"

	mid "github.com/coderbunker/heikenet-backend/middleware"
)

type (
	User struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password []byte `json:"-"`
	}

	NewUser struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func CreateUser(c echo.Context, new_user *NewUser) (*User, error) {
	if new_user.Name == "" || new_user.Email == "" || new_user.Password == "" {
		return nil, errors.New("invalid account details")
	}

	db, err := mid.GetDB(c)
	if err != nil {
		return nil, err
	}

	var user User
	if db.First(&user, "email = ?", new_user.Email).RecordNotFound() {
		hpassword, err := bcrypt.GenerateFromPassword([]byte(new_user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		user = User{
			Name:     new_user.Name,
			Email:    new_user.Email,
			Password: hpassword,
		}
		db.Create(&user)

		return &user, nil
	} else {
		return nil, errors.New("user exists")
	}
}
