package controllers

import (
    "github.com/Qiaorui/zooli/controllers"
    "github.com/Qiaorui/zooli/models"
)

type StoresController struct {
	controllers.BaseController
}

func (c *StoresController) Get() {

	c.Data["stores"] = models.FindStores()
	c.TplName = "best_practice/stores/list.tpl"
}
/*
func (c *StoresController) Edit() {

	companysession := c.GetSession("companiesInfo")

	var company models.Company

	if companysession != nil {
		c.DelSession("companiesInfo")
		company = companysession.(models.Company)
	} else {
		id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
		// load company in DB
		company = models.FindCompanyByID(uint(id))
	}
	if !company.Exists() {
		flash := beego.NewFlash()
		flash.Error("Incorrect company id")
		flash.Store(&c.Controller)
		c.Redirect("/settings/company/getList", 302)
		return
	}

	c.Data["companyForm"] = company
    c.Data["headerTitle"] = "Company Information"

	c.TplName = "best_practice/companies/edit.tpl"
}

func (c *StoresController) New() {

	//get the company session and load if exist
	company := c.GetSession("companiesInfo")
	if company != nil {
		c.DelSession("companiesInfo")
		c.Data["companyForm"] = company.(models.Company)
	}

    c.Data["headerTitle"] = "New Company"
	c.TplName = "best_practice/companies/new.tpl"
}

func (c *StoresController) Create() {
	flash := beego.NewFlash()

	company, err := c.getCompany()
	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("companiesInfo", company)
		c.Redirect("/companies/new", 303)
		return
	}

	err = utils.Validate(company)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("companiesInfo", company)
		c.Redirect("/companies/new", 303)
		return
	}

	company.Insert()

	// load message success and redirect
	flash.Success("You have create the company " + company.Name)
	flash.Store(&c.Controller)
	c.Redirect("/companies", 303)
}

func (c *StoresController) Update() {
	//init object for error control
	flash := beego.NewFlash()

	//get identifier of company
	id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	company, err := c.getCompany()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/companies/" + strconv.Itoa(id), 302)
		return
	}

	company.ID = uint(id)
	if !company.Exists() {
		flash.Error("Incorrect company id")
		flash.Store(&c.Controller)
		c.Redirect("/companies", 302)
		return
	}

	err = utils.Validate(company)

	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("companiesInfo", company)
		c.Redirect("/companies/" + strconv.Itoa(id), 302)
		return
	}

	//update the company
	company.Update()

	// load message success and redirect
	flash.Success("You have update the company " + company.Name)
	flash.Store(&c.Controller)
	c.Redirect("/companies/" + strconv.Itoa(id), 302)
}

func (c *StoresController) Delete() {
	flash := beego.NewFlash()

	//get identifier of company
	id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	var company models.Company
	company.ID = uint(id)
	if !company.Exists() {
		flash.Error("Incorrect company id")
		flash.Store(&c.Controller)
		c.Redirect("/companies", 303)
		return
	}

	company.DeleteSoft()

	// load message success and redirect
	flash.Success("You have deleted company")
	flash.Store(&c.Controller)
	c.Redirect("/companies", 303)
}
*/
func (c *StoresController) getCompany() (models.Company, error) {
	company := models.Company{
		Name: c.GetString("name"),
        Contact: c.GetString("contact"),
		Email: c.GetString("email"),
		PhoneNumber: c.GetString("phonenumber"),
	}

	return company, nil
}