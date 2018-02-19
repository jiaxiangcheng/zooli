package business

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
	c.Data["stores"] = models.FindStores()
	c.TplName = "stores/list.tpl"
}

func (c *StoresController) Edit() {
	storeSession := c.GetSession("storeInfo")
	var s models.Store

	if storeSession != nil {
		c.DelSession("storeInfo")
		s = storeSession.(models.Store)
	} else {
		id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
		s = models.FindStoreByID(uint(id))
	}

	if !s.Exists() {
		flash := beego.NewFlash()
		flash.Error("Incorrect store id")
		flash.Store(&c.Controller)
		c.Redirect("/stores", 302)
		return
	}

	c.Data["companies"] = models.FindCompanies()
	c.Data["services"] = models.FindServices()
	c.Data["storeForm"] = s
	c.TplName = "stores/edit.tpl"
}

func (c *StoresController) findManagersWithoutStoreAssigned() []models.User {
	//var managers []models.User
	var usersWithoutStore []models.User

	// users which has role managar and doesn't have any store assigned
	//managers = models.FindUsersByRole(2)
	//usersWithoutStore = models.FindManagersWithoutStore()

	return usersWithoutStore
}

func (c *StoresController) New() {
	s := c.GetSession("storeInfo")
	if s != nil {
		c.DelSession("storeInfo")
		c.Data["storeForm"] = s.(models.Store)
	}

	c.Data["companies"] = models.FindCompanies()
	c.Data["services"] = models.FindServices()
	//c.Data["managers"] = c.findManagersWithouStoreAssgined()
	c.TplName = "stores/new.tpl"
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

	// relate store with user
	//var user = models.FindUserByID(store.ManagerID)
	//user.Store = store
	//user.StoreID = store.ID
	//user.Update()

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

	// update old manager if exists
	/*
	var oldStore = models.FindStoreByID(uint(id))
	if oldStore.ManagerID != 0 {
		var oldManager = models.FindUserByID(oldStore.ManagerID)
		oldManager.StoreID = 0
		oldManager.Update()
	}
	*/

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

	// relate user
	/*
	var user = models.FindUserByID(store.ManagerID)

	user.StoreID = store.ID
	user.Update()
	*/

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

	store.CompanyID = uint(companyID)

	return store, nil
}
