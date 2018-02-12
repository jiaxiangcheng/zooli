package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
	"encoding/json"
)

type Client struct {
	gorm.Model				`valid:"-"`
	Name			string	`gorm:"not null" valid:"-"`
	PhoneNumber		string	`gorm:"not null" valid:"numeric"`
	PasswordHash 	string	`gorm:"not null" valid:"-"`
	Email			string	`gorm:"not null" valid:"email"`
}


func (c *Client) Insert() {
	DB.Create(&c)
	beego.Debug("Insert Client:", c)
}

func (c *Client) Exists() bool {
	count := 0
	DB.Where("id = ?", c.ID).Find(&Client{}).Count(&count)
	return count > 0
}

func FindClientByID(id uint) Client {
	var c Client
	DB.Where("id = ?", id).Find(&c)
	return c
}

func FindClients() []Client {
	var c []Client
	DB.Find(&c)
	return c
}

func (c *Client) Update() {
	var cDB Client
	cDB.ID = c.ID
	DB.Where(&cDB).First(&cDB)

	cDB.Name = c.Name
	cDB.PasswordHash = c.PasswordHash
	cDB.PhoneNumber = c.PhoneNumber
	cDB.Email = c.Email

	DB.Save(&cDB)
	beego.Debug("Update Client:", cDB)
}

func (c *Client) DeleteSoft() {
	DB.Delete(&c)
	beego.Debug("Delete Client:", c)
}


func (c *Client) SetPassword(pass string) {
	c.PasswordHash = encryptMD5(pass)
}

func (c *Client) ValidPassword(pass string) bool {
	if c.PasswordHash == encryptMD5(pass) {
		return true
	}
	return false
}

func (c Client) String() string {
	out, err := json.Marshal(c)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return string(out)
}