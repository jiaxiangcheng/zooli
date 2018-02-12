package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
	"encoding/json"
)

const ROLE_ADMIN = "admin"
const ROLE_MANAGER = "manager"

type Role struct {
	gorm.Model			`valid:"-"`
	Name	string		`gorm:"not null;unique" valid:"-"`
}

func (r *Role) Insert() {
	DB.Create(&r)
	beego.Debug("Insert Role:", r)
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

	var rDB Role
	rDB.ID = r.ID
	DB.Where(&rDB).First(&rDB)

	rDB.Name = r.Name

	DB.Save(&rDB)
	beego.Debug("Update Role:", rDB)
}

func (r *Role) DeleteSoft() {
	DB.Delete(&r)
	beego.Debug("Delete Role:", r)
}

func (r Role) String() string {
	out, err := json.Marshal(r)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return string(out)
}