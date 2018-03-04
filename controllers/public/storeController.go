package public

import (
	"github.com/Qiaorui/zooli/controllers"
	utils "github.com/Qiaorui/zooli/controllers/utils"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/beego/i18n"
)

type ManagersStoreController struct {
	controllers.BaseController
}

func (c *ManagersStoreController) Edit() {
	store, err := GetCurrentStore(&c.BaseController)
	if err != nil {
		c.Redirect("/dashboard", 302)
	}

	storeSession := c.GetSession("storeInfo")
	if storeSession != nil {
		c.DelSession("storeInfo")
		store = storeSession.(models.Store)
	}

	c.Data["storeForm"] = store
	c.TplName = "public/store/edit.tpl"
}

func (c *ManagersStoreController) Update() {
	flash := beego.NewFlash()
	store, err := c.getStore()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("storeInfo", store)
		c.Redirect("/dashboard/", 302)
		return
	}

	if !store.Exists() {
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_STORE_AREADY_EXISTS))
		flash.Store(&c.Controller)
		c.Redirect("/dashboard", 302)
		return
	}

	err = utils.Validate(store)

	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("storeInfo", store)
		c.Redirect("/dashboard", 302)
		return
	}

	//update the store
	store.Update()

	// load message success and redirect
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_UPDATE_STORE))
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


	imageFiles, err := c.GetFiles("files[]")
	if err == nil {
		for _, h := range imageFiles {
			fmt.Println(h.Filename)
			path, err := c.UploadFileByFile(h, "image")
			if err != nil {
				fmt.Println("ERROR: ", err)
				return store, err
			} else {
				store.Images = append(store.Images, models.StoreImage{
					Image: path,
				})
			}
		}
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

	// get image
	/*
		defaultImage := c.GetString("oldImage")
		path, err := c.UploadFileByKey("image", "image", defaultImage)
		if err != nil {
			return store, err
		} else {
			store.Image = path
		}*/

	return store, nil
}
