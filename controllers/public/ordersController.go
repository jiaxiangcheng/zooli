package public

import (
	"strconv"

	"github.com/Qiaorui/zooli/controllers"
	utils "github.com/Qiaorui/zooli/controllers/utils"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
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
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_ORDER_AREADY_EXISTS))
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
	c.Data["orderLogs"] = models.FindOrderLogsByOrderID(order.ID)

	status := c.getStatusFromConst(order.Status)
	if status != 5 && status != 6 {
		c.Data["nextStep"] = c.getStatusFromInt(status + 1)
	}
	beego.Debug(c.Data["nextStep"])

	c.TplName = "public/orders/edit.tpl"
}

/*
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

	client := models.Customer{
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
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_CREATE_ORDER))
	flash.Store(&c.Controller)
	c.Redirect("/public/orders", 303)
}
*/

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
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_ORDER_AREADY_EXISTS))
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
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_UPDATE_ORDER))
	flash.Store(&c.Controller)
	c.Redirect("/public/orders/"+strconv.Itoa(id), 302)
}

func (c *OrdersController) NextStatus() {
	flash := beego.NewFlash()

	var order models.Order
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	order = models.FindOrderByID(uint(id))

	oldStatusInt := c.getStatusFromConst(order.Status)
	currenStatusInt := oldStatusInt + 1

	order.Status = c.getStatusFromInt(currenStatusInt)

	order.Update()

	// load message success and redirect
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_NEXTSTATUS_ORDER))
	flash.Store(&c.Controller)
	c.Redirect("/public/orders/"+strconv.Itoa(id), 302)
}

func (c *OrdersController) CancelOrder() {
	flash := beego.NewFlash()

	var order models.Order
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	order = models.FindOrderByID(uint(id))
	order.Status = models.CANCELED

	order.Update()

	// load message success and redirect
	flash.Success("Order has been canceled")
	flash.Store(&c.Controller)
	c.Redirect("/public/orders/"+strconv.Itoa(id), 302)
}

/*
func (c *OrdersController) Delete() {
	flash := beego.NewFlash()

	//get identifier of order
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	var order models.Order
	order.ID = uint(id)
	if !order.Exists() {
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_ORDER_AREADY_EXISTS))
		flash.Store(&c.Controller)
		c.Redirect("/public/orders", 303)
		return
	}

	order.DeleteSoft()

	// load message success and redirect
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_DELETE_ORDER))
	flash.Store(&c.Controller)
	c.Redirect("/public/orders", 303)
}
*/

func (c *OrdersController) getOrder() (models.Order, error) {
	var order models.Order

	fee, err := c.GetFloat("fee")
	if err != nil {
		return order, err
	}
	order.Fee = fee

	/*
		productID, err := c.GetInt("product")
		if err != nil {
			return order, err
		}
		beego.Debug(productID)
		order.ProductID = uint(productID)
		order.Product = models.FindProductByID(order.ProductID)
	*/

	status := c.GetString("status")

	s, _ := strconv.Atoi(status)

	order.Status = c.getStatusFromInt(s)

	return order, nil
}

func (c *OrdersController) getStatusFromConst(status models.Status) int {

	switch status {
	case models.ORDERED:
		return 1
	case models.IN_SERVICE:
		return 2
	case models.END_SERVICE:
		return 3
	case models.WAITING_FOR_PAYMENT:
		return 4
	case models.FINISHED:
		return 5
	case models.CANCELED:
		return 6
	}

	return 0
}

func (c *OrdersController) getStatusFromInt(status int) models.Status {
	switch status {
	case 1:
		return models.ORDERED
	case 2:
		return models.IN_SERVICE
	case 3:
		return models.END_SERVICE
	case 4:
		return models.WAITING_FOR_PAYMENT
	case 5:
		return models.FINISHED
	case 6:
		return models.CANCELED
	}

	return models.CANCELED
}
