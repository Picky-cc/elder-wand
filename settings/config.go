package settings

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type DBConfigure struct {
	DriverName              string
	ConnectionUri           string
	MaxIdle                 int
	MaxOpen                 int
	ConnMaxLeftTime         int
	BatchInsertRecordsCount int
	RunMigration            bool
}

type appConfig struct {
	AppName          string
	HttpPort         string
	RunMode          string
	LogsPath         string
	BeegoLogName     string
	LogName          string
	ErrorLogName     string
	FileMaxSize      int
	MaxAge           int
	DBLog            bool
	DBClearingConfig DBConfigure
	EwNodeID         int
	EnableAPM        bool
}

func (config *appConfig) init() {
	config.AppName = beego.AppConfig.String("appname")
	config.HttpPort = beego.AppConfig.String("httpport")
	config.RunMode = beego.AppConfig.String("runmode")
	config.LogsPath = beego.AppConfig.String("LogsPath")
	config.BeegoLogName = beego.AppConfig.String("BeegoLogName")
	config.LogName = beego.AppConfig.String("LogName")
	config.ErrorLogName = beego.AppConfig.String("ErrorLogName")
	config.FileMaxSize = beego.AppConfig.DefaultInt("FileMaxSize", 100)
	config.MaxAge = beego.AppConfig.DefaultInt("MaxAge", 14)
	config.DBLog = beego.AppConfig.DefaultBool("DBLog", false)

	config.DBClearingConfig.DriverName = beego.AppConfig.String("DBClearingDriverName")
	config.DBClearingConfig.ConnectionUri = beego.AppConfig.String("DBClearingConnectionUri")
	config.DBClearingConfig.MaxIdle = beego.AppConfig.DefaultInt("DBClearingMaxIdle", 20)
	config.DBClearingConfig.MaxOpen = beego.AppConfig.DefaultInt("DBClearingMaxOpen", 100)
	config.DBClearingConfig.ConnMaxLeftTime = beego.AppConfig.DefaultInt("DBClearingMaxConnLeftTime", 1800000000000)
	config.DBClearingConfig.BatchInsertRecordsCount = beego.AppConfig.DefaultInt("DBClearingBatchInsertRecordsCount", 500)
	config.DBClearingConfig.RunMigration = beego.AppConfig.DefaultBool("DBClearingRunMigration", false)

	config.EwNodeID = beego.AppConfig.DefaultInt("EwNodeID", 0)
	config.EnableAPM = beego.AppConfig.DefaultBool("EnableAPM", false)
}

func (config *appConfig) IsDev() bool {
	return config.RunMode == "dev"
}

func (config *appConfig) IsTest() bool {
	return config.RunMode == "test"
}

func (config *appConfig) IsProd() bool {
	return config.RunMode == "prod"
}

var Config appConfig

func initConfig() *appConfig {
	err := beego.LoadAppConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}

	Config.init()
	configInfo, _ := json.Marshal(&Config)
	fmt.Printf("initialized settings success! config info is %s\n", string(configInfo))

	return &Config
}
