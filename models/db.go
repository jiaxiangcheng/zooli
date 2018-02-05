package models

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
)

var DB *gorm.DB



func LoadlibDB(db *gorm.DB) {
	db.AutoMigrate(&User{})
}


func Syncdb() {
	createDB()
	if err := Connect(); err != nil {
		beego.Error(err)
		return
	}
	insertUser()
	fmt.Println("database init is complete.")
}

//数据库连接
func Connect() error {
	var dsn string
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	db_path := beego.AppConfig.String("db_path")
	db_sslmode := beego.AppConfig.String("db_sslmode")
	switch db_type {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db_user, db_pass, db_host, db_port, db_name)
		break
	case "postgres":
		dsn = fmt.Sprintf("dbname=%s host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_name, db_host, db_user, db_pass, db_port, db_sslmode)
		break
	case "sqlite3":
		if db_path == "" {
			db_path = "./"
		}
		dsn = fmt.Sprintf("%s%s.db", db_path, db_name)
		break
	case "mssql":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", db_user, db_pass, db_host, db_port, db_name)
		break
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}
	db, err := gorm.Open(db_type, dsn)

	if err == nil {
		DB = db
		LoadlibDB(db)
	} else {
		beego.Warning(err)
	}
	return err
}

//创建数据库
func createDB() {

	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	db_path := beego.AppConfig.String("db_path")
	db_sslmode := beego.AppConfig.String("db_sslmode")

	var dsn string
	var sqlstring string
	switch db_type {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local", db_user, db_pass, db_host, db_port)
		sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", db_name)
		break
	case "postgres":
		dsn = fmt.Sprintf("host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_host, db_user, db_pass, db_port, db_sslmode)
		sqlstring = fmt.Sprintf("CREATE DATABASE %s", db_name)
		break
	case "sqlite3":
		if db_path == "" {
			db_path = "./"
		}
		dsn = fmt.Sprintf("%s%s.db", db_path, db_name)
		os.Remove(dsn)
		sqlstring = "create table init (n varchar(32));drop table init;"
		break
	case "mssql":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", db_user, db_pass, db_host, db_port, db_name)
		sqlstring = fmt.Sprintf("CREATE DATABASE %s", db_name)
		break
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}
	db, err := gorm.Open(db_type, dsn)
	if err != nil {
		panic(err.Error())
	}
	db.Exec(sqlstring)
	log.Println("Database ", db_name, " created")

	defer db.Close()
}

func insertUser() {
	fmt.Println("insert user ...")
	u := new(User)
	u.Username = "admin"
	u.Password = EncryptPasswordMD5("1234")
	DB.Create(&u)
	fmt.Println("insert user end")
}
