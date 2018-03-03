package models

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

type ProductImage struct {
	gorm.Model			`valid:"-"`
	Image		string	`gorm:"not null" valid:"required, url"`
	ProductID	uint	`gorm:"not null" valid:"required"`
}

func (i *ProductImage) Insert() {
	DB.Create(&i)
	beego.Debug("Insert Product Image:", i)
}

func FindImagesByProductID(productID uint) []ProductImage {
	var i []ProductImage
	DB.Where("product_id = ?", productID).Find(&i)
	return i
}

