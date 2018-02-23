package routers

import (
	"github.com/Qiaorui/zooli/controllers"
	"github.com/Qiaorui/zooli/controllers/auth"

	"github.com/Qiaorui/zooli/controllers/business"
	"github.com/Qiaorui/zooli/controllers/rbac"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("", &controllers.MainController{})
	beego.Router("/dashboard", &controllers.MainController{})
	beego.Router("/login", &auth.LoginController{}, "get:LoginForm")
	beego.Router("/login", &auth.LoginController{}, "post:Login")
	beego.Router("/logout", &auth.LoginController{}, "get:Logout")

	beego.Router("/users", &rbac.UsersController{}, "get:Get")
	beego.Router("/users/:id([0-9]+", &rbac.UsersController{}, "get:Edit")
	beego.Router("/users/:id([0-9]+", &rbac.UsersController{}, "post:Update")
	beego.Router("/users/:id([0-9]+", &rbac.UsersController{}, "delete:Delete")
	beego.Router("/users/new", &rbac.UsersController{}, "get:New")
	beego.Router("/users/new", &rbac.UsersController{}, "post:Create")
	beego.Router("/users/:id([0-9]+/assign", &rbac.UsersController{}, "post:AssignStore")

	beego.Router("/companies", &business.CompaniesController{}, "get:Get")
	beego.Router("/companies/new", &business.CompaniesController{}, "get:New")
	beego.Router("/companies/new", &business.CompaniesController{}, "post:Create")
	beego.Router("/companies/:id([0-9]+", &business.CompaniesController{}, "get:Edit")
	beego.Router("/companies/:id([0-9]+", &business.CompaniesController{}, "post:Update")
	beego.Router("/companies/:id([0-9]+", &business.CompaniesController{}, "delete:Delete")

	beego.Router("/stores", &business.StoresController{}, "get:Get")
	beego.Router("/stores/new", &business.StoresController{}, "get:New")
	beego.Router("/stores/new", &business.StoresController{}, "post:Create")
	beego.Router("/stores/:id([0-9]+", &business.StoresController{}, "get:Edit")
	beego.Router("/stores/:id([0-9]+", &business.StoresController{}, "post:Update")
	beego.Router("/stores/:id([0-9]+", &business.StoresController{}, "delete:Delete")

	beego.Router("/manager-store", &business.ManagersStoreController{}, "get:Edit")
	beego.Router("/manager-store", &business.ManagersStoreController{}, "post:Update")

	beego.Router("/services", &business.ServicesController{}, "get:Get")
	beego.Router("/services/new", &business.ServicesController{}, "get:New")
	beego.Router("/services/new", &business.ServicesController{}, "post:Create")
	beego.Router("/services/:id([0-9]+", &business.ServicesController{}, "get:Edit")
	beego.Router("/services/:id([0-9]+", &business.ServicesController{}, "post:Update")
	beego.Router("/services/:id([0-9]+", &business.ServicesController{}, "delete:Delete")

	beego.Router("/orders", &business.OrdersController{}, "get:Get")

	beego.Router("/products", &business.ProductsController{}, "get:Get")
	beego.Router("/products/new", &business.ProductsController{}, "get:New")
	beego.Router("/products/new", &business.ProductsController{}, "post:Create")
	beego.Router("/products/:id([0-9]+", &business.ProductsController{}, "get:Edit")
	beego.Router("/products/:id([0-9]+", &business.ProductsController{}, "post:Update")
	beego.Router("/products/:id([0-9]+", &business.ProductsController{}, "delete:Delete")

	beego.Router("/dont_use-long_name-and_bar", &controllers.MainController{}, "get:RandomData")

	beego.InsertFilter("*", beego.BeforeRouter, LoggedInFilter)
}

var LoggedInFilter = func(ctx *context.Context) {
	user := ctx.Input.Session("user")

	if user == nil && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
	if user != nil && ctx.Request.RequestURI == "/login" {
		ctx.Redirect(302, "/dashboard")
	}
}
