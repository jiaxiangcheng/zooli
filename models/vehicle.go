package models

import (
	"time"
	"github.com/astaxie/beego"
	"encoding/json"
)

type Vehicle struct {
	ID				uint		`gorm:"primary_key" valid:"-"`
	CreatedAt		time.Time	`valid:"-"`
	UpdatedAt		time.Time	`valid:"-"`
	DeletedAt		*time.Time	`sql:"index" valid:"-"`
	Plate			string		`gorm:"not null;unique" valid:"required"`
	Model			string		`valid:"-"`
	Owner			Client		`valid:"-" json:"-"`
	OwnerID			uint		`gorm:"not null" valid:"required"`
}


func (v *Vehicle) Insert() {
	DB.Create(&v)
	beego.Debug("Insert Vehicle:", v)
}

func (v *Vehicle) Exists() bool {
	count := 0
	DB.Where("id = ?", v.ID).Find(&Vehicle{}).Count(&count)
	return count > 0
}

func (v *Vehicle) ExistsPlate() bool {
	count := 0
	DB.Unscoped().Where("plate = ? and id <> ?", v.Plate, v.ID).Find(&Vehicle{}).Count(&count)
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
	var vDB Vehicle
	DB.Where("plate = ?", v.Plate).First(&vDB)

	vDB.Model = v.Model
	vDB.OwnerID = v.OwnerID

	DB.Save(&vDB)
	beego.Debug("Update Vehicle:", v)
}

func (v *Vehicle) DeleteSoft() {
	DB.Delete(&v)
	beego.Debug("Delete Vehicle:", v)
}

func (v Vehicle) String() string {
	out, err := json.Marshal(v)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return string(out)
}