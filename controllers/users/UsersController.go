package controllers

import (
	"github.com/astaxie/beego"
	"github.com/Qiaorui/zooli/controllers"
)

type UsersController struct {
	controllers.BaseController
}

func (c *UsersController) NestPrepare() {

}

func (c *UsersController) Index() {
	beego.ReadFromRequest(&c.Controller)
	c.TplName = "users/index.tpl"
}
