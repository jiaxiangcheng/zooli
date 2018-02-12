package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
	"encoding/json"
)

type Order struct {
	gorm.Model				`valid:"-"`
	Client		Client		`valid:"-" json:"-"`
	ClientID	uint		`gorm:"not null" valid:"-"`
	Product		Product		`valid:"-" json:"-"`
	ProductID	uint		`gorm:"not null" valid:"-"`
	Status		int			`gorm:"not null" valid:"-"`
	Price		float64		`valid:"-"`
}


func (o *Order) Insert() {
	DB.Create(&o)
	beego.Debug("Insert Order:", o)
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
	var oDB Order
	oDB.ID = o.ID
	DB.Where(&oDB).First(&oDB)

	oDB.Status = o.Status
	oDB.Price = o.Price

	DB.Save(&oDB)
	beego.Debug("Update Order:", o)
}

func (o *Order) DeleteSoft() {
	beego.Debug("Delete Order:", o)
	DB.Delete(&o)
}


func (o Order) String() string {
	out, err := json.Marshal(o)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return string(out)
}
