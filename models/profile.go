package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type (
	Profile struct {
		ID        string    `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
		Wallet    string    `json:"wallet"`
		Info      string    `json:"info"`
		Rate      float32   `json:"rate"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		UserID    string    `json:"id" gorm:"type:uuid"`
	}

	NewProfile struct {
		Rate   float32 `json:"rate"`
		Info   string  `json:"info"`
		Wallet string  `json:"wallet"`
	}
)

func CreateProfile(db *gorm.DB, user_id uuid.UUID, new_profile *NewProfile) (*Profile, error) {
	if new_profile.Rate <= 0.0 {
		return nil, errors.New("invalid profile details")
	}

	var user User
	if db.First(&user, "id = ?", user_id).RecordNotFound() {
		return nil, errors.New("user is not exists")
	} else {
		var profile Profile
		if db.Model(&user).Related(&profile).RecordNotFound() {
			profile = Profile{
				Info:   new_profile.Info,
				Rate:   new_profile.Rate,
				Wallet: new_profile.Wallet,
				UserID: user.ID,
			}
			db.Create(&profile)

			return &profile, nil
		}
		return &profile, nil
	}
}

func GetProfile(db *gorm.DB, id uuid.UUID) (*Profile, error) {
	var profile Profile
	if db.First(&profile, "id = ?", id).RecordNotFound() {
		return nil, errors.New("profile not found")
	} else {
		return &profile, nil
	}
}

func UpdateProfile(db *gorm.DB, id uuid.UUID, new_profile *NewProfile) (*Profile, error) {
	if new_profile.Rate <= 0.0 || new_profile.Info == "" || new_profile.Wallet == "" {
		return nil, errors.New("invalid profile details")
	}

	var profile Profile
	if db.First(&profile, "id = ?", id).RecordNotFound() {
		return nil, errors.New("profile not found")
	} else {
		db.Model(&profile).Updates(
			Profile{
				Rate:   new_profile.Rate,
				Info:   new_profile.Info,
				Wallet: new_profile.Wallet,
			},
		)

		return &profile, nil
	}
}

func DeleteProfile(db *gorm.DB, id uuid.UUID) error {
	var profile Profile
	if db.First(&profile, "id = ?", id).RecordNotFound() {
		return errors.New("profile not found")
	} else {
		db.Delete(&profile)
		return nil
	}
}
