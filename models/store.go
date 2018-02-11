package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
)

type Store struct {
	gorm.Model				`valid:"-"`
	Name		string		`valid:"-"`
	Address		string		`valid:"-"`
	Latitude	float64		`valid:"latitude"`
	Longitude	float64		`valid:"longitude"`
	PhoneNumber	string		`valid:"numeric,optional"`
	Image		string		`valid:"url,optional"`
	Company		Company		`valid:"-"`
	CompanyID	uint		`valid:"-"`
	ManagerID	uint		`valid:"-"`
	Services	[]Service	`gorm:"many2many:store_services;" valid:"-"`
}


func (s *Store) Insert() {
	beego.Debug("Insert ", s)
	DB.Create(&s)
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
	beego.Debug("Update ", s)
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
}

func (s *Store) DeleteSoft() {
	beego.Debug("Update ", s)
	DB.Delete(&s)
}
