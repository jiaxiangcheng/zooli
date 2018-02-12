package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
	"encoding/json"
)

type Product struct {
	gorm.Model					`valid:"-"`
	Name			string		`gorm:"not null" valid:"-"`
	Description		string		`gorm:"type:longtext" valid:"-"`
	Value			float64		`valid:"float,optional"`
	Image			string		`valid:"url,optional"`
	Service			Service		`valid:"-" json:"-"`
	ServiceID		uint		`gorm:"not null" valid:"-"`
}


func (p *Product) Insert() {
	DB.Create(&p)
	beego.Debug("Insert Product:", p)
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
	var pDB Product
	pDB.ID = p.ID
	DB.Where(&pDB).First(&pDB)

	pDB.Name = p.Name
	pDB.Description = p.Description
	pDB.Value = p.Value
	pDB.Image = p.Image
	pDB.ServiceID = p.ServiceID

	DB.Save(&pDB)
	beego.Debug("Update Product:", p)
}

func (p *Product) DeleteSoft() {
	DB.Delete(&p)
	beego.Debug("Delete Product:", p)
}

func (p Product) String() string {
	out, err := json.Marshal(p)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return string(out)
}



