package enums

type TaskLifeCycle int

const (
	TaskLifeCycleActive   TaskLifeCycle = 1 // 有效
	TaskLifeCycleFinished TaskLifeCycle = 2 // 已终止
	TaskLifeCycleCancel   TaskLifeCycle = 3 // 作废
)

type TaskType int

const (
	TaskTypeOfOne   TaskType = 1 //任务类型一
	TaskTypeOfTwo   TaskType = 2 //任务类型二
	TaskTypeOfThree TaskType = 3 //任务类型三
)
