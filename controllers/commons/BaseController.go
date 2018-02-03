package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	// Embedding: "Inherit" beego.Controller
	beego.Controller
	Userinfo *models.User
	IsLogin  bool
}

func (this *BaseController) Prepare() {

	//// Overwrite beego.Controller.Layout (string)
	this.Data["user"] = this.GetSession("user")
	this.IsLogin = this.Data["user"] != nil
	if this.IsLogin {
		this.Userinfo = this.GetUserSession()
		this.Data["IsLogin"] = c.IsLogin
		this.Data["Userinfo"] = c.Userinfo
		this.Layout = "auth/login.tpl"
}

func (this *BaseController) GetUserSession *models.User {
	user := &models.User{Username: this.GetSession("user").(string)}
	return user
}

func (this *BaseController) SetUserSession(user *models.User) {
	this.SetSession("user", user.Username);
}

func (this *BaseController) DelUserSession() {
	this.DelSession("user")
}

func (this *BaseController) GoLogin() {
	return this.UrlFor("LoginController.Login")
}
