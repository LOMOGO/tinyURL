package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		global.Logger.Fatal("未连接到数据库", zap.Error(err))
	}
	err = global.DB.AutoMigrate(&model.URLInfo{})
	if err != nil {
		global.Logger.Warn("数据表自动绑定失败", zap.Error(err))
	}
	sqlDB, err := global.DB.DB()
	if err != nil {
		global.Logger.Fatal("数据库初始化错误", zap.Error(err))
	}
	sqlDB.SetConnMaxIdleTime(time.Minute * 4)
	sqlDB.SetMaxOpenConns(global.DataBaseConf.MaxOpenConn)
	sqlDB.SetMaxIdleConns(global.DataBaseConf.MaxIdleConn)
}
