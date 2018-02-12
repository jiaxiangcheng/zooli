package controllers

import (
	"github.com/Qiaorui/zooli/controllers"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
)

type DashController struct {
	controllers.BaseController
}

func (c *DashController) LoadDashboard() {
	c.TplName = "dashboard/dashboard.tpl"
}
