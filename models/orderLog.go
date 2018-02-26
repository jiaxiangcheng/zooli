package models

import (
	"time"
	"github.com/astaxie/beego"
)

type OrderLog struct {
	Status		Status		`gorm:"not null" valid:"required"`
	Timestamp	time.Time	`valid:"required"`
	OrderID		uint		`gorm:"not null" valid:"required"`
}


func (o *OrderLog) Insert() {
	DB.Create(&o)
	beego.Debug("Insert OrderLog:", o)
}


func FindOrderLogsByOrderID(orderID uint) []OrderLog {
	var o []OrderLog
	DB.Where("order_id = ?", orderID).Find(&o)
	return o
}
