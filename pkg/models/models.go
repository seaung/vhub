package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/seaung/vhub/pkg/config"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.SConfig.Databases.DbUser,
		config.SConfig.Databases.DbName,
		config.SConfig.Databases.Host,
		config.SConfig.Databases.DbName, config.SConfig.Databases.Charset)
	fmt.Println("dsn == ", dsn)
	//db, err = gorm.Open("mysql", "user:pass@(localhost)/vuln_db?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("Can't connect to database")
	}
	db.Debug()
	db.LogMode(true)
	db.AutoMigrate(&Users{}, &Message{})
}
