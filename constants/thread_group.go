package constants

type ThreadGroupLifeCycle int

const (
	ThreadGroupLifeCycleActive ThreadGroupLifeCycle = 1
	ThreadGroupLifeCycleCancel ThreadGroupLifeCycle = 2
)

type TaskTaskStatus int

const (
	ThreadTaskStatusWaiting TaskTaskStatus = 1
	ThreadTaskStatusRunning TaskTaskStatus = 2
)
