package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	GoogleUser struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Link          string `json:"link"`
		Picture       string `json:"picture"`
		Gender        string `json:"gender"`
		Locale        string `json:"locale"`
	}

	User struct {
		ID        string    `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  []byte    `json:"-"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Profile   Profile
	}

	NewUser struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginUser struct {
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

func GetUser(db *gorm.DB, id uuid.UUID) (*User, error) {
	var user User
	if db.First(&user, "id = ?", id).RecordNotFound() {
		return nil, errors.New("user not found")
	} else {
		return &user, nil
	}
}

func UpdateUser(db *gorm.DB, id uuid.UUID, new_user *NewUser) (*User, error) {
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

func DeleteUser(db *gorm.DB, id uuid.UUID) error {
	var user User
	if db.First(&user, "id = ?", id).RecordNotFound() {
		return errors.New("user not found")
	} else {
		db.Delete(&user)
		return nil
	}
}
