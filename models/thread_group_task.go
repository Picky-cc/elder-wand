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
	AgreementID   dbUtils.SFID                // 协议号
}

func (*ThreadGroupTask) TableName() string {
	return "t_clearing_thread_group_task"
}
