package tasks

import (
	"context"
	"elder-wand/db"
	"elder-wand/enums"
	"elder-wand/errors"
	"elder-wand/models"
	"elder-wand/services"
	"elder-wand/settings"
	"elder-wand/tasks/manager"
	"elder-wand/utils/dbUtils"
	"elder-wand/utils/log"
	"elder-wand/utils/tracer"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
)

type ThreadGroupManager struct {
	isRunning bool
	lock      sync.Mutex
	wg        sync.WaitGroup
	workers   []*threadGroup
}

// ============================================================================================================
// start
// ============================================================================================================

func (manager *ThreadGroupManager) Start() {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	if manager.isRunning {
		log.Warnf("[ThreadGroupManager] has started")
		return
	}

	log.Info("[ThreadGroupManager] thread group begin to start..")
	defer log.Info("[ThreadGroupManager] thread group finish to start..")

	manager.isRunning = true
	threadGroups, err := services.ThreadGroupService.GetActiveThreadGroupsByNodeID(context.TODO(), settings.Config.EwNodeID)
	if err != nil {
		panic(err)
	}
	manager.workers = make([]*threadGroup, 0)

	for _, group := range threadGroups {
		g := threadGroup{}
		g.Init(group)
		manager.workers = append(manager.workers, &g)
		manager.wg.Add(1)
		go func(g *threadGroup) {
			log.Infof("[ThreadGroupManager] group [%s, %s] start..", g.group.Name, g.group.EwNode)
			g.Run()
			log.Infof("[ThreadGroupManager] group [%s, %s] end!!", g.group.Name, g.group.EwNode)
			manager.wg.Done()
		}(&g)
	}
}

func (g *threadGroup) Init(gp models.ThreadGroup) {
	g.group = gp
	g.tasksChan = make(chan taskPluginGroupStruct, gp.ThreadCount)
	g.quitChan = make(chan int, gp.ThreadCount)
}

func (g *threadGroup) Run() {
	g.lock.Lock()
	if g.isRunning {
		g.lock.Unlock()
		return
	}
	g.isRunning = true

	g.startWorkers()

	g.lock.Unlock()

	g.runProducer()

	close(g.tasksChan)
}

func (g *threadGroup) startWorkers() {
	for i := 0; i < g.group.ThreadCount; i++ {
		g.workerGroup.Add(1)
		go func(a int) {
			defer g.workerGroup.Done()
			for {
				select {
				case taskPluginGroup, ok := <-g.tasksChan:
					if !ok {
						log.Infof("[threadGroup %s]  worker %d, exit with channel closed", g.group.Name, a)
						return
					}
					taskPluginGroup.workerNo = a
					g.processTaskPluginGroup(taskPluginGroup)
					taskPluginGroup.wg.Done()
				case <-g.quitChan:
					log.Infof("[threadGroup %s] worker %d, exit with quit signal", g.group.Name, a)
					return
				}
			}
		}(i)
	}
}

func (g *threadGroup) processTaskPluginGroup(taskPluginGroup taskPluginGroupStruct) {
	defer g.HandlerPanicError()

	log.Infof("[threadGroup %s] worker %d start to run", g.group.Name, taskPluginGroup.workerNo)
	defer func() {
		log.Infof("[threadGroup %s] worker %d start to run", g.group.Name, taskPluginGroup.workerNo)
	}()
	for _, s := range taskPluginGroup.taskPlugins {
		if err := g.runDataService(&taskPluginGroup.task, &s); err != nil {
			log.Infof("[threadGroup %s] plugin [%d] runDataService occur error [%s]", g.group.Name, s.ID, err.Error())
			// todo: 这里到底是break还是continue值得商榷，可能以后会走配置
			continue
		}
	}
}

func (g *threadGroup) runDataService(task *models.Task, taskPlugin *models.TaskPlugin) *errors.Error {
	defer g.HandlerPanicError()
	defer func() {
		now := time.Now()
		_, _ = services.TaskPluginService.UpdateTaskPluginQueryTime(db.ClearingDB.NewConnection(context.TODO()), taskPlugin.ID, now)
	}()

	log.Infof("[threadGroup %s] start to run service: %d, plugin: %d", g.group.Name, taskPlugin.ID, taskPlugin.ServicePlugin)
	defer log.Infof("[threadGroup %s] end to run service: %d, plugin: %d", g.group.Name, taskPlugin.ID, taskPlugin.ServicePlugin)

	plugin := manager.PluginManager.GetPlugin(taskPlugin.ServicePlugin)
	if plugin == nil {
		msg := fmt.Sprintf("[threadGroup %s] not exist plugins matched ServicePlugin: %d", g.group.Name, taskPlugin.ServicePlugin)
		return errors.NewUnknownError(msg)
	}
	tracer := plugin.InitTracerContext()
	tx := tracer.StartTransaction(enums.ServicePluginMap[taskPlugin.ServicePlugin], "plugin")
	defer tx.End()
	err := plugin.Process(task, taskPlugin)
	if err != nil {
		tracer.SendError(err)
		return err
	}
	return nil
}

// ============================================================================================================
// stop
// ============================================================================================================

