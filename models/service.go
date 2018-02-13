package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
	"encoding/json"
)

type Service struct {
	gorm.Model			`valid:"-"`
	Name		string	`gorm:"not null;unique" valid:"required"`
}

func (s *Service) Insert() {
	DB.Create(&s)
	beego.Debug("Insert Service:", s)
}

func (s *Service) Exists() bool {
	count := 0
	DB.Where("id = ?", s.ID).Find(&Service{}).Count(&count)
	return count > 0
}

func (s *Service) ExistsName() bool {
	count := 0
	DB.Where("name = ? and id <> ?", s.Name, s.ID).Find(&Service{}).Count(&count)
	return count > 0

}

func FindServiceByID(id uint) Service {
	var s Service
	DB.Where("id = ?", id).Find(&s)
	return s
}

func FindServiceByName(name string) Service {
	var s Service
	DB.Where("name = ?", name).Find(&s)
	return s
}

func ExistServiceByName(name string) bool {
	var s Service
	count := 0
	DB.Where("name = ?", name).Find(&s).Count(&count)
	return count > 0
}

func FindServices() []Service {
	var s []Service
	DB.Find(&s)
	return s
}

func (s *Service) Update() {
	var sDB Service
	sDB.ID = s.ID
	DB.Where(&sDB).First(&sDB)

	sDB.Name = s.Name

	DB.Save(&sDB)
	beego.Debug("Update Service:", s)
}

func (s *Service) DeleteSoft() {
	DB.Delete(&s)
	beego.Debug("Delete Service:", s)
}


func (s Service) String() string {
	out, err := json.Marshal(s)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return string(out)
}