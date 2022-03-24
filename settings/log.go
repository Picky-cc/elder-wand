package settings

import (
	"elder-wand/utils/log"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/logs"
)

func initLog() {
	if err := os.MkdirAll(Config.LogsPath, os.ModePerm); err != nil {
		panic(err)
	}

	if Config.IsProd() {
		logs.SetLevel(logs.LevelDebug)
	} else if Config.IsDev() {
		logs.SetLevel(logs.LevelDebug)
		log.Development = true
	} else {
		logs.SetLevel(logs.LevelInfo)
		log.Development = false
	}
	// 显示输出文件名和行号
	logs.EnableFuncCallDepth(true)

	config := make(map[string]interface{}, 1)
	config["filename"] = filepath.Join(Config.LogsPath, Config.BeegoLogName)
	config["maxdays"] = 7
	configJsonStr, _ := json.Marshal(config)

	// 同时输出到文件和控制台
	_ = logs.SetLogger(logs.AdapterFile, string(configJsonStr))
	_ = logs.SetLogger(logs.AdapterConsole)
	if Config.IsDev() {
		logs.SetLevel(logs.LevelDebug)
	} else {
		logs.SetLevel(logs.LevelInformational)
	}

	log.FileName = filepath.Join(Config.LogsPath, Config.LogName)
	log.ErrorFileName = filepath.Join(Config.LogsPath, Config.ErrorLogName)
	log.FileMaxSize = Config.FileMaxSize
	log.MaxAge = Config.MaxAge
	log.Init()

	log.Infow("init log success! log config:",
		"Development", log.Development,
		"Level", log.Level,
		"FileName", log.FileName,
		"ErrorFileName", log.ErrorFileName,
		"FileMaxSize", log.FileMaxSize,
		"MaxAge", log.MaxAge,
	)
}
