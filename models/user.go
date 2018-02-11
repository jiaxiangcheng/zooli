package models

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
)

type User struct {
	gorm.Model				`valid:"-"`
	Username     	string	`valid:"alphanum"`   //PK
	PasswordHash 	string	`valid:"-"`
	Email			string	`valid:"email,optional"`
	Name			string	`valid:"-"`
	Role			Role	`valid:"-"`
	RoleID			uint	`valid:"-"`
	Store			Store	`gorm:"foreignkey:ManagerID"`
}


func (u *User) Insert() {
	beego.Debug("Insert ", u)
	DB.Create(&u)
}

func (u *User) Exists() bool {
	count := 0
	DB.Where("id = ?", u.ID).Find(&User{}).Count(&count)
	return count > 0
}

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

func ExistUserByUsername(username string) bool {
	var u User
	count := 0
	DB.Where("username = ?", username).Find(&u).Count(&count)
	return count > 0
}

func FindUsers() []User {
	var u []User
	DB.Find(&u)
	return u
}


func (u *User) Update() {
	beego.Debug("Update ", u)
	var uDB User
	uDB.ID = u.ID
	DB.Where(&uDB).First(&uDB)
	uDB.Username = u.Username
	uDB.PasswordHash = u.PasswordHash
	uDB.Email = u.Email
	uDB.Name = u.Name
	uDB.RoleID = u.RoleID

	DB.Save(&uDB)
}

func (u *User) DeleteSoft() {
	beego.Debug("Delete ", u)
	DB.Delete(&u)
}


func (u *User) SetPassword(pass string) {
	u.PasswordHash = encryptMD5(pass)
}

func (u *User) ValidPassword(pass string) bool {
	if u.PasswordHash == encryptMD5(pass) {
		return true
	}
	return false
}

func encryptMD5(pass string) string {
	h := md5.New()
	io.WriteString(h, pass)
	str := hex.EncodeToString(h.Sum(nil))

	return str
}