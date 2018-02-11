package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	// Embedding: "Inherit" beego.Controller
	beego.Controller
}

func (c *BaseController) Prepare() {

	//// Overwrite beego.Controller.Layout (string)
	u := c.GetSession("user")
	c.Data["user"] = u
}
