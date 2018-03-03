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
	clientCount := 5000

	var services []Service
	statuses := [...]Status{ORDERED, IN_SERVICE, END_SERVICE, WAITING_FOR_PAYMENT, FINISHED}

	startTime := time.Date(2017, 6, 1,0, 0, 0, 0, time.UTC).Unix()
	endTime := time.Now().Unix()
	diffTime := int(endTime - startTime)

	source := rand.NewSource(time.Now().UnixNano())
	ran := rand.New(source)

	for _, i := range [...]string{"Maintenance", "Wash", "Rent", "Repair", "Gas", "Restaurant", "Hotel"} {
		s := Service{Name:i}
		s.Insert()
		services = append(services, s)
	}

	for i := 1; i <= clientCount; i++ {
		c := Customer{
			Name: fake.FullName(),
			PhoneNumber: fake.Phone(),
			Email: fake.EmailAddress(),
		}
		c.SetPassword("111111")

		c.Insert()

		vCount := ran.Intn(3)
		for j := 0; j < vCount ; j++ {
			v := Vehicle{
				Model: fake.Model(),
				OwnerID: c.ID,
				Plate: string(c.Name[0]) + strconv.Itoa(int(c.ID)) + strconv.Itoa(j+1) ,
			}
			v.Insert()
		}
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

		sc := ran.Intn(3) + 1
		ss := Shuffle(services)[:sc]

		for j := 0; j < storeCount/companyCount; j++ {
			s := Store{
				Name: "Store " + strconv.Itoa(j),
				Address: fake.StreetAddress(),
				Latitude: float64(fake.Latitude()),
				Longitude: float64(fake.Longitude()),
				PhoneNumber: fake.Phone(),
				CompanyID: uint(i),
				Image: "",
			}
			ssc := ran.Intn(sc) + 1
			sss := Shuffle(ss)[:ssc]
			s.Services = sss
			s.Insert()

			pCount := ran.Intn(4)
			for k := 0; k < pCount; k++ {
				p := Product{
					Name: fake.ProductName(),
					Value: float64(int(ran.Float64() * 10000)) / 100,
					StoreID: s.ID,
					ServiceID: sss[ran.Intn(len(sss))].ID,
					Description: fake.Paragraph(),
				}

				p.Insert()

				oCount := ran.Intn(200)
				for m := 0; m < oCount; m++ {
					o := Order{
						CustomerID: uint(ran.Intn(clientCount) + 1),
						Fee: p.Value,
						ProductID: p.ID,
					}
					var targetStatus int
					if ran.Intn(10) > 5 {
						targetStatus = len(statuses) - 1
					} else {
						targetStatus = ran.Intn(len(statuses))
					}
					tmpT := time.Unix(startTime + int64(ran.Intn(diffTime)), 0)
					for n := 0; n <= targetStatus; n++ {
						tmpT = tmpT.Add( time.Hour * time.Duration(ran.Intn(24)))
						ol := OrderLog{
							Status: statuses[n],
							Timestamp: tmpT,
						}
						o.Logs = append(o.Logs, ol)
						if ran.Intn(20) < 2 {
							olc := OrderLog{
								Status: CANCELED,
								Timestamp: tmpT.Add( time.Hour * time.Duration(ran.Intn(24))),
							}
							o.Logs = append(o.Logs, olc)
							break
						}
					}
					o.Status = o.Logs[len(o.Logs)-1].Status

					o.Insert()

				}

			}

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
		if ran.Intn(10) > 2 {
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
	db.AutoMigrate(&Customer{})
	db.AutoMigrate(&Vehicle{})
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&OrderLog{})
}


func Syncdb() {
	createDB()
	//insertRoles()
	//insertUser()
	if err := Connect(); err != nil {
		beego.Error(err)
		return
	}
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

	dbType := beego.AppConfig.String("db_type")
	dbHost := beego.AppConfig.String("db_host")
	dbPort := beego.AppConfig.String("db_port")
	dbUser := beego.AppConfig.String("db_user")
	dbPass := beego.AppConfig.String("db_pass")
	dbName := beego.AppConfig.String("db_name")
	dbPath := beego.AppConfig.String("db_path")
	dbSslmode := beego.AppConfig.String("db_sslmode")

	var dsn string
	var sqlstring string
	switch dbType {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort)
		sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", dbName)
		break
	case "postgres":
		dsn = fmt.Sprintf("host=%s  user=%s  password=%s  port=%s  sslmode=%s", dbHost, dbUser, dbPass, dbPort, dbSslmode)
		sqlstring = fmt.Sprintf("CREATE DATABASE %s", dbName)
		break
	case "sqlite3":
		if dbPath == "" {
			dbPath = "./"
		}
		dsn = fmt.Sprintf("%s%s.db", dbPath, dbName)
		os.Remove(dsn)
		sqlstring = "create table init (n varchar(32));drop table init;"
		break
	case "mssql":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPass, dbHost, dbPort, dbName)
		sqlstring = fmt.Sprintf("CREATE DATABASE %s", dbName)
		break
	default:
		beego.Critical("Database driver is not allowed:", dbType)
	}
	db, err := gorm.Open(dbType, dsn)
	if err != nil {
		panic(err.Error())
	}
	db.Exec(sqlstring)
	log.Println("Database ", dbName, " created")

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