type taskPluginGroupStruct struct {
	task        models.Task
	taskPlugins []models.TaskPlugin
	wg          *sync.WaitGroup
	workerNo    int
}

type threadGroup struct {
	group       models.ThreadGroup
	isRunning   bool
	lock        sync.Mutex
	tasksChan   chan taskPluginGroupStruct
	quitChan    chan int
	workerGroup sync.WaitGroup
}

func (g *threadGroup) runProducer() {
	var fromTaskID dbUtils.SFID = 0
	var limitCount = 10
	var currentTime = time.Now()
	var sleepSeconds = time.Duration(g.group.SleepSeconds) * time.Second

	for g.isRunning {
		// 一个threadGroup下有多个协议
		// 一个协议一下有多个dataService组
		// 按目前的数据结构没法很好的按数量取dataService组, 所以保持按协议取
		taskIDs, err := services.ThreadGroupTaskService.QueryWaitProcessTaskIDList(context.TODO(), currentTime, g.group.ID, fromTaskID, limitCount)
		if err != nil {
			time.Sleep(sleepSeconds)
			fromTaskID = 0
			currentTime = time.Now()
			continue
		}
		for i := 0; i < len(taskIDs) && g.isRunning; i++ {
			taskID := taskIDs[i]
			if taskIDs[i] > fromTaskID {
				fromTaskID = taskID
			}
			ret, err := services.ThreadGroupTaskService.CompareAndSwapStatus(db.ClearingDB.NewConnection(context.TODO()), g.group.ID, taskID, enums.ThreadTaskStatusWaiting, enums.ThreadTaskStatusRunning)
			if err != nil {
				continue
			}
			if !ret {
				log.Infof("[threadGroup %s] compareAndSwapStatus failed. (%d, %d, %d)", g.group.Name, taskID, enums.ThreadTaskStatusWaiting, enums.ThreadTaskStatusRunning)
				continue
			}
			task, err := services.TaskService.GetTaskByID(context.TODO(), taskID)
			if err != nil {
				ret, _ = services.ThreadGroupTaskService.CompareAndSwapStatus(db.ClearingDB.NewConnection(context.TODO()), g.group.ID, taskID, enums.ThreadTaskStatusRunning, enums.ThreadTaskStatusWaiting)
				if !ret {
					log.Errorf("[threadGroup %s] compareAndSwapStatus failed. (%d, %d, %d)", g.group.Name, taskID, enums.ThreadTaskStatusWaiting, enums.ThreadTaskStatusRunning)
				}
				continue
			}
			now := time.Now()
			serviceList, err := services.TaskPluginService.GetTaskPluginListByTask(context.TODO(), taskID, enums.LifeCycleActive, &now)
			if err != nil {
				ret, _ = services.ThreadGroupTaskService.CompareAndSwapStatus(db.ClearingDB.NewConnection(context.TODO()), g.group.ID, taskID, enums.ThreadTaskStatusRunning, enums.ThreadTaskStatusWaiting)
				if !ret {
					log.Errorf("[threadGroup %s] compareAndSwapStatus failed. (%d, %d, %d)", g.group.Name, taskID, enums.ThreadTaskStatusWaiting, enums.ThreadTaskStatusRunning)
				}
				continue
			}
			var taskWg sync.WaitGroup
			taskWg.Add(1)
			g.tasksChan <- taskPluginGroupStruct{
				task:        *task,
				taskPlugins: serviceList,
				wg:          &taskWg,
			}
			go func(wg *sync.WaitGroup) {
				wg.Wait()
				ret, _ = services.ThreadGroupTaskService.CompareAndSwapStatus(db.ClearingDB.NewConnection(context.TODO()), g.group.ID, taskID, enums.ThreadTaskStatusRunning, enums.ThreadTaskStatusWaiting)
				if !ret {
					log.Errorf("[threadGroup %s] compareAndSwapStatus failed. (%d, %d, %d)", g.group.Name, taskID, enums.ThreadTaskStatusWaiting, enums.ThreadTaskStatusRunning)
				}
			}(&taskWg)

		}
		if len(taskIDs) == 0 {
			time.Sleep(sleepSeconds)
			fromTaskID = 0
			currentTime = time.Now()
			continue
		}
	}
}

func (g *threadGroup) NumberOfRunningTasks() int {
	sql := `select count(*) count from t_thread_group_task where thread_group_id=? and status=?`

	var count []int
	conn := db.ClearingDB.NewConnection(context.TODO())
	err := conn.Raw(sql, g.group.ID, enums.ThreadTaskStatusRunning).Pluck("count", &count).Error
	if err != nil || len(count) == 0 {
		return -1
	}
	return count[0]
}

func (g *threadGroup) HandlerPanicError() {
	if err := recover(); nil != err {
		e := tracer.Tracer.Recovered(err)
		e.Send()
		buf := make([]byte, 8192)
		n := runtime.Stack(buf, false)
		stackTraces := fmt.Sprintf("%s", buf[:n])
		logs.Error(fmt.Sprintf("[threadGroup %s] system panic: %+v", g.group.Name, stackTraces))
	}
}
