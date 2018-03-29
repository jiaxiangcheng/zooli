package rbac

import (
	"strconv"

	"github.com/Qiaorui/zooli/controllers"
	utils "github.com/Qiaorui/zooli/controllers/utils"
	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
	"github.com/pkg/errors"
	"github.com/beego/i18n"
)

type UsersController struct {
	controllers.BaseController
}

func (c *UsersController) Get() {
	beego.Info("sadsad")
	c.Data["users"] = models.FindUsers()
	c.TplName = "admin/users/list.tpl"
}

func (c *UsersController) Edit() {

	userSession := c.GetSession("userInfo")

	var u models.User

	if userSession != nil {
		c.DelSession("userInfo")
		u = userSession.(models.User)
	} else {
		id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
		// load user in DB
		u = models.FindUserByID(uint(id))
	}
	if !u.Exists() {
		flash := beego.NewFlash()
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_USER_AREADY_EXISTS))
		flash.Store(&c.Controller)
		c.Redirect("/admin/users", 302)
		return
	}

	c.Data["userForm"] = u
	c.Data["roles"] = models.FindRoles()
	c.Data["stores"] = models.FindStores()
	c.TplName = "admin/users/edit.tpl"
}

func (c *UsersController) New() {

	//get the user session and load if exist
	u := c.GetSession("userInfo")
	if u != nil {
		c.DelSession("userInfo")
		c.Data["userForm"] = u.(models.User)
	}

	c.Data["stores"] = models.FindStores()
	c.Data["roles"] = models.FindRoles()
	c.TplName = "admin/users/new.tpl"
}

func (c *UsersController) Create() {
	flash := beego.NewFlash()

	u, err := c.getUser()
	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("userInfo", u)
		c.Redirect("/admin/users/new", 303)
		return
	}

	err = utils.Validate(u)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("userInfo", u)
		c.Redirect("/admin/users/new", 303)
		return
	}

	u.Insert()

	// load message success and redirect
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_CREATE_USER) + " " + u.Name)
	flash.Store(&c.Controller)
	c.Redirect("/admin/users", 303)
}

func (c *UsersController) Update() {
	//init object for error control
	flash := beego.NewFlash()

	//get identifier of user
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	u, err := c.getUser()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/admin/users/"+strconv.Itoa(id), 302)
		return
	}

	u.ID = uint(id)
	if !u.Exists() {
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_USER_AREADY_EXISTS))
		flash.Store(&c.Controller)
		c.Redirect("/admin/users", 302)
		return
	}

	err = utils.Validate(u)

	//load the error, save the form fields and redirect
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.SetSession("userInfo", u)
		c.Redirect("/admin/users/"+strconv.Itoa(id), 302)
		return
	}

	//update the user
	u.Update()

	if u.StoreID != uint(0) {
		store := models.FindStoreByID(u.StoreID)
		if !store.Exists() {
			flash.Error("Store selected not existed")
			flash.Store(&c.Controller)
			c.Redirect("/admin/users", 302)
			return
		}
		u.AssignStore(store)
	}

	// load message success and redirect
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_UPDATE_USER) + " " + u.Name)
	flash.Store(&c.Controller)
	c.Redirect("/admin/users/"+strconv.Itoa(id), 302)
}

func (c *UsersController) Delete() {
	flash := beego.NewFlash()

	//get identifier of user
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	var u models.User
	u.ID = uint(id)
	if !u.Exists() {
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_USER_AREADY_EXISTS))
		flash.Store(&c.Controller)
		c.Redirect("/admin/users", 303)
		return
	}

	u.DeleteSoft()

	// load message success and redirect
	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_DELETE_USER) + " " + u.Name)
	flash.Store(&c.Controller)
	c.Redirect("/admin/users", 303)
}

func (c *UsersController) AssignStore() {
	flash := beego.NewFlash()
	//get identifier of user
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	u := models.FindUserByID(uint(id))
	if !u.Exists() {
		flash.Error(i18n.Tr(c.Lang, utils.ERROR_USER_AREADY_EXISTS))
		flash.Store(&c.Controller)
		c.Redirect("/admin/stores", 303)
		return
	}

	storeID, err := c.GetInt("storeID")
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/admin/stores", 303)
		return
	}

	s := models.FindStoreByID(uint(storeID))
	if !s.Exists() {
		flash.Error("Incorrect store id")
		flash.Store(&c.Controller)
		c.Redirect("/admin/stores", 303)
		return
	}

	u.AssignStore(s)

	flash.Success(i18n.Tr(c.Lang, utils.SUCCESS_ASSIGN_USER) + " " + u.Name +
	 " " + i18n.Tr(c.Lang, utils.TO_STORE) + " " + s.Name)
	flash.Store(&c.Controller)
	c.Redirect("/admin/stores", 303)
}

func (c *UsersController) getUser() (models.User, error) {
	u := models.User{
		Username: c.GetString("username"),
		Email:    c.GetString("email"),
		Name:     c.GetString("name"),
	}
	pwd := c.GetString("password")
	if pwd == "" {
		return u, errors.New("Password can not be empty")
	}
	u.SetPassword(pwd)
	roleID, err := c.GetInt("role")
	if err != nil {
		return u, err
	}
	u.RoleID = uint(roleID)

	beego.Info(c.GetString("stores"))
	if c.GetString("stores") != "" {
		storeID, err := strconv.Atoi(c.GetString("stores"))
		beego.Info(storeID)
		if err != nil {
			return u, err
		}
		u.StoreID = uint(storeID)
	}

	return u, nil
}
