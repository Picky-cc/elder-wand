package db

import (
	"context"
	"elder-wand/enums"
	"elder-wand/errors"
	"elder-wand/settings"
	"elder-wand/utils/dbUtils"
	"elder-wand/utils/log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.elastic.co/apm/module/apmgorm"
	"regexp"
	"strings"
	"sync"
	"time"
)

type DBInstance struct {
	gdb *gorm.DB

	CashFlowRemarkHighCounter uint64
	CashFlowRemarkLowCounter  uint64
	Lock                      sync.Mutex
}

type dbLogger struct {
}

func formatSQL(sql string) string {
	nSQL := strings.ReplaceAll(sql, "\n", " ")
	nSQL = strings.TrimSpace(nSQL)

	replaceWhiteSpaces := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	nSQL = replaceWhiteSpaces.ReplaceAllString(nSQL, " ")

	return nSQL
}

func (*dbLogger) Print(v ...interface{}) {
	if len(v) >= 6 {
		sql := formatSQL(fmt.Sprint(v[3]))
		log.Debugf("%s ( %+v ) [%d rows affected or returned] [%+v], %+v", sql, v[4], v[5], v[2], v[1])
	}
}

var EwDB *DBInstance

func (instance *DBInstance) NewConnection(ctx context.Context) *gorm.DB {
	db := instance.gdb.New()
	if !settings.Config.EnableAPM {
		return db
	}
	db = apmgorm.WithContext(ctx, db)
	return db
}

func (instance *DBInstance) Init(config *settings.DBConfigure) {
	var (
		db  *gorm.DB
		err error
	)
	if settings.Config.EnableAPM {
		db, err = apmgorm.Open("mysql", config.ConnectionUri)
	} else {
		db, err = gorm.Open("mysql", config.ConnectionUri)
	}
	if err != nil {
		panic(err)
	}
	db.SetLogger(&dbLogger{})

	db.DB().SetMaxIdleConns(config.MaxIdle)
	db.DB().SetMaxOpenConns(config.MaxOpen)

	maxLeftTime := time.Duration(config.ConnMaxLeftTime)
	if maxLeftTime < 30*time.Minute && maxLeftTime != 0 {
		maxLeftTime = 30 * time.Minute
	}
	db.DB().SetConnMaxLifetime(maxLeftTime)

	db.LogMode(settings.Config.DBLog)

	instance.initCallback(db)

	instance.gdb = db
}

func (instance *DBInstance) initCallback(db *gorm.DB) {
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("updated", time.Now())
	}
}

func Init() {
	dbUtils.Init()
	EwDB = &DBInstance{}
	EwDB.Init(&settings.Config.DBConfig)

	initDefaultThreadGroup(EwDB.gdb)
	resetThreadGroupTaskStatus(EwDB.gdb, settings.Config.EwNodeID)
}

func resetThreadGroupTaskStatus(conn *gorm.DB, ewNodeID int) (int64, *errors.Error) {
	sql := `
	update t_thread_group_task set status=?, updated=now() where status=? and thread_group_id in (select id from t_thread_group where ew_node=?)
`
	conn = conn.Exec(sql, enums.ThreadTaskStatusWaiting, enums.ThreadTaskStatusRunning, fmt.Sprint(ewNodeID))
	if conn.Error != nil {
		return 0, errors.NewDBError(conn.Error.Error())
	}
	return conn.RowsAffected, nil
}
