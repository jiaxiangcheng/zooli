package models

import (
	"github.com/astaxie/beego"
	"time"
)

type StoreImage struct {
	ID			uint		`gorm:"primary_key" valid:"-"`
	CreatedAt	time.Time	`valid:"-"`
	Image      string `gorm:"not null" valid:"required, url"`
	StoreID    uint   `gorm:"not null" valid:"required"`
}

func (i *StoreImage) Insert() {
	DB.Create(&i)
	beego.Debug("Insert Store Image:", i)
}

func FindImagesByStoreID(storeID uint) []StoreImage {
	var i []StoreImage
	DB.Where("store_id = ?", storeID).Find(&i)
	return i
}
