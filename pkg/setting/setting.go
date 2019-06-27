package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type AppConf struct {
	JwtSecret       string
	PageSize int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string
	LogSavePath    string
	LogSaveName    string
	LogFileExt     string
	TimeFormat     string
}

type ServerConf struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DatabaseConf struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

// var 方式定义全局变量
var App = &AppConf{}
var Server = &ServerConf{}
var Database = &DatabaseConf{}

func Setup() {
	file, e := ini.Load("conf/app.ini")
	if e != nil {
		log.Fatalf("Fail to parse app.ini: %v", e)
	}

	e = file.Section("app").MapTo(App)
	if e != nil {
		log.Fatalf("app.ini MapTo App err: %v", e)
	}

	App.ImageMaxSize = App.ImageMaxSize * 1024 * 1024

	e = file.Section("server").MapTo(Server)
	if e != nil {
		log.Fatalf("app.ini MapTo Server err: %v", e)
	}

	Server.ReadTimeout = Server.ReadTimeout * time.Second
	Server.WriteTimeout = Server.WriteTimeout * time.Second

	e = file.Section("database").MapTo(Database)
	if e != nil {
		log.Fatalf("app.ini MapTo Database err: %v", e)
	}
}
