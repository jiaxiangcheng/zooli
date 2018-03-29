package public

import (
	"strconv"

	"github.com/Qiaorui/zooli/controllers"
	utils "github.com/Qiaorui/zooli/controllers/utils"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type ProductsController struct {
	controllers.BaseController
}

func (c *ProductsController) Get() {
	s, err := GetCurrentStore(&c.BaseController)
	if err != nil {
		c.Redirect("/dashboard", 302)
	}

	c.Data["products"] = models.FindProductsByStoreID(s.ID)
	c.TplName = "public/products/list.tpl"
}

func (c *ProductsController) Edit() {
	s, err := GetCurrentStore(&c.BaseController)
	if err != nil {
		c.Redirect("/dashboard", 302)
	}

	productSession := c.GetSession("productsInfo")

	var product models.Product

	if productSession != nil {
		c.DelSession("productInfo")
		product = productSession.(models.Product)
	} else {
		id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
		// load product in DB
		product = models.FindProductByID(uint(id))
	}
	if !product.Exists() || product.StoreID != s.ID {
		flash := beego.NewFlash()
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_PRODUCT_AREADY_EXISTS))
		flash.Store(&c.Controller)
		c.Redirect("/public/products", 302)
		return
	}

	c.Data["services"] = s.Services
	c.Data["productForm"] = product
	c.TplName = "public/products/edit.tpl"
}

func (c *ProductsController) New() {
	s, err := GetCurrentStore(&c.BaseController)
	if err != nil {
		c.Redirect("/dashboard", 302)
	}

	//get the product session and load if exist
	product := c.GetSession("productInfo")
	if product != nil {
		c.DelSession("productInfo")
		c.Data["productForm"] = product.(models.Product)
	}

	c.Data["services"] = s.Services
	c.TplName = "public/products/new.tpl"
}

func (c *ProductsController) Create() {
	flash := beego.NewFlash()

	s, err := GetCurrentStore(&c.BaseController)
	if err != nil {
		c.Redirect("/dashboard", 302)
	}

	product, err := c.getProduct()
	product.StoreID = s.ID
	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("productInfo", product)
		c.Redirect("/public/products/new", 303)
		return
	}

	err = utils.Validate(product)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("productInfo", product)
		c.Redirect("/public/products/new", 303)
		return
	}

	product.Insert()

	// load message success and redirect
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_CREATE_PRODUCT) + " " + product.Name)
	flash.Store(&c.Controller)
	c.Redirect("/public/products", 303)
}

func (c *ProductsController) Update() {
	//init object for error control
	flash := beego.NewFlash()

	s, err := GetCurrentStore(&c.BaseController)
	if err != nil {
		c.Redirect("/dashboard", 302)
	}

	//get identifier of product
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	product, err := c.getProduct()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/public/products"+strconv.Itoa(id), 302)
		return
	}

	pDB := models.FindProductByID(uint(id))
	if !pDB.Exists() || pDB.StoreID != s.ID {
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_PRODUCT_AREADY_EXISTS))
		flash.Store(&c.Controller)
		c.Redirect("/public/products", 302)
		return
	}
	product.ID = pDB.ID
	product.StoreID = pDB.StoreID

	err = utils.Validate(product)

	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("productInfo", product)
		c.Redirect("/public/products/"+strconv.Itoa(id), 302)
		return
	}

	//update the product
	product.Update()

	// load message success and redirect
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_UPDATE_PRODUCT) + " " + product.Name)
	flash.Store(&c.Controller)
	c.Redirect("/public/products/"+strconv.Itoa(id), 302)
}

func (c *ProductsController) Delete() {
	flash := beego.NewFlash()

	s, err := GetCurrentStore(&c.BaseController)
	if err != nil {
		c.Redirect("/dashboard", 302)
	}

	//get identifier of product
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	product := models.FindProductByID(uint(id))
	if !product.Exists() || product.StoreID != s.ID {
		flash.Error("Incorrect product id")
		flash.Store(&c.Controller)
		c.Redirect("/public/products", 303)
		return
	}

	product.DeleteSoft()

	// load message success and redirect
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_DELETE_PRODUCT))
	flash.Store(&c.Controller)
	c.Redirect("/public/products", 303)
}

func (c *ProductsController) getProduct() (models.Product, error) {
	product := models.Product{
		Name:        c.GetString("name"),
		Description: c.GetString("description"),
	}

	if c.GetString("value") != "" {
		value, err := c.GetFloat("value")
		if err != nil {
			return product, err
		}
		product.Value = value
	}

	serviceID, err := c.GetInt("service")
	if err != nil {
		return product, err
	}
	product.ServiceID = uint(serviceID)
	product.Service = models.FindServiceByID(product.ServiceID)

	// get image
	/*
		defaultImage := c.GetString("oldImage")
		path, err := c.UploadFileByKey("image", "image", defaultImage)
		if err != nil {
			return product, err
		} else {
			product.Image = path
		}*/

	return product, nil
}
