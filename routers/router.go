package routers

import (
	"github.com/Qiaorui/zooli/controllers"
	"github.com/Qiaorui/zooli/controllers/admin"
	"github.com/Qiaorui/zooli/controllers/auth"
	"github.com/Qiaorui/zooli/controllers/public"
	"github.com/Qiaorui/zooli/controllers/rbac"
	"github.com/Qiaorui/zooli/models"
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

	beego.Router("admin/companies", &admin.CompaniesController{}, "get:Get")
	beego.Router("admin/companies/new", &admin.CompaniesController{}, "get:New")
	beego.Router("admin/companies/new", &admin.CompaniesController{}, "post:Create")
	beego.Router("admin/companies/:id([0-9]+", &admin.CompaniesController{}, "get:Edit")
	beego.Router("admin/companies/:id([0-9]+", &admin.CompaniesController{}, "post:Update")
	beego.Router("admin/companies/:id([0-9]+", &admin.CompaniesController{}, "delete:Delete")

	beego.Router("admin/stores", &admin.StoresController{}, "get:Get")
	beego.Router("admin/stores/new", &admin.StoresController{}, "get:New")
	beego.Router("admin/stores/new", &admin.StoresController{}, "post:Create")
	beego.Router("admin/stores/:id([0-9]+", &admin.StoresController{}, "get:Edit")
	beego.Router("admin/stores/:id([0-9]+", &admin.StoresController{}, "post:Update")
	beego.Router("admin/stores/:id([0-9]+", &admin.StoresController{}, "delete:Delete")

	beego.Router("public/store", &public.ManagersStoreController{}, "get:Edit")
	beego.Router("public/store", &public.ManagersStoreController{}, "post:Update")

	beego.Router("admin/services", &admin.ServicesController{}, "get:Get")
	beego.Router("admin/services/new", &admin.ServicesController{}, "get:New")
	beego.Router("admin/services/new", &admin.ServicesController{}, "post:Create")
	beego.Router("admin/services/:id([0-9]+", &admin.ServicesController{}, "get:Edit")
	beego.Router("admin/services/:id([0-9]+", &admin.ServicesController{}, "post:Update")
	beego.Router("admin/services/:id([0-9]+", &admin.ServicesController{}, "delete:Delete")

	beego.Router("public/orders", &public.OrdersController{}, "get:Get")
	beego.Router("public/orders/new", &public.OrdersController{}, "get:New")
	beego.Router("public/orders/new", &public.OrdersController{}, "post:Create")
	beego.Router("public/orders/:id([0-9]+", &public.OrdersController{}, "get:Edit")
	beego.Router("public/orders/:id([0-9]+", &public.OrdersController{}, "post:Update")
	beego.Router("public/orders/:id([0-9]+", &public.OrdersController{}, "delete:Delete")

	beego.Router("public/products", &public.ProductsController{}, "get:Get")
	beego.Router("public/products/new", &public.ProductsController{}, "get:New")
	beego.Router("public/products/new", &public.ProductsController{}, "post:Create")
	beego.Router("public/products/:id([0-9]+", &public.ProductsController{}, "get:Edit")
	beego.Router("public/products/:id([0-9]+", &public.ProductsController{}, "post:Update")
	beego.Router("public/products/:id([0-9]+", &public.ProductsController{}, "delete:Delete")

	beego.Router("/dont_use-long_name-and_bar", &controllers.MainController{}, "get:RandomData")

	beego.InsertFilter("*", beego.BeforeRouter, LoggedInFilter)
	beego.InsertFilter("admin/*", beego.BeforeRouter, AdminFilter)
	beego.InsertFilter("public/*", beego.BeforeRouter, ManagerFilter)

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

var AdminFilter = func(ctx *context.Context) {
	user := ctx.Input.Session("user").(models.User)

	if user.Role.Name != models.ROLE_ADMIN {
		ctx.Redirect(302, "/dashboard")
	}
}

var ManagerFilter = func(ctx *context.Context) {
	user := ctx.Input.Session("user").(models.User)

	if user.Role.Name != models.ROLE_MANAGER {
		ctx.Redirect(302, "/dashboard")
	}
}
