package controllers

import (
	"github.com/Qiaorui/zooli/controllers"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
)

type UsersController struct {
	controllers.BaseController
}

func (c *UsersController) Prepare() {
	//// Overwrite beego.Controller.Layout (string)
	_ = beego.ReadFromRequest(&c.Controller)

	u := c.GetSession("user")
	c.Data["user"] = u
	c.Data["roleManager"] = models.ROLE_MANAGER
	c.Data["roleAdmin"] = models.ROLE_ADMIN

}

func (c *UsersController) LoadUser() {
	id, _ := c.GetInt64("id")
	user := models.FindUserByID(uint(id))
	c.Data["UserInfo"] = user
	beego.Info(c.Data["UserInfo"])

	c.Layout = "common/content.html"
	c.TplName = "users/user.tpl"
}

func (c *UsersController) DeleteUser() {
	id, _ := c.GetInt64("id")
	user := models.FindUserByID(uint(id))
	user.DeleteSoft()
	c.Get()
}

func (c *UsersController) CreateUser() {
	c.Layout = "common/content.html"
	c.TplName = "users/create.tpl"
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
	email := c.GetString("email")
	name := c.GetString("name")

	new_user := models.User{
		Username:     user_name,
		PasswordHash: password,
		Email:        email,
		Name:         name,
		RoleID:       uint(1)}

	new_user.SetPassword(password)
	new_user.Insert()
	c.Redirect("/users", 302)
}

func (c *UsersController) SaveUser() {
	user_name := c.GetString("username")

	email := c.GetString("email")
	name := c.GetString("name")

	user := models.FindUserByUsername(user_name)
	user.Email = email
	user.Name = name

	user.Update()

	//flash
}
