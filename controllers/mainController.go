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
	c.Data["companycount"] = models.NumCompanies()
	c.Data["servicecount"] = models.NumServices()
	c.Data["storecount"] = models.NumStores()
	c.TplName = "dashboard.tpl"
}

func (c *MainController) RandomData() {
	models.GenerateRandomDataset()
	c.Redirect("/dashboard", 302)
}
