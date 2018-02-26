package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/astaxie/beego"
)

type OrderLog struct {
	gorm.Model				`valid:"-"`
	Status		Status		`gorm:"not null" valid:"required"`
	Timestamp	time.Time	`gorm:"not null" valid:"required"`
	OrderID		uint		`gorm:"not null" valid:"required"`
}


func (o *OrderLog) Insert() {
	DB.Create(&o)
	beego.Debug("Insert OrderLog:", o)
}

func FindOrderLogByIO(id uint) OrderLog {
	var o OrderLog
	DB.Where("id = ?", id).Find(&o)
	return o
}

func FindOrderLogsByOrderID(orderID uint) []OrderLog {
	var o []OrderLog
	DB.Where("order_id = ?", orderID).Find(&o)
	return o
}
