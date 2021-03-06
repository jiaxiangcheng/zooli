package auth

import (
	"github.com/Qiaorui/zooli/controllers"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	utils "github.com/Qiaorui/zooli/controllers/utils"
)

type LoginController struct {
	controllers.BaseController
}

func (c *LoginController) LoginForm() {
	if u := c.GetSession("userInfo"); u != nil {
		c.DelSession("userInfo")
		c.Data["userForm"] = u.(models.User)
	}

	c.Layout = "layout.tpl"
	c.TplName = "login.tpl"
}

func (c *LoginController) Login() {
	flash := beego.NewFlash()
	if !models.ExistUserByUsername(c.GetString("username")) {
		u := models.User{
			Username: c.GetString("username"),
		}
		c.SetSession("userInfo", u)
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_LOGIN_COMBO))
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
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_LOGIN_COMBO))
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
	}
}

func (c *LoginController) Logout() {
	c.DelSession("user")
	flash := beego.NewFlash()
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_LOG_OUT))
	flash.Store(&c.Controller)
	c.Redirect("/login", 302)
}
