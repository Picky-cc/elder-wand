package controllers

import (
	"elder-wand/errors"
	"elder-wand/utils/tracer"
	"fmt"
	"runtime"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go.elastic.co/apm"
)

type BaseController struct {
	beego.Controller
}

func (b *BaseController) HandlerPanicError() {
	if err := recover(); nil != err && err != beego.ErrAbort {
		e := tracer.Tracer.Recovered(err)
		e.SetTransaction(apm.TransactionFromContext(b.Ctx.Request.Context()))
		e.Send()
		buf := make([]byte, 8192)
		n := runtime.Stack(buf, false)
		stackTraces := fmt.Sprintf("%s", buf[:n])
		logs.Error(fmt.Sprintf("system panic: %+v", stackTraces))

		b.Error(errors.NewUnknownError(stackTraces))
	}
}

func (b *BaseController) Success(data interface{}) {
	b.response(errors.Success, "", data)
}

func (b *BaseController) SuccessList(list interface{}, total int64) {
	b.SuccessListWithExtraData(list, total, nil)
}
func (b *BaseController) SuccessListWithExtraData(list interface{}, total int64, extraData map[string]interface{}) {
	data := map[string]interface{}{
		"List":       list,
		"TotalCount": total,
	}
	for k, v := range extraData {
		data[k] = v
	}
	b.response(errors.Success, "", data)
}

func (b *BaseController) Error(err error) {
	b.ErrorData(err, nil)
}

func (b *BaseController) ErrorData(err error, data interface{}) {
	switch t := err.(type) {
	case *errors.Error:
		b.response(t.Code, t.Detail, data)
	default:
		b.response(errors.UnknownError, err.Error(), data)
	}
}

func (b *BaseController) response(code errors.ResponseCode, errorDetail string, data interface{}) {
	msg := errors.GetMessage(code)
	if code == errors.Success {
		b.Ctx.Input.SetData("prom:resp_status", "succ")
	} else {
		b.Ctx.Input.SetData("prom:resp_status", "fail")
		b.Ctx.Input.SetData("prom:fail_msg", msg)
	}
	if len(errorDetail) > 0 {
		msg = fmt.Sprintf("%s|%s", msg, errorDetail)
	}

	b.Data["json"] = &errors.JSONResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	b.ServeJSON()
	b.StopRun()
}
