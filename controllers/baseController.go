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
	"github.com/beego/i18n"
)

const EXTERNAL_FILE_STORAGE = "static/storage/"


var langTypes []*langType // Languages are supported.

// langType represents a language type.
type langType struct {
	Lang, Name string
}

type BaseController struct {
	beego.Controller
	i18n.Locale
}


func init() {
	beego.AddFuncMap("i18n", i18n.Tr)

	// Initialized language type list.
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
	names := strings.Split(beego.AppConfig.String("lang::names"), "|")
	langTypes = make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	// Load locale files according to language types.
	for _, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}

func (c *BaseController) Prepare() {
	//// Overwrite beego.Controller.Layout (string)
	_ = beego.ReadFromRequest(&c.Controller)

	u := c.GetSession("user")
	c.Data["user"] = u
	c.Data["roleManager"] = models.ROLE_MANAGER
	c.Data["roleAdmin"] = models.ROLE_ADMIN

	if c.setLangVer() {
		i := strings.Index(c.Ctx.Request.RequestURI, "?")
		c.Redirect(c.Ctx.Request.RequestURI[:i], 302)
		return
	}

	if !c.IsAjax() {
		c.Layout = "layoutAuth.tpl"
	}
}

// setLangVer sets site language version.
func (c *BaseController) setLangVer() bool {
	isNeedRedir := false
	hasCookie := false

	// 1. Check URL arguments.
	lang := c.Input().Get("lang")

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = c.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := c.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 4. Default language is English.
	if len(lang) == 0 {
		lang = "en-US"
		isNeedRedir = false
	}

	curLang := langType{
		Lang: lang,
	}

	// Save language information in cookies.
	if !hasCookie {
		c.Ctx.SetCookie("lang", curLang.Lang, 1<<31-1, "/")
	}

	restLangs := make([]*langType, 0, len(langTypes)-1)
	for _, v := range langTypes {
		if lang != v.Lang {
			restLangs = append(restLangs, v)
		} else {
			curLang.Name = v.Name
		}
	}

	// Set language properties.
	c.Lang = lang
	c.Data["Lang"] = curLang.Lang
	c.Data["CurLang"] = curLang.Name
	c.Data["RestLangs"] = restLangs

	return isNeedRedir
}


func (c *BaseController) GetString(key string, def ...string) string {
	str := c.Controller.GetString(key, def...)
	return strings.TrimSpace(str)
}

/*
Using form key and specify format to upload a file.
Return path and error if exists.
*/
func (c *BaseController) UploadFile(key string, format string, defaultPath string) (string, error) {
	path := defaultPath
	f, h, err := c.GetFile(key)
	if err != nil {
		return path, nil
	}
	defer f.Close()
	if h.Filename == "" {
		return path, nil
	}
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
		path = EXTERNAL_FILE_STORAGE + strconv.FormatInt(time.Now().Unix(), 10) + h.Filename
		// save to server
		if err := c.SaveToFile(key, path); err != nil {
			return "", err
		}
	} else {
		return "", errors.New("the upload file has incorrect format")
	}
	return c.Ctx.Input.Site() + ":" + strconv.Itoa(c.Ctx.Input.Port()) + "/" + path, nil
}
