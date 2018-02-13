package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
	"encoding/json"
)

type Store struct {
	gorm.Model				`valid:"-"`
	Name		string		`gorm:"not null" valid:"required"`
	Address		string		`valid:"-"`
	Latitude	float64		`gorm:"not null" valid:"latitude,required"`
	Longitude	float64		`gorm:"not null" valid:"longitude,required"`
	PhoneNumber	string		`valid:"numeric,optional"`
	Image		string		`valid:"url,optional"`
	Company		Company		`valid:"-" json:"-"`
	CompanyID	uint		`gorm:"not null" valid:"required"`
	ManagerID	uint		`valid:"-"`
	Services	[]Service	`gorm:"many2many:store_services;" valid:"-" json:"-"`
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
	DB.Where("id = ?", id).Find(&s)
	return s
}

func FindStores() []Store {
	var s []Store
	DB.Find(&s)
	return s
}

func (s *Store) Update() {
	var sDB Store
	sDB.ID = s.ID
	DB.Where(&sDB).First(&sDB)

	sDB.Name = s.Name
	sDB.Address = s.Address
	sDB.Latitude = s.Latitude
	sDB.Longitude = s.Longitude
	sDB.PhoneNumber = s.PhoneNumber
	sDB.Image = s.Image
	sDB.CompanyID = s.CompanyID
	sDB.ManagerID = s.ManagerID
	DB.Model(&sDB).Association("Services").Replace(s.Services)

	DB.Save(&sDB)
	beego.Debug("Update Store:", s)
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