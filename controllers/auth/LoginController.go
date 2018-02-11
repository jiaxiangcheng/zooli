package controllers

import (
	"github.com/Qiaorui/zooli/controllers"
	"github.com/Qiaorui/zooli/models"
)

type LoginController struct {
	controllers.BaseController
}

func (c *LoginController) LoginForm() {
	c.TplName = "auth/login.tpl"
}

func (c *LoginController) Login() {

	u := models.FindUserByUsername(c.GetString("Username"))
	if u.ValidPassword(c.GetString("PasswordHash")) {
		c.SetSession("user", u)
		// role admin
		c.Redirect("/users", 302)

		// role storer
		// TODO: clients view

	} else {
		c.Redirect("/login", 302)
	}
}

func (c *LoginController) Logout() {
	c.DelSession("user")
	c.Redirect("/", 302)
}
