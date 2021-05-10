package service

//Load DB config from config.yaml and access to DB

import (
	"github.com/dgrijalva/jwt-go"
	"golang-demo/model/request"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
)

type Data struct {
	SQL SQL
	JWT JWT
}

type SQL struct {
	Username string
	Password string
	DBName string `yaml:db-name`
	DBPort string `yaml:db-port`
}


type JWT struct {
	SigningKey  string `yaml:"signing-key"`    // jwt signature
}

type CustomClaim struct {
	user string
	roll string
	jwt.StandardClaims
}

func getConfig() *Data {
	conf,err := ioutil.ReadFile("config.yaml")
	if err != nil{
		log.Fatal(err)
	}
	d := &Data{}
	err = yaml.Unmarshal(conf,d)
	if err != nil{
		log.Fatal(err)
	}
	return d
}
//dsn format: username:password@protocol(address)/dbname?param=value
func connectToDB(d *Data,u request.User) (string,error){
	dsn := d.SQL.Username+":"+d.SQL.Password+"@tcp("+d.SQL.DBName+")/demo?charset=utf8&parseTime=True&loc=Local"
	//gorm mysql config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	// Open: Using mysql driver and gorm default config
	db,err := gorm.Open(mysql.New(mysqlConfig),&gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}
	var User = request.User{}
	// If sql ORM execute successfully, .Error return nil
	err = db.Where("username = ? and password = ?",u.Username,u.Password).First(&User).Error
	if err != nil {
		return "",err
	}

	var claim = CustomClaim{
		user: u.Username,
		roll: "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 3600,
		},
	}
	key := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	token, err := key.SignedString([]byte(d.JWT.SigningKey))
	return token,nil
}

func GenerateJWT(u request.User) (string,error) {
	return connectToDB(getConfig(),u)
}

