package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
	"encoding/json"
	"time"
)


type Status int
const (
	ORDERED Status = iota + 1
	IN_SERVICE
	END_SERVICE
	WAITING_FOR_PAYMENT
	FINISHED
	CANCELED
)

var statuses = [...]string {
	"ORDERED",
	"IN SERVICE",
	"END SERVICE",
	"WAITING FOR PAYMENT",
	"FINISHED",
	"CANCELED",
}

func (status Status) String() string {
	return statuses[status - 1]
}



type Order struct {
	gorm.Model				`valid:"-"`
	Customer	Customer	`valid:"-" json:"-"`
	CustomerID	uint		`gorm:"not null" valid:"required"`
	Product		Product		`valid:"-" json:"-"`
	ProductID	uint		`gorm:"not null" valid:"required"`
	Status		Status		`gorm:"not null" valid:"required"`
	Fee			float64		`valid:"-"`
	Logs		[]OrderLog	`valid:"-" json:"-"`
}


func (o *Order) Insert() {
	DB.Create(&o)
	if len(o.Logs) == 0 {
		l := OrderLog{
			o.Status,
			time.Now(),
			o.ID,
		}
		l.Insert()
	}

	beego.Debug("Insert Order:", o)
}

func (o *Order) Exists() bool {
	count := 0
	DB.Where("id = ?", o.ID).Find(&Order{}).Count(&count)
	return count > 0
}

func FindOrderByID(id uint) Order {
	var o Order
	DB.Where("id = ?", id).Preload("Customer").Preload("Product").Preload("Logs").Find(&o)
	return o
}

func FindOrdersByStoreID(storeID uint) []Order {
	var o []Order
	DB.Where("product_id in (?)",
		DB.Table("products").Select("id").Where("store_id = ?", storeID).QueryExpr()).
		Preload("Customer").Preload("Product").Find(&o)

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

	if oDB.Status != o.Status {
		l := OrderLog{
			Status: o.Status,
			Timestamp: time.Now(),
			OrderID: o.ID,
		}
		l.Insert()
	}
	oDB.Status = o.Status
	oDB.Fee = o.Fee

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

func NumOrders() int {
	var count int = 0
	var c []Order
	DB.Find(&c).Count(&count)
	return count
}
