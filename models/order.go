package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
)

type Order struct {
	gorm.Model				`valid:"-"`
	Client		Client		`valid:"-"`
	ClientID	uint		`gorm:"not null" valid:"-"`
	Product		Product		`valid:"-"`
	ProductID	uint		`gorm:"not null" valid:"-"`
	Status		int			`gorm:"not null" valid:"-"`
	Price		float64		`valid:"-"`
}


func (o *Order) Insert() {
	beego.Debug("Insert ", o)
	DB.Create(&o)
}

func (o *Order) Exists() bool {
	count := 0
	DB.Where("id = ?", o.ID).Find(&Order{}).Count(&count)
	return count > 0
}

func FindOrderByID(id uint) Order {
	var o Order
	DB.Where("id = ?", id).Find(&o)
	return o
}


func FindOrders() []Order {
	var o []Order
	DB.Find(&o)
	return o
}

func (o *Order) Update() {
	beego.Debug("Update ", o)
	var oDB Order
	oDB.ID = o.ID
	DB.Where(&oDB).First(&oDB)

	oDB.Status = o.Status
	oDB.Price = o.Price

	DB.Save(&oDB)
}

func (o *Order) DeleteSoft() {
	beego.Debug("Update ", o)
	DB.Delete(&o)
}


