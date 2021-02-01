package model

type Database struct {
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	DBName      string `mapstructure:"dbname"`
	MaxOpenConn int    `mapstructure:"maxopenconn"`
	MaxIdleConn int    `mapstructure:"maxidleconn"`
}

type App struct {
	URL  string `mapstructure:"url"`
	Port string `mapstructure:"port"`
}
