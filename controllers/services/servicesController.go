package controllers

import (
    "github.com/Qiaorui/zooli/controllers"
    "github.com/Qiaorui/zooli/models"
    "github.com/astaxie/beego"
    "strconv"
    utils "github.com/Qiaorui/zooli/controllers/utils"
)

type ServicesController struct {
	controllers.BaseController
}

func (c *ServicesController) Get() {
	c.Data["services"] = models.FindServices()
	c.TplName = "services/list.tpl"
}

func (c *ServicesController) Edit() {

	serviceSession := c.GetSession("servicesInfo")

	var service models.Service

	if serviceSession != nil {
		c.DelSession("serviceInfo")
		service = serviceSession.(models.Service)
	} else {
		id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
		// load service in DB
		service = models.FindServiceByID(uint(id))
	}
	if !service.Exists() {
		flash := beego.NewFlash()
		flash.Error("Incorrect service id")
		flash.Store(&c.Controller)
		c.Redirect("/services", 302)
		return
	}

	c.Data["serviceForm"] = service
	c.TplName = "services/edit.tpl"
}

func (c *ServicesController) New() {

	//get the service session and load if exist
	service := c.GetSession("serviceInfo")
	if service != nil {
		c.DelSession("serviceInfo")
		c.Data["serviceForm"] = service.(models.Service)
	}

	c.TplName = "services/new.tpl"
}

func (c *ServicesController) Create() {
	flash := beego.NewFlash()

	service, err := c.getService()
	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("serviceInfo", service)
		c.Redirect("/services/new", 303)
		return
	}

	err = utils.Validate(service)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("serviceInfo", service)
		c.Redirect("/services/new", 303)
		return
	}

	service.Insert()

	// load message success and redirect
	flash.Success("You have created the service " + service.Name)
	flash.Store(&c.Controller)
	c.Redirect("/services", 303)
}

func (c *ServicesController) Update() {
	//init object for error control
	flash := beego.NewFlash()

	//get identifier of service
	id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	service, err := c.getService()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/services/" + strconv.Itoa(id), 302)
		return
	}

	service.ID = uint(id)
	if !service.Exists() {
		flash.Error("Incorrect service id")
		flash.Store(&c.Controller)
		c.Redirect("/services", 302)
		return
	}

	err = utils.Validate(service)

	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("serviceInfo", service)
		c.Redirect("/services/" + strconv.Itoa(id), 302)
		return
	}

	//update the service
	service.Update()

	// load message success and redirect
	flash.Success("You have updated the service " + service.Name)
	flash.Store(&c.Controller)
	c.Redirect("/services/" + strconv.Itoa(id), 302)
}


func (c *ServicesController) Delete() {
	flash := beego.NewFlash()

	//get identifier of service
	id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	var service models.Service
	service.ID = uint(id)
	if !service.Exists() {
		flash.Error("Incorrect service id")
		flash.Store(&c.Controller)
		c.Redirect("/services", 303)
		return
	}

	service.DeleteSoft()

	// load message success and redirect
	flash.Success("You have deleted service")
	flash.Store(&c.Controller)
	c.Redirect("/services", 303)
}

func (c *ServicesController) getService() (models.Service, error) {
	service := models.Service{
		Name: c.GetString("name"),
	}
	return service, nil
}
