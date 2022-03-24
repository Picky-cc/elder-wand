package enums

type ThreadGroupLifeCycle int

const (
	ThreadGroupLifeCycleActive ThreadGroupLifeCycle = 1
	ThreadGroupLifeCycleCancel ThreadGroupLifeCycle = 2
)

type ThreadGroupTaskStatus int

const (
	ThreadTaskStatusWaiting ThreadGroupTaskStatus = 1
	ThreadTaskStatusRunning ThreadGroupTaskStatus = 2
)
