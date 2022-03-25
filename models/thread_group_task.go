package models

import (
	"elder-wand/enums"
	"elder-wand/utils/dbUtils"
)

// ThreadGroupTask 线程组任务
type ThreadGroupTask struct {
	BaseModel
	ThreadGroupID dbUtils.SFID                // 线程组ID
	Status        enums.ThreadGroupTaskStatus // 任务状态
	TaskID        dbUtils.SFID                // 任务ID
}

func (*ThreadGroupTask) TableName() string {
	return "t_thread_group_task"
}
