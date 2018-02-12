package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Prepare() {
	c.BaseController.Prepare()
}

func (c *MainController) Get() {
	c.Layout = "common/content.html"
	c.TplName = "dashboard/dashboard.tpl"
}
