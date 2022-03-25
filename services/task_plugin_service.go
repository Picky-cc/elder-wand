package services

import (
	"context"
	"elder-wand/db"
	"elder-wand/enums"
	"elder-wand/errors"
	"elder-wand/models"
	"elder-wand/utils/dbUtils"
	"github.com/jinzhu/gorm"
	"time"
)

type taskPluginService struct {
}

var TaskPluginService taskPluginService

func (a *taskPluginService) Create(conn *gorm.DB, taskPlugin models.TaskPlugin) *errors.Error {
	err := conn.Create(&taskPlugin).Error
	if err != nil {
		return errors.NewDBError(err.Error())
	}
	return nil
}

func (a *taskPluginService) GetTaskPluginListByTask(ctx context.Context, taskID dbUtils.SFID, lifeCycle enums.CommonLifeCycle, queryTime *time.Time) ([]models.TaskPlugin, *errors.Error) {
	results := make([]models.TaskPlugin, 0)

	sql := "select * from t_task_plugin where task_id=? and life_cycle=? and next_query_time <= ? "

	conn := db.ClearingDB.NewConnection(ctx)
	err := conn.Raw(sql, taskID, lifeCycle, queryTime).Scan(&results).Error
	if err != nil {
		return nil, errors.NewDBError(err.Error())
	}
	return results, nil
}

func (a *taskPluginService) UpdateTaskPluginQueryTime(conn *gorm.DB, serviceID dbUtils.SFID, queryTime time.Time) (int64, *errors.Error) {
	sql := "update t_task_plugin set last_query_time=?, next_query_time=date_add(?, interval `interval` second) where id=?"
	conn = conn.Exec(sql, queryTime, queryTime, serviceID)
	if conn.Error != nil {
		return 0, errors.NewDBError(conn.Error.Error())
	}
	return conn.RowsAffected, nil
}
