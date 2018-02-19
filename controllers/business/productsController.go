package business

import (
    "github.com/Qiaorui/zooli/controllers"
    "github.com/Qiaorui/zooli/models"
)

type ProductsController struct {
	controllers.BaseController
}

func (c *ProductsController) Get() {
	c.Data["products"] = models.FindProducts()
	c.TplName = "products/list.tpl"
}
/*
func (c *ProductsController) Edit() {

	productSession := c.GetSession("productsInfo")

	var product models.Order

	if productSession != nil {
		c.DelSession("productInfo")
		product = productSession.(models.Order)
	} else {
		id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
		// load product in DB
		product = models.FindOrderByID(uint(id))
	}
	if !product.Exists() {
		flash := beego.NewFlash()
		flash.Error("Incorrect product id")
		flash.Store(&c.Controller)
		c.Redirect("/products", 302)
		return
	}

	c.Data["productForm"] = product
	c.TplName = "products/edit.tpl"
}

func (c *ProductsController) New() {

	//get the product session and load if exist
	product := c.GetSession("productInfo")
	if product != nil {
		c.DelSession("productInfo")
		c.Data["productForm"] = product.(models.Order)
	}

	c.TplName = "products/new.tpl"
}

func (c *ProductsController) Create() {
	flash := beego.NewFlash()

	product, err := c.getOrder()
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

	product, err := c.getOrder()
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

	var product models.Order
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
}*/

/*
func (c *ProductsController) getOrder() (models.Order, error) {
	product := models.Order{
		Name: c.GetString("name"),
	}
	return product, nil
}*/
