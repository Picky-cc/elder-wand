package main

import (
	"elder-wand/db"
	"elder-wand/routers"
	"elder-wand/settings"
	"elder-wand/tasks"
	"elder-wand/utils/tracer"

	"github.com/astaxie/beego"
	"go.elastic.co/apm/module/apmbeego"
)

func initApp() {
	settings.Init()
	routers.Init()
	db.Init()
	tracer.Init()
}

func runApp() {
	tasks.Start()
	if settings.Config.EnableAPM {
		opt := apmbeego.WithTracer(tracer.Tracer)
		beego.RunWithMiddleWares("", apmbeego.Middleware(opt))
	} else {
		beego.Run()
	}
}

func main() {
	initApp()
	runApp()
}
