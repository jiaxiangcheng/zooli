package controllers

import (
	"github.com/Qiaorui/zooli/controllers"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
	"strconv"
)

type UsersController struct {
	controllers.BaseController
}

func (c *UsersController) Get() {
	c.Data["users"] = models.FindUsers()
	c.TplName = "best_practice/users/list.tpl"
}

func (c *UsersController) Edit() {
	id, _ := c.GetInt64("id")
	user := models.FindUserByID(uint(id))
	c.Data["userForm"] = user
	c.Data["roles"] = models.FindRoles()

	c.TplName = "best_practice/users/edit.tpl"
}



func (c *UsersController) New() {
	//c.TplName = "users/create.tpl"

	c.Data["roles"] = models.FindRoles()
	c.TplName = "best_practice/users/new.tpl"
}

/*
func (c *UsersController) ExistUserIf() {
	user_name := c.GetString("username")
	existed := models.ExistUserByUsername(user_name)
	c.Data["json"] = map[string]interface{}{"existed": existed}
	c.ServeJSON()
}
*/

func (c *UsersController) Create() {


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

func (c *UsersController) Update() {
	user_name := c.GetString("username")

	email := c.GetString("email")
	name := c.GetString("name")

	user := models.FindUserByUsername(user_name)
	user.Email = email
	user.Name = name

	user.Update()

	//flash
}


func (c *UsersController) Delete() {
	flash := beego.NewFlash()

	//get identifier of user
	id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	//error if id is not valid
	/*if err != nil {
		flash.Error("Wrong id")
		flash.Store(&c.Controller)
		c.Redirect("/users", 302)
		return
	}*/

	var u models.User
	u.ID = uint(id)
	if !u.Exists() {
		flash.Error("method")
		flash.Store(&c.Controller)
		c.Redirect("/users", 303)
		return
	}

	u.DeleteSoft()

	// load message success and redirect
	flash.Success("You have deleted user")
	flash.Store(&c.Controller)
	c.Redirect("/users", 303)
}