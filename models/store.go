package models

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

type Store struct {
	gorm.Model					`valid:"-"`
	Name        string			`gorm:"not null" valid:"required"`
	Address     string			`valid:"-"`
	Latitude    float64			`gorm:"not null" valid:"required"`
	Longitude   float64			`gorm:"not null" valid:"required"`
	PhoneNumber string			`valid:"phone,optional"`
	Images		[]StoreImage	`valid:"-"`
	Company     Company			`valid:"-" json:"-"`
	CompanyID   uint			`gorm:"not null" valid:"required"`
	Managers	[]User			`valid:"-" json:"-"`
	Services    []Service		`gorm:"many2many:store_services;" valid:"-" json:"-"`
}

func (s *Store) Insert() {
	DB.Create(&s)
	beego.Debug("Insert Store:", s)
}

func (s *Store) Exists() bool {
	count := 0
	DB.Where("id = ?", s.ID).Find(&Store{}).Count(&count)
	return count > 0
}

func FindStoreByID(id uint) Store {
	var s Store
	DB.Where("id = ?", id).Preload("Images").Preload("Managers").Preload("Services").Preload("Company").Find(&s)
	return s
}

func FindStores() []Store {
	var s []Store
	DB.Preload("Services").Preload("Managers").Preload("Company").Find(&s)
	return s
}

func (s *Store) Update() {
	var sDB Store
	sDB.ID = s.ID
	DB.Where(&sDB).First(&sDB)

	DB.Model(&sDB).Association("Services").Replace(s.Services)
	DB.Model(&sDB).Association("Managers").Replace(s.Managers)
	DB.Model(&sDB).Association("Images").Replace(s.Images)
	sDB.Name = s.Name
	sDB.Address = s.Address
	sDB.Latitude = s.Latitude
	sDB.Longitude = s.Longitude
	sDB.PhoneNumber = s.PhoneNumber
	sDB.CompanyID = s.CompanyID

	DB.Save(&sDB)
	beego.Debug("Update Store:", sDB)
}

func (s *Store) DeleteSoft() {
	DB.Delete(&s)
	beego.Debug("Delete Store:", s)
}

func (s Store) String() string {
	out, err := json.Marshal(s)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return string(out)
}

func NumStores() int {
	var count int = 0
	var st []Store
	DB.Find(&st).Count(&count)
	return count
}
