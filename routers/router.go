package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var apiV1 = beego.NewNamespace("/api/v1",
	beego.NSCond(func(ctx *context.Context) bool {
		if ua := ctx.Request.UserAgent(); ua != "" {
			return true
		}
		return false
	}),
	ApiV1,
	//clearFlowV1,
)

func Init() {
	beego.AddNamespace(apiV1)
}
