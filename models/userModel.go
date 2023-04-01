package models

import (
	"errors"

	"github.com/omenejoseph/go-crud/inititalizers"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func GetUserByID(uid uint) (User, error) {

	var u User
	db := inititalizers.ConnectToDB()
	
	if err := db.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}
