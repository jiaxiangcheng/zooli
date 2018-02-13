package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Prepare() {
	c.BaseController.Prepare()
}

func (c *MainController) Get() {

	c.TplName = "best_practice/dashboard.tpl"
}
