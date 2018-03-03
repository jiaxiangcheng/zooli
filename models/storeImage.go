package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
)

type StoreImage struct {
	gorm.Model			`valid:"-"`
	Image		string	`gorm:"not null" valid:"required, url"`
	StoreID		uint	`gorm:"not null" valid:"required"`
}

func (i *StoreImage) Insert() {
	DB.Create(&i)
	beego.Debug("Insert Product Image:", i)
}

func FindImagesByStoreID(storeID uint) []StoreImage {
	var i []StoreImage
	DB.Where("product_id = ?", storeID).Find(&i)
	return i
}


