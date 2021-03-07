package initialize

import (
	"github.com/spf13/viper"
	"log"
	"tinyURL/global"
)

func Config() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("配置文件读取失败：", err)
	}
	if err = viper.UnmarshalKey("mysql", &global.DataBaseConf); err != nil {
		log.Println("mysql配置读取失败:", err)
	}
	if err = viper.UnmarshalKey("app", &global.AppConf); err != nil {
		log.Println("app配置读取失败：", err)
	}
	if err = viper.UnmarshalKey("zap", &global.ZapConf); err != nil {
		log.Println("zap配置读取失败：", err)
	}

	if err = viper.UnmarshalKey("qrcode", &global.QRCodeConf); err != nil {
		log.Println("qrcode配置失败", err)
	}
}
