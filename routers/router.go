package routers

import (
	"github.com/Qiaorui/zooli/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
