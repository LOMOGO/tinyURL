package config

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

type Zap struct {
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxsize"`
	MaxBackups int    `mapstructure:"maxbackups"`
	MaxAge     int    `mapstructure:"maxage"`
	Compress   bool   `mapstructure:"compress"`
	Level      string `mapstructure:"debug"`
}

type QRCode struct {
	ImgHeight int    `mapstructure:"imgheight"`
	FilePath  string `mapstructure:"filepath"`
}
