package admin

import (
	"strconv"

	"strings"

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
	c.TplName = "admin/stores/list.tpl"
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
		c.Redirect("/admin/stores", 302)
		return
	}

	m := models.FindRoleByName(models.ROLE_MANAGER)
	managers := models.FindUsersByRoleID(m.ID)
	c.Data["managers"] = managers
	c.Data["companies"] = models.FindCompanies()
	c.Data["services"] = models.FindServices()

	c.Data["storeForm"] = s
	c.TplName = "admin/stores/edit.tpl"
}

func (c *StoresController) New() {
	s := c.GetSession("storeInfo")
	if s != nil {
		c.DelSession("storeInfo")
		c.Data["storeForm"] = s.(models.Store)
	}

	m := models.FindRoleByName(models.ROLE_MANAGER)
	managers := models.FindUsersByRoleID(m.ID)
	c.Data["managers"] = managers
	c.Data["companies"] = models.FindCompanies()
	c.Data["services"] = models.FindServices()
	c.TplName = "admin/stores/new.tpl"
}

func (c *StoresController) Create() {
	flash := beego.NewFlash()

	store, err := c.getStore()
	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("storeInfo", store)
		c.Redirect("/admin/stores/new", 303)
		return
	}

	err = utils.Validate(store)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("storeInfo", store)
		c.Redirect("/admin/stores/new", 303)
		return
	}

	store.Insert()

	// load message success and redirect
	flash.Success("You have create the stores " + store.Name)
	flash.Store(&c.Controller)
	c.Redirect("/admin/stores", 303)
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
		c.Redirect("/admin/stores/"+strconv.Itoa(id), 302)
		return
	}

	store.ID = uint(id)
	if !store.Exists() {
		flash.Error("Incorrect store id")
		flash.Store(&c.Controller)
		c.Redirect("/admin/stores", 302)
		return
	}

	err = utils.Validate(store)

	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("storeInfo", store)
		c.Redirect("/admin/stores/"+strconv.Itoa(id), 302)
		return
	}

	//update the store
	store.Update()

	// load message success and redirect
	flash.Success("You have update the store " + store.Name)
	flash.Store(&c.Controller)
	c.Redirect("/admin/stores/"+strconv.Itoa(id), 302)
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
		c.Redirect("/admin/stores", 303)
		return
	}

	store.DeleteSoft()

	// load message success and redirect
	flash.Success("You have deleted store")
	flash.Store(&c.Controller)
	c.Redirect("/admin/stores", 303)
}

func (c *StoresController) getStore() (models.Store, error) {

	store := models.Store{
		Address:     c.GetString("address"),
		Name:        c.GetString("name"),
		PhoneNumber: c.GetString("phone"),
	}

	latitude, err := c.GetFloat("latitude")
	if err != nil {
		return store, err
	}
	store.Latitude = latitude

	longitude, err := c.GetFloat("longitude")
	if err != nil {
		return store, err
	}
	store.Longitude = longitude

	companyID, err := c.GetInt("company")
	if err != nil {
		return store, err
	}
	store.CompanyID = uint(companyID)

	// get services
	str := c.GetString("services")
	if str != "" {
		services := strings.Split(str, ",")
		for _, serviceName := range services {
			s := models.FindServiceByName(serviceName)
			store.Services = append(store.Services, s)
		}
	}

	// get managers
	str = c.GetString("managers")
	if str != "" {
		managers := strings.Split(str, ",")
		for _, managerID := range managers {
			i, err := strconv.Atoi(managerID)
			if err != nil {
				return store, err
			}
			m := models.FindUserByID(uint(i))
			store.Managers = append(store.Managers, m)
		}
	}

	// get image
	defaultImage := c.GetString("oldImage")
	path, err := c.UploadFile("image", "image", defaultImage)
	if err != nil {
		return store, err
	} else {
		store.Image = path
	}



	return store, nil
}
