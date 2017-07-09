package models

import (
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Email string `gorm:"primary_key"`
	Password string `gorm:"size:255";gorm:"-"`
}


func (u *User) GetAuthenicatedUser (email string, password string) *User {
	authenticatedUser := User{}
	password = EncodePassword(password)
	Database.Find(&authenticatedUser).Where(&User{Email: email, Password: password})
	return &authenticatedUser
}


func EncodePassword(password string) string {
	data := []byte(password)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
