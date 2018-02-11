package controllers

import (
	"github.com/Qiaorui/zooli/controllers"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
)

type UsersController struct {
	controllers.BaseController
}

func (c *UsersController) NestPrepare() {

}

func (c *UsersController) Index() {
	beego.ReadFromRequest(&c.Controller)

	c.Data["users"] = models.FindUsers()
	c.TplName = "users/user_list.tpl"
}

func (c *UsersController) LoadUser() {
	id, _ := c.GetInt64("id")
	user := models.FindUserByID(uint(id))
	c.Data["UserInfo"] = user

	c.TplName = "users/user.tpl"
}

func (c *UsersController) DeleteUser() {
	id, _ := c.GetInt64("id")
	user := models.FindUserByID(uint(id))
	user.DeleteSoft()
	c.Index()
}

func (c *UsersController) CreateUser() {
	c.TplName = "users/new_user.tpl"
}

func (c *UsersController) ExistUserIf() {
	user_name := c.GetString("username")
	existed := models.ExistUserByUsername(user_name)
	c.Data["json"] = map[string]interface{}{"existed": existed}
	c.ServeJSON()
}

func (c *UsersController) InsertUser() {
	user_name := c.GetString("username")
	password := c.GetString("password")
	new_user := models.User{Username: user_name, PasswordHash: password}

	new_user.SetPassword(password)
	new_user.Insert()
	c.Redirect("/users", 302)
}
