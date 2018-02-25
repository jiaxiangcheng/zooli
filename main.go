package main

import (
	_ "github.com/Qiaorui/zooli/routers"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
	"github.com/asaskevich/govalidator"
	"strconv"
	"os"
	"regexp"
)

func main() {

	if port := os.Getenv("PORT"); port != "" {
		beego.BConfig.Listen.HTTPPort, _ = strconv.Atoi(port)
	}

	govalidator.SetFieldsRequiredByDefault(true)
	if err := models.Connect(); err != nil {
		models.Syncdb()
	}
	//models.DB.LogMode(true)

	beego.Run()


	govalidator.TagMap["phone"] = govalidator.Validator(func(str string) bool {
		match, _ := regexp.MatchString("^[\\d-+]+$", str)
		return match
	})
}
