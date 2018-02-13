package routers

import (
	"github.com/Qiaorui/zooli/controllers"
	auth "github.com/Qiaorui/zooli/controllers/auth"
	users "github.com/Qiaorui/zooli/controllers/users"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("", &controllers.MainController{})
	beego.Router("/dashboard", &controllers.MainController{})
	beego.Router("/login", &auth.LoginController{}, "get:LoginForm")
	beego.Router("/login", &auth.LoginController{}, "post:Login")
	beego.Router("/logout", &auth.LoginController{}, "get:Logout")

	beego.Router("/users", &users.UsersController{})
	beego.Router("/users/:id", &users.UsersController{}, "get:Edit")
	beego.Router("/users/new", &users.UsersController{}, "get:New")
	beego.Router("/users/new", &users.UsersController{}, "post:Create")
	//beego.Router("/users/existUserIf", &users.UsersController{}, "post:ExistUserIf")
	beego.Router("/users/:id", &users.UsersController{}, "post:Update")
	beego.Router("/users/:id", &users.UsersController{}, "delete:Delete")

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
