package routers

import (
	"github.com/Qiaorui/zooli/controllers"
	adminBusiness "github.com/Qiaorui/zooli/controllers/admin/business"
	"github.com/Qiaorui/zooli/controllers/admin/rbac"
	"github.com/Qiaorui/zooli/controllers/auth"
	publicBusiness "github.com/Qiaorui/zooli/controllers/public/business"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("", &controllers.MainController{})
	beego.Router("/dashboard", &controllers.MainController{})
	beego.Router("/login", &auth.LoginController{}, "get:LoginForm")
	beego.Router("/login", &auth.LoginController{}, "post:Login")
	beego.Router("/logout", &auth.LoginController{}, "get:Logout")

	beego.Router("admin/users", &rbac.UsersController{}, "get:Get")
	beego.Router("admin/users/:id([0-9]+", &rbac.UsersController{}, "get:Edit")
	beego.Router("admin/users/:id([0-9]+", &rbac.UsersController{}, "post:Update")
	beego.Router("admin/users/:id([0-9]+", &rbac.UsersController{}, "delete:Delete")
	beego.Router("admin/users/new", &rbac.UsersController{}, "get:New")
	beego.Router("admin/users/new", &rbac.UsersController{}, "post:Create")
	beego.Router("admin/users/:id([0-9]+/assign", &rbac.UsersController{}, "post:AssignStore")

	beego.Router("admin/companies", &adminBusiness.CompaniesController{}, "get:Get")
	beego.Router("admin/companies/new", &adminBusiness.CompaniesController{}, "get:New")
	beego.Router("admin/companies/new", &adminBusiness.CompaniesController{}, "post:Create")
	beego.Router("admin/companies/:id([0-9]+", &adminBusiness.CompaniesController{}, "get:Edit")
	beego.Router("admin/companies/:id([0-9]+", &adminBusiness.CompaniesController{}, "post:Update")
	beego.Router("admin/companies/:id([0-9]+", &adminBusiness.CompaniesController{}, "delete:Delete")

	beego.Router("admin/stores", &adminBusiness.StoresController{}, "get:Get")
	beego.Router("admin/stores/new", &adminBusiness.StoresController{}, "get:New")
	beego.Router("admin/stores/new", &adminBusiness.StoresController{}, "post:Create")
	beego.Router("admin/stores/:id([0-9]+", &adminBusiness.StoresController{}, "get:Edit")
	beego.Router("admin/stores/:id([0-9]+", &adminBusiness.StoresController{}, "post:Update")
	beego.Router("admin/stores/:id([0-9]+", &adminBusiness.StoresController{}, "delete:Delete")

	beego.Router("public/store", &publicBusiness.ManagersStoreController{}, "get:Edit")
	beego.Router("public/store", &publicBusiness.ManagersStoreController{}, "post:Update")

	beego.Router("admin/services", &adminBusiness.ServicesController{}, "get:Get")
	beego.Router("admin/services/new", &adminBusiness.ServicesController{}, "get:New")
	beego.Router("admin/services/new", &adminBusiness.ServicesController{}, "post:Create")
	beego.Router("admin/services/:id([0-9]+", &adminBusiness.ServicesController{}, "get:Edit")
	beego.Router("admin/services/:id([0-9]+", &adminBusiness.ServicesController{}, "post:Update")
	beego.Router("admin/services/:id([0-9]+", &adminBusiness.ServicesController{}, "delete:Delete")

	beego.Router("public/orders", &publicBusiness.OrdersController{}, "get:Get")

	beego.Router("public/products", &publicBusiness.ProductsController{}, "get:Get")
	beego.Router("public/products/new", &publicBusiness.ProductsController{}, "get:New")
	beego.Router("public/products/new", &publicBusiness.ProductsController{}, "post:Create")
	beego.Router("public/products/:id([0-9]+", &publicBusiness.ProductsController{}, "get:Edit")
	beego.Router("public/products/:id([0-9]+", &publicBusiness.ProductsController{}, "post:Update")
	beego.Router("public/products/:id([0-9]+", &publicBusiness.ProductsController{}, "delete:Delete")


	beego.Router("/dont_use-long_name-and_bar", &controllers.MainController{}, "get:RandomData")

	beego.InsertFilter("*", beego.BeforeRouter, LoggedInFilter)
}

var LoggedInFilter = func(ctx *context.Context) {
	user := ctx.Input.Session("user")

	//TODO: if user is admin, then has permission admin/*

	if user == nil && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
	if user != nil && ctx.Request.RequestURI == "/login" {
		ctx.Redirect(302, "/dashboard")
	}
}
