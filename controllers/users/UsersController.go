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
	c.TplName = "users/index.tpl"
}

func (c *UsersController) LoadUser() {
	id, _ := c.GetInt64("id")
	userInfo := models.FindUserByID(uint(id))
	c.Data["UserInfo"] = userInfo
	c.TplName = "users/user.tpl"
}

func (c *UsersController) DeleteUser(username string) {

}

func (c *UsersController) UpdateUser(username string) {

}
