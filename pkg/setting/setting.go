package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
)

type AppConf struct {
	JwtSecret       string
	PageSize        int
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

type RedisConf struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

// var 方式定义全局变量
var App = &AppConf{}
var Server = &ServerConf{}
// 设置默认值
var Database = &DatabaseConf{Type: "mysql"}
var Redis = &RedisConf{}

var config *ini.File

func Setup() {
	var err error
	config, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse app.ini: %v", err)
	}

	mapTo("app", App)
	mapTo("redis", Redis)
	mapTo("server", Server)
	mapTo("database", Database)
	fmt.Println(Database)
	App.ImageMaxSize = App.ImageMaxSize * 1024 * 1024

	Redis.IdleTimeout = Redis.IdleTimeout * time.Second

	Server.ReadTimeout = Server.ReadTimeout * time.Second
	Server.WriteTimeout = Server.WriteTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := config.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("config.MapTo %s err: %v", section, err)
	}
}
