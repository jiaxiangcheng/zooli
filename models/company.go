package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
)

type Company struct {
	gorm.Model				`valid:"-"`
	Name			string	`valid:"-"`
	Contact			string	`valid:"-"`
	PhoneNumber		string	`valid:"numeric,optional"`
	Email			string	`valid:"email,optional"`
}



func (c *Company) Insert() {
	beego.Debug("Insert ", c)
	DB.Create(&c)
}

func (c *Company) Exists() bool {
	count := 0
	DB.Where("id = ?", c.ID).Find(&Company{}).Count(&count)
	return count > 0
}

func (c *Company) ExistsName() bool {
	count := 0
	DB.Where("name = ? and id <> ?", c.Name, c.ID).Find(&Company{}).Count(&count)
	return count > 0
}

func FindCompanyByID(id uint) Company {
	var c Company
	DB.Where("id = ?", id).Find(&c)
	return c
}

func FindCompanyByName(name string) Company {
	var c Company
	DB.Where("name = ?", name).Find(&c)
	return c
}

func ExistCompanyByName(name string) bool {
	var c Company
	count := 0
	DB.Where("name = ?", name).Find(&c).Count(&count)
	return count > 0
}

func FindCompanys() []Company {
	var c []Company
	DB.Find(&c)
	return c
}

func (c *Company) Update() {
	beego.Debug("Update ", c)
	var cDB Company
	cDB.ID = c.ID
	DB.Where(&cDB).First(&cDB)

	cDB.Name = c.Name
	cDB.Contact = c.Contact
	cDB.PhoneNumber = c.PhoneNumber
	cDB.Email = c.Email

	DB.Save(&cDB)
}

func (c *Company) DeleteSoft() {
	beego.Debug("Delete ", c)
	DB.Delete(&c)
}

