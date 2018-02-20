package controllers

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Qiaorui/zooli/models"
	"github.com/astaxie/beego"
)

const EXTERNAL_FILE_STORAGE = "static/storage/"

type BaseController struct {
	// Embedding: "Inherit" beego.Controller
	beego.Controller
}

func (c *BaseController) Prepare() {
	//// Overwrite beego.Controller.Layout (string)
	_ = beego.ReadFromRequest(&c.Controller)

	u := c.GetSession("user")
	c.Data["user"] = u
	c.Data["roleManager"] = models.ROLE_MANAGER
	c.Data["roleAdmin"] = models.ROLE_ADMIN

	if !c.IsAjax() {
		c.Layout = "layoutAuth.tpl"
	}
}

func (c *BaseController) GetString(key string, def ...string) string {
	str := c.Controller.GetString(key, def...)
	return strings.TrimSpace(str)
}

/*
Using form key and specify format to upload a file.
Return path and error if exists.
*/
func (c *BaseController) UploadFile(key string, format string) (string, error) {
	path := ""
	f, header, err := c.GetFile(key)

	defer f.Close()
	buff := make([]byte, 512) // docs tell that it take only first 512 bytes into consideration
	if _, err = f.Read(buff); err != nil {
		return "", err
	}
	if strings.Contains(http.DetectContentType(buff), format) {
		// if dir not exists, mkdir
		if _, err := os.Stat(EXTERNAL_FILE_STORAGE); os.IsNotExist(err) {
			os.Mkdir(EXTERNAL_FILE_STORAGE, os.ModePerm)
		}

		// get the filename
		fileName := header.Filename
		path = EXTERNAL_FILE_STORAGE + strconv.FormatInt(time.Now().Unix(), 10) + fileName
		// save to server
		if err := c.SaveToFile(key, path); err != nil {
			return "", err
		}
	} else {
		return "", errors.New("the upload file has incorrect format")
	}
	return path, nil
}
