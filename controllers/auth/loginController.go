package controllers

import (
	"github.com/Qiaorui/zooli/controllers"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	controllers.BaseController
}

func (c *LoginController) LoginForm() {
	c.Layout = "best_practice/layout.tpl"
	c.TplName = "best_practice/login.tpl"
}

func (c *LoginController) Login() {
	flash := beego.NewFlash()
	u := models.FindUserByUsername(c.GetString("username"))
	if u.ValidPassword(c.GetString("password")) {
		c.SetSession("user", u)
		c.Redirect("/dashboard", 302)
	} else {
		flash.Error("Wrong username password combination")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
	}
}

func (c *LoginController) Logout() {
	c.DelSession("user")
	flash := beego.NewFlash()
	flash.Success("Log out successfully")
	flash.Store(&c.Controller)
	c.Redirect("/login", 302)
}
