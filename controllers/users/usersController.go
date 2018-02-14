package controllers

import (
	"github.com/Qiaorui/zooli/controllers"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
	"strconv"
	utils "github.com/Qiaorui/zooli/controllers/utils"
)

type UsersController struct {
	controllers.BaseController
}

func (c *UsersController) Get() {
	c.Data["users"] = models.FindUsers()
	c.TplName = "best_practice/users/list.tpl"
}

func (c *UsersController) Edit() {

	userSession := c.GetSession("userInfo")

	var u models.User

	if userSession != nil {
		c.DelSession("userInfo")
		u = userSession.(models.User)
	} else {
		id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
		// load user in DB
		u = models.FindUserByID(uint(id))
	}
	if !u.Exists() {
		flash := beego.NewFlash()
		flash.Error("Incorrect user id")
		flash.Store(&c.Controller)
		c.Redirect("/settings/user/getList", 302)
		return
	}

	c.Data["userForm"] = u
	c.Data["roles"] = models.FindRoles()

	c.TplName = "best_practice/users/edit.tpl"
}



func (c *UsersController) New() {

	//get the user session and load if exist
	u := c.GetSession("userInfo")
	if u != nil {
		c.DelSession("userInfo")
		c.Data["userForm"] = u.(models.User)
	}

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
	flash := beego.NewFlash()

	u, err := c.getUser()
	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("userInfo", u)
		c.Redirect("/users/new", 303)
		return
	}

	err = utils.Validate(u)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("userInfo", u)
		c.Redirect("/users/new", 303)
		return
	}

	u.Insert()

	// load message success and redirect
	flash.Success("You have create the user " + u.Name)
	flash.Store(&c.Controller)
	c.Redirect("/users", 303)
}

func (c *UsersController) Update() {
	//init object for error control
	flash := beego.NewFlash()

	//get identifier of user
	id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	u, err := c.getUser()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/users", 302)
		return
	}

	u.ID = uint(id)
	if !u.Exists() {
		flash.Error("Incorrect user id")
		flash.Store(&c.Controller)
		c.Redirect("/users", 302)
		return
	}

	err = utils.Validate(u)

	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("userInfo", u)
		c.Redirect("/users/" + strconv.Itoa(id), 302)
		return
	}

	//update the user
	u.Update()

	// load message success and redirect
	flash.Success("You have update the user " + u.Name)
	flash.Store(&c.Controller)
	c.Redirect("/users", 302)
}


func (c *UsersController) Delete() {
	flash := beego.NewFlash()

	//get identifier of user
	id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	var u models.User
	u.ID = uint(id)
	if !u.Exists() {
		flash.Error("Incorrect user id")
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

func (c *UsersController) getUser() (models.User, error) {
	u := models.User{
		Username: c.GetString("username"),
		Email: c.GetString("email"),
		Name: c.GetString("name"),
	}
	u.SetPassword(c.GetString("password"))
	roleID, err := c.GetInt("role")
	if err != nil {
		return u, err
	}
	u.RoleID = uint(roleID)

	return u, nil
}