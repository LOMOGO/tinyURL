package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
	"tinyURL/global"
	"tinyURL/model"
)

func DataBase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.DataBaseConf.User,
		global.DataBaseConf.Password,
		global.DataBaseConf.Host,
		global.DataBaseConf.Port,
		global.DataBaseConf.DBName,
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("未连接到数据库")
	}
	err = global.DB.AutoMigrate(&model.URLInfo{})
	if err != nil {
		log.Println(err)
	}
	sqlDB, err := global.DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetConnMaxIdleTime(time.Minute * 4)
	sqlDB.SetMaxOpenConns(global.DataBaseConf.MaxOpenConn)
	sqlDB.SetMaxIdleConns(global.DataBaseConf.MaxIdleConn)
}
