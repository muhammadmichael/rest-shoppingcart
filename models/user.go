package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `form:"name" json: "name" validate:"required"`
	Username   string `form:"username" json: "username" validate:"required"`
	Email      string `form:"email" json: "email" validate:"required"`
	Password   string `form:"password" json: "password" validate:"required"`
	Cart       Cart
	Transaksis []Transaksi
}

func CreateUser(db *gorm.DB, newUser *User) (err error) {
	err = db.Create(newUser).Error
	if err != nil {
		return err
	}
	return nil
}

func FindUserByUsername(db *gorm.DB, user *User, username string) (err error) {
	err = db.Where("username=?", username).First(user).Error
	if err != nil {
		return err
	}
	return nil
}
