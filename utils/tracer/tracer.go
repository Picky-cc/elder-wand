package tracer

import (
	"context"
	"elder-wand/settings"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmsql"
	"log"
	"sync"
)

var Tracer *apm.Tracer

func Init() {
	//if !settings.Config.EnableAPM {
	//	return
	//}
	tracer, err := apm.NewTracer(settings.Config.AppName, "") //version.VERSION)
	if err != nil {
		log.Fatal(err)
	}
	Tracer = tracer
}

type TracerContext struct {
	ctx  context.Context
	once sync.Once
}

func (c *TracerContext) GetContext() context.Context {
	c.once.Do(func() {
		c.ctx = context.Background()
	})
	return c.ctx
}

func (c *TracerContext) StartTransaction(name, transactionType string) *apm.Transaction {
	tx := Tracer.StartTransaction(name, "plugin")
	tx.Context.SetLabel("EW_NODE_ID", settings.Config.EwNodeID)
	ctx := c.GetContext()
	c.ctx = apm.ContextWithTransaction(ctx, tx)
	return tx
}

func (c *TracerContext) WrapFunction(name string, parent *apm.Span, fn func(ctx context.Context)) {
	tx := apm.TransactionFromContext(c.ctx)
	span := tx.StartSpan(name, "function", parent)
	defer span.End()
	fn(c.ctx)
}

func (c *TracerContext) SendError(err error) {
	e := apm.CaptureError(c.ctx, err)
	e.Send()
}

func GormExec(conn *gorm.DB, sql string, values ...interface{}) *gorm.DB {
	value, ok := conn.Get("elasticapm:context")
	if !ok {
		conn = conn.Exec(sql, values...)
		return conn
	}
	ctx := value.(context.Context)
	driverName := conn.Dialect().GetName()
	switch driverName {
	case "postgres":
		driverName = "postgresql"
	}
	spanName := apmsql.QuerySignature(sql)
	spanType := fmt.Sprintf("db.%s.exec", driverName)
	span, _ := apm.StartSpan(ctx, spanName, spanType)
	defer span.End()
	// span.Context.SetDestinationAddress(dsnInfo.Address, dsnInfo.Port)
	span.Context.SetDatabase(apm.DatabaseSpanContext{
		// Instance:  dsnInfo.Database,
		Statement: sql,
		Type:      "sql",
		// User:      dsnInfo.User,
	})
	conn = conn.Exec(sql, values...)
	span.Context.SetLabel("RowsAffected", conn.RowsAffected)
	return conn
}
