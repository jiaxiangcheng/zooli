package controllers

import (
	"github.com/Qiaorui/zooli/controllers"
    "github.com/Qiaorui/zooli/models"
    "github.com/astaxie/beego"
    utils "github.com/Qiaorui/zooli/controllers/utils"
)

type CompaniesController struct {
	controllers.BaseController
}

func (c *CompaniesController) Get() {
	c.Data["companies"] = models.FindCompanies()
	c.TplName = "best_practice/companies/list.tpl"
}
/*
func (c *CompaniesController) Edit() {

	companiesession := c.GetSession("companiesInfo")

	var u models.User

	if companiesession != nil {
		c.DelSession("companiesInfo")
		u = companiesession.(models.User)
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

	c.TplName = "best_practice/companies/edit.tpl"
}
*/


func (c *CompaniesController) New() {

	//get the user session and load if exist
	company := c.GetSession("companiesInfo")
	if company != nil {
		c.DelSession("companiesInfo")
		c.Data["companyForm"] = company.(models.Company)
	}

    c.Data["headerTitle"] = "New Company"
	c.TplName = "best_practice/companies/new.tpl"
}

/*
func (c *CompaniesController) ExistUserIf() {
	user_name := c.GetString("username")
	existed := models.ExistUserByUsername(user_name)
	c.Data["json"] = map[string]interface{}{"existed": existed}
	c.ServeJSON()
}*/

func (c *CompaniesController) Create() {
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
/*
func (c *CompaniesController) Update() {
	//init object for error control
	flash := beego.NewFlash()

	//get identifier of user
	id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	u, err := c.getUser()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/companies/" + strconv.Itoa(id), 302)
		return
	}

	u.ID = uint(id)
	if !u.Exists() {
		flash.Error("Incorrect user id")
		flash.Store(&c.Controller)
		c.Redirect("/companies", 302)
		return
	}

	err = utils.Validate(u)

	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("companiesInfo", u)
		c.Redirect("/companies/" + strconv.Itoa(id), 302)
		return
	}

	//update the user
	u.Update()

	// load message success and redirect
	flash.Success("You have update the user " + u.Name)
	flash.Store(&c.Controller)
	c.Redirect("/companies/" + strconv.Itoa(id), 302)
}

func (c *CompaniesController) Delete() {
	flash := beego.NewFlash()

	//get identifier of user
	id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	var u models.User
	u.ID = uint(id)
	if !u.Exists() {
		flash.Error("Incorrect user id")
		flash.Store(&c.Controller)
		c.Redirect("/companies", 303)
		return
	}

	u.DeleteSoft()

	// load message success and redirect
	flash.Success("You have deleted user")
	flash.Store(&c.Controller)
	c.Redirect("/companies", 303)
}*/

func (c *CompaniesController) getCompany() (models.Company, error) {
	company := models.Company{
		Name: c.GetString("name"),
        Contact: c.GetString("contact"),
		Email: c.GetString("email"),
		PhoneNumber: c.GetString("phonenumber"),
	}

	return company, nil
}
