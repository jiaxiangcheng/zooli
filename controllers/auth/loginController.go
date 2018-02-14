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
	if u := c.GetSession("userInfo"); u != nil {
		c.DelSession("userInfo")
		c.Data["userForm"] = u.(models.User)
	}

	c.Layout = "best_practice/layout.tpl"
	c.TplName = "best_practice/login.tpl"
}

func (c *LoginController) Login() {
	flash := beego.NewFlash()
	if !models.ExistUserByUsername(c.GetString("username")) {
		u := models.User{
			Username: c.GetString("username"),
		}
		c.SetSession("userInfo", u)
		flash.Error("Wrong username password combination")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}
	u := models.FindUserByUsername(c.GetString("username"))
	if u.ValidPassword(c.GetString("password")) {
		c.SetSession("user", u)
		c.Redirect("/dashboard", 302)
	} else {
		c.SetSession("userInfo", u)
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
