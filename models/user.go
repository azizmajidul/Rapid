package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	IdUser       int     `form:"iduser" json:"iduser" validate:"required" gorm:"primaryKey"`
	Name     string  `form:"nameuser" json:"nameuser" validate:"required"`
	Email     string  `form:"emailuser" json:"emailuser" validate:"required"`
	Password     string  `form:"passworduser" json:"passworduser" validate:"required"`
	
}

type SignIn struct {
	Email    string `form:"email" json:"email" validate:"required"`
	Password string `form:"password"  json:"password" validate:"required"`
}

func Registrasi(db *gorm.DB, newUser *User) (err error) {
	plainPassword := newUser.Password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(plainPassword),10)
	sHash := string(bytes)
	fmt.Println("Hash password: ", sHash)
	newUser.Password = sHash
	err = db.Create(newUser).Error
	if err != nil {
		return err
	}
	return nil
}



func GetUserByEmail(db *gorm.DB, users *User, email string) (err error) {
	err = db.Where("email=?", email).First(users).Error
	if err != nil {
		return err
	}
	return nil
}


