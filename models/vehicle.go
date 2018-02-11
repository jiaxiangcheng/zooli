package models

import (
	"time"
	"github.com/astaxie/beego"
)

type Vehicle struct {
	Plate			string		`gorm:"primary_key" valid:"alphanum"`
	CreatedAt		time.Time	`valid:"-"`
	UpdatedAt		time.Time	`valid:"-"`
	DeletedAt		*time.Time	`sql:"index" valid:"-"`
	Model			string		`valid:"-"`
	Owner			Client		`valid:"-"`
	OwnerID			uint		`valid:"-"`
}


func (v *Vehicle) Insert() {
	beego.Debug("Insert ", v)
	DB.Create(&v)
}

func (v *Vehicle) Exists() bool {
	count := 0
	DB.Where("plate = ?", v.Plate).Find(&Vehicle{}).Count(&count)
	return count > 0
}

func FindVehicleByPlate(id string) Vehicle {
	var v Vehicle
	DB.Where("plate = ?", id).Find(&v)
	return v
}

func ExistVehicleByPlate(plate string) bool {
	var v Vehicle
	count := 0
	DB.Where("plate = ?", plate).Find(&v).Count(&count)
	return count > 0
}

func FindVehicles() []Vehicle {
	var v []Vehicle
	DB.Find(&v)
	return v
}

func (v *Vehicle) Update() {
	beego.Debug("Update ", v)
	var vDB Vehicle
	DB.Where("plate = ?", v.Plate).First(&vDB)

	vDB.Model = v.Model
	vDB.OwnerID = v.OwnerID

	DB.Save(&vDB)
}

func (v *Vehicle) DeleteSoft() {
	beego.Debug("Update ", v)
	DB.Delete(&v)
}
