package routers

import (
	"github.com/Qiaorui/zooli/controllers"
	auth "github.com/Qiaorui/zooli/controllers/auth"
	users "github.com/Qiaorui/zooli/controllers/users"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/login", &auth.LoginController{}, "get:LoginForm")
	beego.Router("/login", &auth.LoginController{}, "post:Login")

	beego.Router("/users", &users.UsersController{}, "get:Index")
	beego.Router("/users/{id}", &users.UsersController{}, "post:LoadUser")
	beego.InsertFilter("/*", beego.BeforeRouter, FilterLoggedIn)
}

var FilterLoggedIn = func(ctx *context.Context) {
	user := ctx.Input.Session("user")

	if user == nil && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}
