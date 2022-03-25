package services

import (
	"context"
	"elder-wand/db"
	"elder-wand/enums"
	"elder-wand/errors"
	"elder-wand/models"
	"elder-wand/utils/dbUtils"
	"github.com/jinzhu/gorm"
)

type taskService struct {
}

var TaskService taskService

func (a *taskService) Create(conn *gorm.DB, task models.Task) *errors.Error {
	err := conn.Create(&task).Error
	if err != nil {
		return errors.NewDBError(err.Error())
	}
	return nil
}

func (a *taskService) GetTaskByProjectAndTaskType(ctx context.Context, projectID dbUtils.SFID, tType enums.TaskType) (*models.Task, *errors.Error) {
	task := models.Task{}
	conn := db.EwDB.NewConnection(ctx)

	sql := ` select * from t_task where project_id = ? and type = ? `
	conn = conn.Raw(sql, projectID, tType).Find(&task)
	if conn.RecordNotFound() {
		return nil, nil
	} else if conn.Error != nil {
		return nil, errors.NewDBError(conn.Error.Error())
	}
	return &task, nil
}

func (a *taskService) GetTaskByID(ctx context.Context, ID dbUtils.SFID) (*models.Task, *errors.Error) {
	task := models.Task{}
	conn := db.EwDB.NewConnection(ctx)

	sql := ` select * from t_task where id=? `
	conn = conn.Raw(sql, ID).Find(&task)
	if conn.RecordNotFound() {
		return nil, nil
	} else if conn.Error != nil {
		return nil, errors.NewDBError(conn.Error.Error())
	}
	return &task, nil
}

func (a *taskService) GetActiveTaskByID(ctx context.Context, ID dbUtils.SFID, lifeCycle enums.TaskLifeCycle) (*models.Task, *errors.Error) {
	task := models.Task{}
	conn := db.EwDB.NewConnection(ctx)

	sql := ` select * from t_task where id=? and life_cycle = ? `
	conn = conn.Raw(sql, ID, lifeCycle).Find(&task)
	if conn.RecordNotFound() {
		return nil, errors.NewNotFoundErrorf("Task %d not exist", ID)
	} else if conn.Error != nil {
		return nil, errors.NewDBError(conn.Error.Error())
	}
	return &task, nil
}

func (a *taskService) GetActiveTask(conn *gorm.DB) ([]models.Task, *errors.Error) {
	list := make([]models.Task, 0)

	sql := ` select * from t_task where life_cycle = ? `
	err := conn.Raw(sql, enums.TaskLifeCycleActive).Find(&list).Error
	if err != nil {
		return list, errors.NewDBError(conn.Error.Error())
	}
	return list, nil
}
