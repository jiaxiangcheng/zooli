package routers

import (
	"github.com/Qiaorui/zooli/controllers"
	auth "github.com/Qiaorui/zooli/controllers/auth"
	users "github.com/Qiaorui/zooli/controllers/users"
	companies "github.com/Qiaorui/zooli/controllers/companies"
	stores "github.com/Qiaorui/zooli/controllers/stores"
	services "github.com/Qiaorui/zooli/controllers/services"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("", &controllers.MainController{})
	beego.Router("/dashboard", &controllers.MainController{})
	beego.Router("/login", &auth.LoginController{}, "get:LoginForm")
	beego.Router("/login", &auth.LoginController{}, "post:Login")
	beego.Router("/logout", &auth.LoginController{}, "get:Logout")

	beego.Router("/users", &users.UsersController{}, "get:Get")
	beego.Router("/users/:id([0-9]+", &users.UsersController{}, "get:Edit")
	beego.Router("/users/:id([0-9]+", &users.UsersController{}, "post:Update")
	beego.Router("/users/:id([0-9]+", &users.UsersController{}, "delete:Delete")
	beego.Router("/users/new", &users.UsersController{}, "get:New")
	beego.Router("/users/new", &users.UsersController{}, "post:Create")

	beego.Router("/companies", &companies.CompaniesController{}, "get:Get")
	beego.Router("/companies/new", &companies.CompaniesController{}, "get:New")
	beego.Router("/companies/new", &companies.CompaniesController{}, "post:Create")
	beego.Router("/companies/:id([0-9]+", &companies.CompaniesController{}, "get:Edit")
	beego.Router("/companies/:id([0-9]+", &companies.CompaniesController{}, "post:Update")
	beego.Router("/companies/:id([0-9]+", &companies.CompaniesController{}, "delete:Delete")

	beego.Router("/stores", &stores.StoresController{}, "get:Get")

	beego.Router("/services", &services.ServicesController{}, "get:Get")
	beego.Router("/services/new", &services.ServicesController{}, "get:New")
	beego.Router("/services/new", &services.ServicesController{}, "post:Create")
	beego.Router("/services/:id([0-9]+", &services.ServicesController{}, "get:Edit")
	beego.Router("/services/:id([0-9]+", &services.ServicesController{}, "post:Update")
	beego.Router("/services/:id([0-9]+", &services.ServicesController{}, "delete:Delete")

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
