package business

import (
    "github.com/Qiaorui/zooli/controllers"
    "github.com/Qiaorui/zooli/models"
    "github.com/astaxie/beego"
    "strconv"
    utils "github.com/Qiaorui/zooli/controllers/utils"
)

type ProductsController struct {
	controllers.BaseController
}

func (c *ProductsController) Get() {
	c.Data["products"] = models.FindProducts()
	c.TplName = "products/list.tpl"
}

func (c *ProductsController) Edit() {

	productSession := c.GetSession("productsInfo")

	var product models.Product

	if productSession != nil {
		c.DelSession("productInfo")
		product = productSession.(models.Product)
	} else {
		id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
		// load product in DB
		product = models.FindProductByID(uint(id))
	}
	if !product.Exists() {
		flash := beego.NewFlash()
		flash.Error("Incorrect product id")
		flash.Store(&c.Controller)
		c.Redirect("/products", 302)
		return
	}
    c.Data["services"] = models.FindServices()
	c.Data["productForm"] = product
	c.TplName = "products/edit.tpl"
}

func (c *ProductsController) New() {

	//get the product session and load if exist
	product := c.GetSession("productInfo")
	if product != nil {
		c.DelSession("productInfo")
		c.Data["productForm"] = product.(models.Product)
	}
    c.Data["services"] = models.FindServices()
	c.TplName = "products/new.tpl"
}

func (c *ProductsController) Create() {
	flash := beego.NewFlash()

	product, err := c.getProduct()
	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("productInfo", product)
		c.Redirect("/products/new", 303)
		return
	}

	err = utils.Validate(product)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("productInfo", product)
		c.Redirect("/products/new", 303)
		return
	}

	product.Insert()

	// load message success and redirect
	flash.Success("You have created the product " + product.Name)
	flash.Store(&c.Controller)
	c.Redirect("/products", 303)
}

func (c *ProductsController) Update() {
	//init object for error control
	flash := beego.NewFlash()

	//get identifier of product
	id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	product, err := c.getProduct()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/products/" + strconv.Itoa(id), 302)
		return
	}

	product.ID = uint(id)
	if !product.Exists() {
		flash.Error("Incorrect product id")
		flash.Store(&c.Controller)
		c.Redirect("/products", 302)
		return
	}

	err = utils.Validate(product)

	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("productInfo", product)
		c.Redirect("/products/" + strconv.Itoa(id), 302)
		return
	}

	//update the product
	product.Update()

	// load message success and redirect
	flash.Success("You have updated the product " + product.Name)
	flash.Store(&c.Controller)
	c.Redirect("/products/" + strconv.Itoa(id), 302)
}

func (c *ProductsController) Delete() {
	flash := beego.NewFlash()

	//get identifier of product
	id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	var product models.Product
	product.ID = uint(id)
	if !product.Exists() {
		flash.Error("Incorrect product id")
		flash.Store(&c.Controller)
		c.Redirect("/products", 303)
		return
	}

	product.DeleteSoft()

	// load message success and redirect
	flash.Success("You have deleted product")
	flash.Store(&c.Controller)
	c.Redirect("/products", 303)
}

func (c *ProductsController) getProduct() (models.Product, error) {
	product := models.Product{
		Name: c.GetString("name"),
        Description: c.GetString("description"),

        Image: c.GetString("image"),
        //falta get del service
	}

    value, err := c.GetFloat("value")
	if err != nil {
		return product, err
	}
	product.Value = value
	return product, nil
}
