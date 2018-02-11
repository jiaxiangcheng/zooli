package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
)

type Client struct {
	gorm.Model				`valid:"-"`
	Name			string	`valid:"-"`
	PhoneNumber		string	`valid:"numeric"`
	PasswordHash 	string	`valid:"-"`
	Email			string	`valid:"email"`
}


func (c *Client) Insert() {
	beego.Debug("Insert ", c)
	DB.Create(&c)
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
	beego.Debug("Update ", c)
	var cDB Client
	cDB.ID = c.ID
	DB.Where(&cDB).First(&cDB)

	cDB.Name = c.Name
	cDB.PasswordHash = c.PasswordHash
	cDB.PhoneNumber = c.PhoneNumber
	cDB.Email = c.Email

	DB.Save(&cDB)
}

func (c *Client) DeleteSoft() {
	beego.Debug("Delete ", c)
	DB.Delete(&c)
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
