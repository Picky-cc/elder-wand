package models

import (
	"elder-wand/enums"
)

type ThreadGroup struct {
	BaseModel
	Name         string                     // 名称
	ThreadCount  int                        // 线程数量
	EwNode       string                     // 节点，保留字段，现无用
	SleepSeconds int                        // 空闲时，休眠时长，单位为秒
	LifeCycle    enums.ThreadGroupLifeCycle // 生命周期，有效，作废

}

func (*ThreadGroup) TableName() string {
	return "t_thread_group"
}
