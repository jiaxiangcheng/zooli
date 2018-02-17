package controllers

import (
	"strconv"

	"github.com/Qiaorui/zooli/controllers"
	utils "github.com/Qiaorui/zooli/controllers/utils"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
)

type StoresController struct {
	controllers.BaseController
}

func (c *StoresController) Get() {
	//c.Data["stores"] = models.FindStores()
	c.TplName = "best_practice/stores/list.tpl"
}

func (c *StoresController) Edit() {
	var store models.Store

	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	// load store in DB
	store = models.FindStoreByID(uint(id))

	if !store.Exists() {
		flash := beego.NewFlash()
		flash.Error("Incorrect store id")
		flash.Store(&c.Controller)
		c.Redirect("/settings/store/getList", 302)
		return
	}

	c.Data["storeForm"] = store
	c.Data["headerTitle"] = "Store Information"

	c.TplName = "best_practice/store/edit.tpl"
}

func (c *StoresController) findManagers() []models.User {
	var managers []models.User
	managers = models.FindUsersByRole(2)
	return managers
}

func (c *StoresController) New() {
	c.Data["companies"] = models.FindCompanies()
	c.Data["managers"] = c.findManagers()

	c.Data["headerTitle"] = "New Store"
	c.TplName = "best_practice/stores/new.tpl"
}

func (c *StoresController) Create() {
	flash := beego.NewFlash()

	store, err := c.getStore()
	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/stores/new", 303)
		return
	}

	err = utils.Validate(store)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/stores/new", 303)
		return
	}

	store.Insert()

	// load message success and redirect
	flash.Success("You have create the stores " + store.Name)
	flash.Store(&c.Controller)
	c.Redirect("/stores", 303)
}

func (c *StoresController) Update() {
	//init object for error control
	flash := beego.NewFlash()

	//get identifier of store
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	store, err := c.getStore()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/stores/"+strconv.Itoa(id), 302)
		return
	}

	store.ID = uint(id)
	if !store.Exists() {
		flash.Error("Incorrect store id")
		flash.Store(&c.Controller)
		c.Redirect("/stores", 302)
		return
	}

	err = utils.Validate(store)

	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/stores/"+strconv.Itoa(id), 302)
		return
	}

	//update the store
	store.Update()

	// load message success and redirect
	flash.Success("You have update the store " + store.Name)
	flash.Store(&c.Controller)
	c.Redirect("/stores/"+strconv.Itoa(id), 302)
}

func (c *StoresController) Delete() {
	flash := beego.NewFlash()

	//get identifier of store
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	var store models.Store
	store.ID = uint(id)
	if !store.Exists() {
		flash.Error("Incorrect store id")
		flash.Store(&c.Controller)
		c.Redirect("/stores", 303)
		return
	}

	store.DeleteSoft()

	// load message success and redirect
	flash.Success("You have deleted store")
	flash.Store(&c.Controller)
	c.Redirect("/stores", 303)
}

func (c *StoresController) getStore() (models.Store, error) {

	store := models.Store{
		Address:     c.GetString("address"),
		Name:        c.GetString("name"),
		PhoneNumber: c.GetString("phone"),
	}

	latitude, err := c.GetFloat("latitude")
	beego.Debug(latitude)

	if err != nil {
		return store, err
	}

	store.Latitude = latitude

	longitude, err := c.GetFloat("longitude")
	beego.Debug(longitude)

	if err != nil {
		return store, err
	}

	store.Longitude = longitude

	companyID, err := c.GetInt("company")
	if err != nil {
		return store, err
	}

	managerID, err := c.GetInt("manager")
	if err != nil {
		return store, err
	}

	store.CompanyID = uint(companyID)
	store.ManagerID = uint(managerID)

	return store, nil
}
