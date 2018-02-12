package routers

import (
	"github.com/Qiaorui/zooli/controllers"
	auth "github.com/Qiaorui/zooli/controllers/auth"
	users "github.com/Qiaorui/zooli/controllers/users"
	dashboard "github.com/Qiaorui/zooli/controllers/dashboard"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/dashboard", &controllers.MainController{})
	beego.Router("/login", &auth.LoginController{}, "get:LoginForm")
	beego.Router("/login", &auth.LoginController{}, "post:Login")

	beego.Router("/dashboard", &dashboard.DashController{}, "get:LoadDashboard")
	beego.Router("/users", &users.UsersController{}, "get:Index")
	beego.Router("/users/:id", &users.UsersController{}, "post:LoadUser")
	beego.Router("/users/new", &users.UsersController{}, "post:CreateUser")
	beego.Router("/users/insert", &users.UsersController{}, "post:InsertUser")
	beego.Router("/users/existUserIf", &users.UsersController{}, "post:ExistUserIf")

	beego.InsertFilter("/*", beego.BeforeRouter, LoggedInFilter)
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