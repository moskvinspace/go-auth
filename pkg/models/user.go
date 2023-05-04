package models

import (
	"github.com/moskvinspace/simple-web-app/pkg/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:""`
}

func (u *User) Create() error {
	if err := database.DB.
		Create(&u).
		Error; err != nil {
		return err
	}

	return nil
}

func IsEmailExist(email string) bool {
	var count int64

	database.DB.
		Model(&User{}).
		Where("email = ?", email).
		Count(&count)

	return count > 0
}

func GetUser(email string) (*User, error) {
	var user User

	if err := database.DB.
		Where("email = ?", email).
		First(&user).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}
