package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
	"encoding/json"
)

type Customer struct {
	gorm.Model					`valid:"-"`
	Name			string		`gorm:"not null" valid:"-"`
	PhoneNumber		string		`gorm:"not null" valid:"required,phone"`
	PasswordHash 	string		`gorm:"not null" valid:"required,alphanum"`
	Email			string		`gorm:"not null" valid:"required,email"`
	Vehicles		[]Vehicle	`valid:"-" json:"-"`
}


func (c *Customer) Insert() {
	DB.Create(&c)
	beego.Debug("Insert Customer:", c)
}

func (c *Customer) Exists() bool {
	count := 0
	DB.Where("id = ?", c.ID).Find(&Customer{}).Count(&count)
	return count > 0
}

func FindCustomerByID(id uint) Customer {
	var c Customer
	DB.Where("id = ?", id).Preload("Vehicles").Find(&c)
	return c
}

func FindCustomers() []Customer {
	var c []Customer
	DB.Preload("Vehicles").Find(&c)
	return c
}

func (c *Customer) Update() {
	var cDB Customer
	cDB.ID = c.ID
	DB.Where(&cDB).First(&cDB)

	cDB.Name = c.Name
	cDB.PasswordHash = c.PasswordHash
	cDB.PhoneNumber = c.PhoneNumber
	cDB.Email = c.Email

	DB.Save(&cDB)
	beego.Debug("Update Customer:", cDB)
}

func (c *Customer) DeleteSoft() {
	DB.Delete(&c)
	beego.Debug("Delete Customer:", c)
}


func (c *Customer) SetPassword(pass string) {
	c.PasswordHash = encryptMD5(pass)
}

func (c *Customer) ValidPassword(pass string) bool {
	if c.PasswordHash == encryptMD5(pass) {
		return true
	}
	return false
}

func (c Customer) String() string {
	out, err := json.Marshal(c)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return string(out)
}

func NumCustomers() int {
	var count int = 0
	var c []Customer
	DB.Find(&c).Count(&count)
	return count
}
