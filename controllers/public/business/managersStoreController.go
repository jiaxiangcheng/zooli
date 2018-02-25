package business

import (
	"github.com/Qiaorui/zooli/controllers"
	utils "github.com/Qiaorui/zooli/controllers/utils"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
)

type ManagersStoreController struct {
	controllers.BaseController
}

func (c *ManagersStoreController) Edit() {
	store, _ := GetCurrentStore(&c.BaseController)

	c.Data["storeForm"] = store
	c.TplName = "public/store/edit.tpl"
}

func (c *ManagersStoreController) Update() {
	flash := beego.NewFlash()
	storeUpdated, err := c.getStore()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/dashboard/", 302)
		return
	}

	if !storeUpdated.Exists() {
		flash.Error("Incorrect store id")
		flash.Store(&c.Controller)
		c.Redirect("/dashboard", 302)
		return
	}

	err = utils.Validate(storeUpdated)

	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/dashboard", 302)
		return
	}

	//update the store
	storeUpdated.Update()

	// load message success and redirect
	flash.Success("You have update the store")
	flash.Store(&c.Controller)
	c.Redirect("/public/store", 302)
}

/*
func (c *ManagersStoreController) GetStoreFromDB() models.Store {
	manager := c.GetSession("user")
	id := manager.(models.User).StoreID
	var store models.Store
	store = models.FindStoreByID(uint(id))

	return store
}
*/

func (c *ManagersStoreController) getStore() (models.Store, error) {

	// var store models.Store
	storeDb, _ := GetCurrentStore(&c.BaseController)
	store := models.Store{
		Company:     storeDb.Company,
		CompanyID:   storeDb.CompanyID,
		Managers:    storeDb.Managers,
		Services:    storeDb.Services,
		Address:     c.GetString("address"),
		Name:        c.GetString("name"),
		PhoneNumber: c.GetString("phone"),
	}

	store.ID = storeDb.ID

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
