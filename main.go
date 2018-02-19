package main

import (
	_ "github.com/Qiaorui/zooli/routers"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
	"github.com/asaskevich/govalidator"
)

func main() {

	govalidator.SetFieldsRequiredByDefault(true)
	if err := models.Connect(); err != nil {
		models.Syncdb()
	}
	models.DB.LogMode(true)

	beego.Run()
}
