package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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

func CreateUser(db *gorm.DB, new_user *NewUser) (*User, error) {
	if new_user.Name == "" || new_user.Email == "" || new_user.Password == "" {
		return nil, errors.New("invalid account details")
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

func GetUser(db *gorm.DB, id int) (*User, error) {
	var user User
	if db.First(&user, "id = ?", id).RecordNotFound() {
		return nil, errors.New("user not found")
	} else {
		return &user, nil
	}
}

func UpdateUser(db *gorm.DB, id int, new_user *NewUser) (*User, error) {
	if new_user.Name == "" || new_user.Email == "" || new_user.Password == "" {
		return nil, errors.New("invalid account details")
	}

	var user User
	if db.First(&user, "id = ?", id).RecordNotFound() {
		return nil, errors.New("user not found")
	} else {
		hpassword, err := bcrypt.GenerateFromPassword([]byte(new_user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		db.Model(&user).Updates(
			User{
				Name:     new_user.Name,
				Email:    new_user.Email,
				Password: hpassword,
			},
		)

		return &user, nil
	}
}

func DeleteUser(db *gorm.DB, id int) error {
	var user User
	if db.First(&user, "id = ?", id).RecordNotFound() {
		return errors.New("user not found")
	} else {
		db.Delete(&user)
		return nil
	}
}
