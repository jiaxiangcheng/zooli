package controllers

import (
	"strconv"

	"github.com/Qiaorui/zooli/controllers"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	controllers.BaseController
}

func (c *UserController) Get() {
	id := c.Ctx.Input.Param(":id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err)
	}
	user := models.FindUserByID(uint(i))
	c.Data["UserInfo"] = user
	c.TplName = "users/user.tpl"
}

func (c *UserController) ExistUserIf() {
	user_name := c.GetString("username")
	existed := models.ExistUserByUsername(user_name)
	c.Data["json"] = map[string]interface{}{"existed": existed}
	c.ServeJSON()
}

func (c *UserController) RegisterUser() {
	flash := beego.NewFlash()

	user_name := c.GetString("username")
	password := c.GetString("password")

	new_user := models.User {Username: user_name, PasswordHash: password}

	//flash.Error("Username already exists")

	new_user.Insert()
	flash.Success("Successfully created user")
	flash.Store(&c.Controller)

	c.Redirect("/users", 302)

	/*if (!new_user.ExistsUsername()) {
		new_user.Insert()
		c.Redirect("/users", 302)
	}
	c.Redirect("/users/register", 302)*/
}
