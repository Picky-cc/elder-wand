package models

import (
	"elder-wand/enums"
	"elder-wand/utils/dbUtils"
)

// Task 任务表
type Task struct {
	BaseModel
	ProjectID dbUtils.SFID        //机构ID
	Name      string              //任务名称
	Type      enums.TaskType      //类型
	LifeCycle enums.TaskLifeCycle //生命周期
}

func (*Task) TableName() string {
	return "t_task"
}
