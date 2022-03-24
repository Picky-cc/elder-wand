package enums

//清算状态
type ClearStatus int

const (
	ClearStatusPending ClearStatus = 1 // 未清算
	ClearStatusPartial ClearStatus = 2 // 部分清算
	ClearStatusFully   ClearStatus = 3 // 全额清算
)

type CommonLifeCycle int

const (
	LifeCycleActive   CommonLifeCycle = 1 // 有效
	LifeCycleCanceled CommonLifeCycle = 2 // 作废
)
