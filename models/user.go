package models

import (
	"fmt"
	"github.com/moskvinspace/go-auth/database"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggertype:"string"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
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

func GetUser(field, value string) (*User, error) {
	var user User

	if err := database.DB.
		Where(fmt.Sprintf("%s = ?", field), value).
		First(&user).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}
