package business

import (
    "github.com/Qiaorui/zooli/controllers"
    "github.com/Qiaorui/zooli/models"
)

type OrdersController struct {
	controllers.BaseController
}

func (c *OrdersController) Get() {
	c.Data["orders"] = models.FindOrders()
	c.TplName = "orders/list.tpl"
}
/*
func (c *OrdersController) Edit() {

	orderSession := c.GetSession("ordersInfo")

	var order models.Order

	if orderSession != nil {
		c.DelSession("orderInfo")
		order = orderSession.(models.Order)
	} else {
		id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
		// load order in DB
		order = models.FindOrderByID(uint(id))
	}
	if !order.Exists() {
		flash := beego.NewFlash()
		flash.Error("Incorrect order id")
		flash.Store(&c.Controller)
		c.Redirect("/orders", 302)
		return
	}

	c.Data["orderForm"] = order
	c.TplName = "orders/edit.tpl"
}

func (c *OrdersController) New() {

	//get the order session and load if exist
	order := c.GetSession("orderInfo")
	if order != nil {
		c.DelSession("orderInfo")
		c.Data["orderForm"] = order.(models.Order)
	}

	c.TplName = "orders/new.tpl"
}

func (c *OrdersController) Create() {
	flash := beego.NewFlash()

	order, err := c.getOrder()
	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("orderInfo", order)
		c.Redirect("/orders/new", 303)
		return
	}

	err = utils.Validate(order)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("orderInfo", order)
		c.Redirect("/orders/new", 303)
		return
	}

	order.Insert()

	// load message success and redirect
	flash.Success("You have created the order " + order.Name)
	flash.Store(&c.Controller)
	c.Redirect("/orders", 303)
}

func (c *OrdersController) Update() {
	//init object for error control
	flash := beego.NewFlash()

	//get identifier of order
	id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	order, err := c.getOrder()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/orders/" + strconv.Itoa(id), 302)
		return
	}

	order.ID = uint(id)
	if !order.Exists() {
		flash.Error("Incorrect order id")
		flash.Store(&c.Controller)
		c.Redirect("/orders", 302)
		return
	}

	err = utils.Validate(order)

	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("orderInfo", order)
		c.Redirect("/orders/" + strconv.Itoa(id), 302)
		return
	}

	//update the order
	order.Update()

	// load message success and redirect
	flash.Success("You have updated the order " + order.Name)
	flash.Store(&c.Controller)
	c.Redirect("/orders/" + strconv.Itoa(id), 302)
}


func (c *OrdersController) Delete() {
	flash := beego.NewFlash()

	//get identifier of order
	id , _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	var order models.Order
	order.ID = uint(id)
	if !order.Exists() {
		flash.Error("Incorrect order id")
		flash.Store(&c.Controller)
		c.Redirect("/orders", 303)
		return
	}

	order.DeleteSoft()

	// load message success and redirect
	flash.Success("You have deleted order")
	flash.Store(&c.Controller)
	c.Redirect("/orders", 303)
}*/

/*
func (c *OrdersController) getOrder() (models.Order, error) {
	order := models.Order{
		Name: c.GetString("name"),
	}
	return order, nil
}*/
