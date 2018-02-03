package controllers

import models "zooli/models/users"

type LoginController struct {
}

func (this *LoginController) Login() {

	if this.IsLogin {
		this.BaseController.Ctx.Redirect(302, this.UrlFor("UsersController.Index"))
		return
	}

	this.TplNames = "auth/login.tpl"
	username := c.GetString("Username")
	password := c.GetString("Password")

	user, err := this.Authenticate(username, password)

	if err != nil {

		return
	}

	this.SetUserSession(user)
	this.Redirect(c.UrlFor("UsersController.Index"), 303)
}

func (this *LoginController) Authenticate(username string, password string) (user *models.User, err error) {
	user = &models.User{Username: username}

	// TODO: validate user
	return user, nil

}

func (this *LoginController) Logout() {
	this.DelUserSession()
	this.Ctx.Redirect(302, c.UrlFor("LoginController.Login"))
}
