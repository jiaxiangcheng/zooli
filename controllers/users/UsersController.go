package controllers

import (
	"github.com/astaxie/beego"
)

type UsersController struct {
	BaseController
}

func (this *UsersController) NestPrepare() {
	if !this.IsLogin {
		this.Ctx.Redirect(302, this.GoLogin())
	}
}

func (this *UsersController) Index() {
	beego.ReadFromRequest(&this.Controller)
	c.TplNames = "users/index.tpl"
}
