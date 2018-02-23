package models

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
	"github.com/icrowley/fake"
	"strconv"
	"math/rand"
	"time"
)

var DB *gorm.DB


func Shuffle(vals []Service) []Service {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]Service, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}

func GenerateRandomDataset() {
	dbName := beego.AppConfig.String("db_name")
	fmt.Println("drop database ...")
	DB.Exec(fmt.Sprintf("DROP DATABASE %s;", dbName))
	Syncdb()
	companyCount := 10
	managerCount := 400
	storeCount := 200
	var services []Service


	for _, i := range [...]string{"Maintenance", "Wash", "Rent", "Repair", "Gas", "Restaurant", "Hotel"} {
		s := Service{Name:i}
		s.Insert()
		services = append(services, s)
	}

	for i := 1; i <= companyCount; i++ {
		c := Company{
			Name: fake.Company(),//"Company " + strconv.Itoa(i),
			Contact: fake.FullName(),
			PhoneNumber: fake.Phone(),
			Email: fake.EmailAddress(),
		}
		c.ID = uint(i)
		c.Insert()
		source := rand.NewSource(time.Now().UnixNano())
		ran := rand.New(source)
		sc := ran.Intn(3) + 1
		ss := Shuffle(services)[:sc]

		for j := 0; j < storeCount/companyCount; j++ {
			source = rand.NewSource(time.Now().UnixNano())
			ran = rand.New(source)
			s := Store{
				Name: "Store " + strconv.Itoa(j),
				Address: fake.StreetAddress(),
				Latitude: ran.Float64()*90,
				Longitude: ran.Float64()*180,
				PhoneNumber: fake.Phone(),
				CompanyID: uint(i),
			}
			source = rand.NewSource(time.Now().UnixNano())
			ran = rand.New(source)
			ssc := ran.Intn(sc) + 1
			sss := Shuffle(ss)[:ssc]
			s.Services = sss
			s.Insert()
		}
	}

	managerRoleID := FindRoleByName(ROLE_MANAGER).ID
	for i := 2; i <= managerCount + 1; i++ {
		m := User{
			Username: "manager" + strconv.Itoa(i),
			Email: fake.EmailAddress(),
			Name: fake.FullName(),
			RoleID: managerRoleID,
		}
		m.ID = uint(i)
		m.SetPassword("111111")
		source := rand.NewSource(time.Now().UnixNano())
		ran := rand.New(source)
		if ran.Intn(10) > 4 {
			m.StoreID = uint(ran.Intn(storeCount) + 1)
		}
		m.Insert()
	}


}



func LoadlibDB(db *gorm.DB) {
	db.AutoMigrate(&Service{})
	db.AutoMigrate(&Company{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Store{})
	db.AutoMigrate(&Client{})
	db.AutoMigrate(&Vehicle{})
	db.AutoMigrate(&Order{})
}


func Syncdb() {
	createDB()
	if err := Connect(); err != nil {
		beego.Error(err)
		return
	}
	insertRoles()
	insertUser()
	fmt.Println("database init is complete.")
}

//数据库连接
func Connect() error {
	var dsn string
	dbType := beego.AppConfig.String("db_type")
	dbHost := beego.AppConfig.String("db_host")
	dbPort := beego.AppConfig.String("db_port")
	dbUser := beego.AppConfig.String("db_user")
	dbPass := beego.AppConfig.String("db_pass")
	dbName := beego.AppConfig.String("db_name")
	dbPath := beego.AppConfig.String("db_path")
	dbSslmode := beego.AppConfig.String("db_sslmode")
	switch dbType {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
		break
	case "postgres":
		dsn = fmt.Sprintf("dbname=%s host=%s  user=%s  password=%s  port=%s  sslmode=%s", dbName, dbHost, dbUser, dbPass, dbPort, dbSslmode)
		break
	case "sqlite3":
		if dbPath == "" {
			dbPath = "./"
		}
		dsn = fmt.Sprintf("%s%s.db", dbPath, dbName)
		break
	case "mssql":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPass, dbHost, dbPort, dbName)
		break
	default:
		beego.Critical("Database driver is not allowed:", dbType)
	}
	db, err := gorm.Open(dbType, dsn)

	if err == nil {
		DB = db
		LoadlibDB(db)
		if len(FindRoles()) == 0 {
			insertRoles()
			insertUser()
		}

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

func insertRoles() {
	fmt.Println("insert roles ...")
	admin := Role{
		Name: ROLE_ADMIN,
	}
	admin.Insert()
	manager := Role{
		Name: ROLE_MANAGER,
	}
	manager.Insert()
	fmt.Println("insert roles end")
}


func insertUser() {
	fmt.Println("insert user ...")
	u := User{
		Username: "admin",
		Email: "admin@zooli.com",
		Name: "Boss",
		Role: FindRoleByName(ROLE_ADMIN),
	}
	u.SetPassword("1234")
	u.Insert()
	fmt.Println("insert user end")
}
