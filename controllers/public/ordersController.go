package public

import (
	"strconv"

	"github.com/Qiaorui/zooli/controllers"
	utils "github.com/Qiaorui/zooli/controllers/utils"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
)

type OrdersController struct {
	controllers.BaseController
}


func (c *OrdersController) Get() {
	s, err := GetCurrentStore(&c.BaseController)
	if err != nil {
		c.Redirect("/dashboard", 302)
	}

	c.Data["ordered"] = models.ORDERED
	c.Data["inService"] = models.IN_SERVICE
	c.Data["endService"] = models.END_SERVICE
	c.Data["waitingForPayment"] = models.WAITING_FOR_PAYMENT
	c.Data["orderFinished"] = models.FINISHED
	c.Data["orderedCanceled"] = models.CANCELED
	c.Data["orders"] = models.FindOrdersByStoreID(s.ID)
	c.TplName = "public/orders/list.tpl"
}

func (c *OrdersController) Edit() {

	orderSession := c.GetSession("ordersInfo")

	var order models.Order

	if orderSession != nil {
		c.DelSession("orderInfo")
		order = orderSession.(models.Order)
	} else {
		id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
		// load order in DB
		order = models.FindOrderByID(uint(id))
	}
	if !order.Exists() {
		flash := beego.NewFlash()
		flash.Error("Incorrect order id")
		flash.Store(&c.Controller)
		c.Redirect("/public/orders", 302)
		return
	}

	c.Data["ordered"] = models.ORDERED
	c.Data["inService"] = models.IN_SERVICE
	c.Data["endService"] = models.END_SERVICE
	c.Data["waitingForPayment"] = models.WAITING_FOR_PAYMENT
	c.Data["orderFinished"] = models.FINISHED
	c.Data["orderedCanceled"] = models.CANCELED
	c.Data["orderForm"] = order
	c.TplName = "public/orders/edit.tpl"
}

func (c *OrdersController) New() {

	beego.Debug("new order")
	//get the order session and load if exist
	order := c.GetSession("orderInfo")
	if order != nil {
		c.DelSession("orderInfo")
		initialOrder := order.(models.Order)
		initialOrder.Status = models.ORDERED
		c.Data["orderForm"] = initialOrder
	}

	manager := c.GetSession("user")
	id := manager.(models.User).StoreID

	c.Data["products"] = models.FindProductsByStoreID(uint(id))
	beego.Debug(c.Data["products"])

	c.TplName = "public/orders/new.tpl"
}

func (c *OrdersController) Create() {
	flash := beego.NewFlash()

	beego.Debug("create order")
	// fake client
	clientName := c.GetString("name")

	client := models.Client{
		Name:         clientName,
		PhoneNumber:  "1234",
		PasswordHash: "1234",
		Email:        "1234@fake.com",
	}
	client.Insert()

	order, err := c.getOrder()
	beego.Debug(order)

	//load the error, save the form fields and redirect
	if err != nil {
		beego.Debug(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("orderInfo", order)
		c.Redirect("/public/orders/new", 303)
		return
	}

	beego.Debug("validate order")
	err = utils.Validate(order)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("orderInfo", order)
		c.Redirect("/public/orders/new", 303)
		return
	}

	beego.Debug("go insert order")
	order.Insert()

	// load message success and redirect
	flash.Success("You have created the order")
	flash.Store(&c.Controller)
	c.Redirect("/public/orders", 303)
}

func (c *OrdersController) Update() {
	//init object for error control
	flash := beego.NewFlash()

	//get identifier of order
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	order, err := c.getOrder()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/public/orders/"+strconv.Itoa(id), 302)
		return
	}

	order.ID = uint(id)
	if !order.Exists() {
		flash.Error("Incorrect order id")
		flash.Store(&c.Controller)
		c.Redirect("/public/orders", 302)
		return
	}

	err = utils.Validate(order)

	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("orderInfo", order)
		c.Redirect("/public/orders/"+strconv.Itoa(id), 302)
		return
	}

	//update the order
	order.Update()

	// load message success and redirect
	flash.Success("You have updated the order")
	flash.Store(&c.Controller)
	c.Redirect("/public/orders/"+strconv.Itoa(id), 302)
}

func (c *OrdersController) Delete() {
	flash := beego.NewFlash()

	//get identifier of order
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	var order models.Order
	order.ID = uint(id)
	if !order.Exists() {
		flash.Error("Incorrect order id")
		flash.Store(&c.Controller)
		c.Redirect("/public/orders", 303)
		return
	}

	order.DeleteSoft()

	// load message success and redirect
	flash.Success("You have deleted order")
	flash.Store(&c.Controller)
	c.Redirect("/public/orders", 303)
}

func (c *OrdersController) getOrder() (models.Order, error) {
	var order models.Order

	fee, err := c.GetFloat("fee")
	if err != nil {
		return order, err
	}
	beego.Debug(fee)
	order.Fee = fee

	productID, err := c.GetInt("product")
	if err != nil {
		return order, err
	}
	beego.Debug(productID)
	order.ProductID = uint(productID)
	order.Product = models.FindProductByID(order.ProductID)

	status := c.GetString("status")

	beego.Debug(status)

	switch s, _ := strconv.Atoi(status); s {
	case 1:
		order.Status = models.ORDERED
	case 2:
		order.Status = models.IN_SERVICE
	case 3:
		order.Status = models.END_SERVICE
	case 4:
		order.Status = models.WAITING_FOR_PAYMENT
	case 5:
		order.Status = models.FINISHED
	case 6:
		order.Status = models.CANCELED
	}

	return order, nil
}
