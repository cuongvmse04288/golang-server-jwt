package service

import (
	"golang-demo/model"
	"golang-demo/model/request"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//dsn format: username:password@protocol(address)/dbname?param=value
//Load DB config from config.yaml and connect to DB
func connectToDB(c *model.Config) (*gorm.DB, error) {
	dsn := c.SQL.Username + ":" + c.SQL.Password + "@tcp(" + c.SQL.DBName +
		")/demo?charset=utf8&parseTime=True&loc=Local"
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
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func VerifyLogin(config *model.Config, user request.User) (bool, error) {
	var db, err = connectToDB(config)
	if err != nil {
		return false, err
	}
	var u = request.User{}
	// If sql ORM execute successfully, .Error return nil
	err = db.Where("username = ? and password = ?", user.Username, user.Password).First(&u).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
