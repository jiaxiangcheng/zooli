package admin

import (
	"strconv"

	"github.com/Qiaorui/zooli/controllers"
	utils "github.com/Qiaorui/zooli/controllers/utils"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type CompaniesController struct {
	controllers.BaseController
}

func (c *CompaniesController) Get() {
	c.Data["companies"] = models.FindCompanies()
	c.TplName = "admin/companies/list.tpl"
}

func (c *CompaniesController) Edit() {

	companySession := c.GetSession("companyInfo")

	var company models.Company

	if companySession != nil {
		c.DelSession("companyInfo")
		company = companySession.(models.Company)
	} else {
		id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
		// load company in DB
		company = models.FindCompanyByID(uint(id))
	}
	if !company.Exists() {
		flash := beego.NewFlash()
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_COMPANY_AREADY_EXISTS))
		flash.Store(&c.Controller)
		c.Redirect("/admin/companies", 302)
		return
	}

	c.Data["companyForm"] = company

	c.TplName = "admin/companies/edit.tpl"
}

func (c *CompaniesController) New() {

	//get the company session and load if exist
	company := c.GetSession("companyInfo")
	if company != nil {
		c.DelSession("companyInfo")
		c.Data["companyForm"] = company.(models.Company)
	}

	c.TplName = "admin/companies/new.tpl"
}

func (c *CompaniesController) Create() {
	flash := beego.NewFlash()

	company, err := c.getCompany()
	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("companiesInfo", company)
		c.Redirect("/admin/companies/new", 303)
		return
	}

	err = utils.Validate(company)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("companyInfo", company)
		c.Redirect("/admin/companies/new", 303)
		return
	}

	company.Insert()

	// load message success and redirect
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_CREATE_COMPANY) + " " + company.Name)
	flash.Store(&c.Controller)
	c.Redirect("/admin/companies", 303)
}

func (c *CompaniesController) Update() {
	//init object for error control
	flash := beego.NewFlash()

	//get identifier of company
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	company, err := c.getCompany()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/admin/companies/"+strconv.Itoa(id), 302)
		return
	}

	company.ID = uint(id)
	if !company.Exists() {
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_COMPANY_AREADY_EXISTS))
		flash.Store(&c.Controller)
		c.Redirect("/admin/companies", 302)
		return
	}

	err = utils.Validate(company)

	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("companyInfo", company)
		c.Redirect("/admin/companies/"+strconv.Itoa(id), 302)
		return
	}

	//update the company
	company.Update()

	// load message success and redirect
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_UPDATE_COMPANY) + " " + company.Name)
	flash.Store(&c.Controller)
	c.Redirect("/admin/companies/"+strconv.Itoa(id), 302)
}

func (c *CompaniesController) Delete() {
	flash := beego.NewFlash()

	//get identifier of company
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	var company models.Company
	company.ID = uint(id)
	if !company.Exists() {
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_COMPANY_AREADY_EXISTS))
		flash.Store(&c.Controller)
		c.Redirect("/admin/companies", 303)
		return
	}

	company.DeleteSoft()

	// load message success and redirect
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_DELETE_COMPANY) + " " + company.Name)
	flash.Store(&c.Controller)
	c.Redirect("/admin/companies", 303)
}

func (c *CompaniesController) getCompany() (models.Company, error) {
	company := models.Company{
		Name:        c.GetString("name"),
		Contact:     c.GetString("contact"),
		Email:       c.GetString("email"),
		PhoneNumber: c.GetString("phoneNumber"),
	}

	return company, nil
}
