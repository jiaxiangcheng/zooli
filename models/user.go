package models

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model   `valid:"-"`
	Username     string `gorm:"not null;unique" valid:"alphanum,required"`
	PasswordHash string `gorm:"not null" valid:"required,alphanum"`
	Email        string `valid:"email,optional"`
	Name         string `gorm:"not null" valid:"required"`
	Role         Role   `valid:"-" json:"-"`
	RoleID       uint   `gorm:"not null" valid:"required"`
	StoreID      uint   `valid:"-"`
}

func (u *User) Insert() {
	DB.Create(&u)
	beego.Debug("Insert User:", u)
}

func (u *User) Exists() bool {
	count := 0
	DB.Where("id = ?", u.ID).Find(&User{}).Count(&count)
	return count > 0
}

func (u *User) ExistsUsername() bool {
	count := 0
	DB.Unscoped().Where("username = ? and id <> ?", u.Username, u.ID).Find(&User{}).Count(&count)
	return count > 0
}

func FindUserByID(id uint) User {
	var u User
	DB.Preload("Role").Where("id = ?", id).Find(&u)
	return u
}

func FindUserByUsername(username string) User {
	var u User
	DB.Preload("Role").Where("username = ?", username).Find(&u)
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
	DB.Preload("Role").Find(&u)
	return u
}

func FindUsersByRoleID(roleId uint) []User {
	var u []User
	DB.Preload("Role").Where("role_id = ?", roleId).Find(&u)
	return u
}

/*
func FindManagersWithoutStore() []User {
	var u []User
	DB.Where("role_id = ? and store_id = ?", 2, 0).Preload("Role").Find(&u)
	return u
}
*/

func (u *User) Update() {
	var uDB User
	uDB.ID = u.ID
	DB.Where(&uDB).First(&uDB)

	uDB.Username = u.Username
	uDB.PasswordHash = u.PasswordHash
	uDB.Email = u.Email
	uDB.Name = u.Name
	uDB.RoleID = u.RoleID

	DB.Save(&uDB)
	beego.Debug("Update User:", u)
}

func (u *User) AssignStore(store Store) {
	u.StoreID = store.ID
	DB.Save(&u)
	beego.Debug("Update User:", u)
}

func (u *User) DeleteSoft() {
	DB.Delete(&u)
	beego.Debug("Delete User:", u)
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

func (u User) String() string {
	out, err := json.Marshal(u)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return string(out)
}

func NumUsers() int {
	var count int = 0
	var u []User
	DB.Find(&u).Count(&count)
	return count
}
