package controllers

import (
	"fmt"

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
	c.Layout = "users/index.tpl"
	c.TplName = "users/header.html"
}

func (c *UsersController) LoadUser() {
	id, _ := c.GetInt64("id")
	user := models.FindUserByID(uint(id))
	c.Data["UserInfo"] = user
	fmt.Print(c.Data["UserInfo"])

	c.Redirect("/users/"+c.GetString("id"), 302)
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
