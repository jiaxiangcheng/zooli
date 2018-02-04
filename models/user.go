package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model 			`valid:"-"`
	Username   string	`valid:"alphanum"`
	Password   string	`valid:"-"`
}

func (u *User) SetPassword(pass string) {
	u.Password = encryptPasswordMD5(pass)
}

func (u *User) ValidPassword(pass string) bool {
	if u.Password == encryptPasswordMD5(pass) {
		return true
	}
	return false
}

func encryptPasswordMD5(pass string) string {
	h := md5.New()
	io.WriteString(h, pass)
	str := hex.EncodeToString(h.Sum(nil))

	return str
}

func (u *User) Insert() {
	DB.Create(u)
}

func (u *User) Exists() bool {
	count := 0
	DB.Where("id = ?", u.ID).Find(&User{}).Count(&count)
	return count > 0
}

// ExistsUsername returns true if there is already an user with this username in the DB
func (u *User) ExistsUsername() bool {
	count := 0
	DB.Where("username = ? and id <> ?", u.Username, u.ID).Find(&User{}).Count(&count)
	return count > 0

}

func FindUserByID(id uint) User {
	var u User
	DB.Where("id = ?", id).Find(&u)
	return u
}

func FindUserByUsername(username string) User {
	var u User
	DB.Where("username = ?", username).Find(&u)
	return u
}

func FindUsers() []User {
	var u []User
	DB.Find(&u)
	return u
}

func (u *User) ChangePassword(pass string) {

	DB.Where("username = ?", u.Username).First(&u)

	u.SetPassword(pass)
	//hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	//u.Hash = string(hashedPassword[:])

	DB.Save(&u)
}

func (u *User) Update() {
	var uDB User
	uDB.ID = u.ID
	DB.Where(&uDB).First(&uDB)

	fmt.Print(uDB)
	uDB.Username = u.Username

	DB.Save(&uDB)
}

func (u *User) DeleteSoft() {
	DB.Delete(&u)
}
