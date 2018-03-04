package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
	"encoding/json"
)

type Product struct {
	gorm.Model						`valid:"-"`
	Name			string			`gorm:"not null" valid:"required"`
	Description		string			`gorm:"type:longtext" valid:"-"`
	Value			float64			`valid:"optional"`
	Images			[]ProductImage	`valid:"-"`
	Service			Service			`valid:"-" json:"-"`
	ServiceID		uint			`gorm:"not null" valid:"required"`
	Store			Store			`valid:"-" json:"-"`
	StoreID			uint			`gorm:"not null" valid:"required"`
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
	DB.Where("id = ?", id).Preload("Images").Preload("Service").Preload("Store").Find(&p)
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
	DB.Preload("Service").Preload("Store").Find(&p)
	return p
}

func FindProductsByStoreID(storeID uint) []Product {
	var p []Product
	DB.Where("store_id = ?", storeID).Preload("Service").Preload("Store").Find(&p)
	return p
}

func (p *Product) Update() {
	var pDB Product
	pDB.ID = p.ID
	DB.Where(&pDB).First(&pDB)

	DB.Model(&pDB).Association("Images").Replace(p.Images)

	pDB.Name = p.Name
	pDB.Description = p.Description
	pDB.Value = p.Value
	pDB.ServiceID = p.ServiceID
	pDB.StoreID = p.StoreID

	DB.Save(&pDB)
	beego.Debug("Update Product:", pDB)
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



