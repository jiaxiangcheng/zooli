package controllers

import 	"github.com/Qiaorui/zooli/models"

type MainController struct {
	BaseController
}

func (c *MainController) Prepare() {
	c.BaseController.Prepare()
}

func (c *MainController) Get() {
    c.Data["usercount"] = models.NumUsers()
	c.TplName = "dashboard.tpl"
}
