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

// InitDB initializes the database connection and configurations
func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.SConfig.Databases.DbUser,
		config.SConfig.Databases.DbPass,
		config.SConfig.Databases.Host,
		config.SConfig.Databases.DbPort,
		config.SConfig.Databases.DbName,
		config.SConfig.Databases.Charset)

	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// 设置连接池配置
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)

	// 开启调试模式
	db.Debug()
	db.LogMode(true)

	// 自动迁移数据库表结构
	if err := db.AutoMigrate(&Users{}, &Message{}).Error; err != nil {
		return fmt.Errorf("failed to auto migrate tables: %v", err)
	}

	return nil
}
