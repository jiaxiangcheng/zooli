package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
)

type Product struct {
	gorm.Model					`valid:"-"`
	Name			string		`gorm:"not null" valid:"-"`
	Description		string		`gorm:"type:longtext" valid:"-"`
	Value			float64		`valid:"float,optional"`
	Image			string		`valid:"url,optional"`
	Service			Service		`valid:"-"`
	ServiceID		uint		`gorm:"not null" valid:"-"`
}


func (p *Product) Insert() {
	beego.Debug("Insert ", p)
	DB.Create(&p)
}

func (p *Product) Exists() bool {
	count := 0
	DB.Where("id = ?", p.ID).Find(&Product{}).Count(&count)
	return count > 0
}

func (p *Product) ExistsName() bool {
	count := 0
	DB.Where("name = ? and id <> ?", p.Name, p.ID).Find(&Product{}).Count(&count)
	return count > 0

}

func FindProductByID(id uint) Product {
	var p Product
	DB.Where("id = ?", id).Find(&p)
	return p
}

func ExistProductByName(name string) bool {
	var p Product
	count := 0
	DB.Where("name = ?", name).Find(&p).Count(&count)
	return count > 0
}

func FindProducts() []Product {
	var p []Product
	DB.Find(&p)
	return p
}

func (p *Product) Update() {
	beego.Debug("Update ", p)
	var pDB Product
	pDB.ID = p.ID
	DB.Where(&pDB).First(&pDB)

	pDB.Name = p.Name
	pDB.Description = p.Description
	pDB.Value = p.Value
	pDB.Image = p.Image
	pDB.ServiceID = p.ServiceID

	DB.Save(&pDB)
}

func (p *Product) DeleteSoft() {
	beego.Debug("Update ", p)
	DB.Delete(&p)
}


