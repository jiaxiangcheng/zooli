package main

import (
	_ "github.com/Qiaorui/zooli/routers"
	_ "github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
	"github.com/asaskevich/govalidator"
)

func main() {


	govalidator.SetFieldsRequiredByDefault(true)

	beego.Run()
}

