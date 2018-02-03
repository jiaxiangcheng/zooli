package models

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {

	DB = CreateGolibDB()

	profileAdmin := &User{
		Username: "admin",
	}

	if profileAdmin.Exists() == false {
		profileAdmin.SetPassword("1234")

		DB.Save(profileAdmin)
	}

}

func CreateGolibDB() *gorm.DB {
	var mysqlConnect bytes.Buffer

	mysqluser := beego.AppConfig.String("mysqluser")
	fmt.Println("mysqluser:" + mysqluser)
	mysqlConnect.WriteString(mysqluser)

	mysqlConnect.WriteString(":")

	var mysqlpass string
	mysqlpass = os.Getenv("MYSQL_PASS")
	if mysqlpass == "" {
		mysqlpass = beego.AppConfig.String("mysqlpass")
	}
	mysqlConnect.WriteString(mysqlpass)

	mysqlConnect.WriteString("@tcp(")

	mysqlhost := os.Getenv("MYSQL_HOST")
	if mysqlhost == "" {
		mysqlhost = beego.AppConfig.String("mysqlhost")
	}
	fmt.Println("mysqlhost:" + mysqlhost)

	mysqlConnect.WriteString(mysqlhost)

	mysqlport := beego.AppConfig.String("mysqlport")
	if mysqlport == "" {
		mysqlport = "3306"
	}
	fmt.Println("mysqlport:" + mysqlport)
	mysqlConnect.WriteString(":" + mysqlport + ")/")

	mysqldb := beego.AppConfig.String("mysqldb")
	fmt.Println("mysqldb:" + mysqldb)

	mysqlConnect.WriteString(mysqldb)
	mysqlConnect.WriteString("?charset=utf8&parseTime=True&loc=Local")

	db, err := gorm.Open("mysql", mysqlConnect.String())

	if err != nil {
		fmt.Println("Failed to connect database " + err.Error())

		fmt.Println("Trying to create a database: " + mysqldb)

		if mErr, ok := err.(*mysql.MySQLError); ok && mErr.Number == 1049 {
			create(mysqluser, mysqlpass, mysqlhost, mysqlport, mysqldb)
		}
	}

	LoadlibDB(db)

	return db
}

func LoadlibDB(db *gorm.DB) {

	DB = db

	DB.AutoMigrate(&User{})
}

func create(mysqluser, mysqlpass, mysqlhost, mysqlport, mysqldb string) *gorm.DB {
	var err error

	var mysqlConnect []string

	mysqlConnect = append(mysqlConnect, mysqluser, ":", mysqlpass)
	mysqlConnect = append(mysqlConnect, "@tcp(", mysqlhost, ":", mysqlport, ")/?charset=utf8&parseTime=True&loc=Local")

	db, err := gorm.Open("mysql", strings.Join(mysqlConnect, ""))

	if err != nil {
		panic(err)
	}

	fmt.Println("CREATE DATABASE " + mysqldb)
	db.Exec("CREATE DATABASE " + mysqldb)

	fmt.Println("USE " + mysqldb)
	db.Exec("USE " + mysqldb)

	return db

}
