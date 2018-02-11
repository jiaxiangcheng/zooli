package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
)

type Role struct {
	gorm.Model			`valid:"-"`
	Name	string		`valid:"-"`
}


func (r *Role) Insert() {
	beego.Debug("Insert ", r)
	DB.Create(&r)
}

func (r *Role) Exists() bool {
	count := 0
	DB.Where("id = ?", r.ID).Find(&Role{}).Count(&count)
	return count > 0
}

func (r *Role) ExistsName() bool {
	count := 0
	DB.Where("name = ? and id <> ?", r.Name, r.ID).Find(&Role{}).Count(&count)
	return count > 0

}

func FindRoleByID(id uint) Role {
	var r Role
	DB.Where("id = ?", id).Find(&r)
	return r
}

func FindRoleByName(name string) Role {
	var r Role
	DB.Where("name = ?", name).Find(&r)
	return r
}

func ExistRoleByName(name string) bool {
	var r Role
	count := 0
	DB.Where("name = ?", name).Find(&r).Count(&count)
	return count > 0
}

func FindRoles() []Role {
	var r []Role
	DB.Find(&r)
	return r
}

func (r *Role) Update() {
	beego.Debug("Update ", r)
	var rDB Role
	rDB.ID = r.ID
	DB.Where(&rDB).First(&rDB)

	rDB.Name = r.Name

	DB.Save(&rDB)
}

func (r *Role) DeleteSoft() {
	beego.Debug("Update ", r)
	DB.Delete(&r)
}


