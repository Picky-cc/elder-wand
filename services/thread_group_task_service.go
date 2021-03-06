package services

import (
	"context"
	"elder-wand/db"
	"elder-wand/enums"
	"elder-wand/errors"
	"elder-wand/models"
	"elder-wand/utils/dbUtils"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type threadGroupTaskService struct {
}

var ThreadGroupTaskService threadGroupTaskService

func (t *threadGroupTaskService) Create(conn *gorm.DB, threadGroupTask models.ThreadGroupTask) *errors.Error {
	err := conn.Create(&threadGroupTask).Error
	if err != nil {
		return errors.NewDBError(err.Error())
	}
	return nil
}

func (t *threadGroupTaskService) CompareAndSwapStatus(conn *gorm.DB, threadGroupID, taskID dbUtils.SFID, fromStatus, toStatus enums.ThreadGroupTaskStatus) (bool, *errors.Error) {
	sql := `
		update t_thread_group_task set status=?, updated=now() where thread_group_id=? and task_id=? and status=?
	`
	conn = conn.Exec(sql, toStatus, threadGroupID, taskID, fromStatus)
	if conn.Error != nil {
		return false, errors.NewDBError(conn.Error.Error())
	}
	return conn.RowsAffected == 1, nil
}

func (t *threadGroupTaskService) ResetThreadGroupTaskStatus(conn *gorm.DB, ewNodeID int) (int64, *errors.Error) {
	sql := `
	update t_thread_group_task set status=?, updated=now() where status=? and thread_group_id in (select id from t_clearing_thread_group where ew_node=?)
`
	conn = conn.Exec(sql, enums.ThreadTaskStatusWaiting, enums.ThreadTaskStatusRunning, fmt.Sprint(ewNodeID))
	if conn.Error != nil {
		return 0, errors.NewDBError(conn.Error.Error())
	}
	return conn.RowsAffected, nil
}

func (t *threadGroupTaskService) QueryWaitProcessTaskIDList(ctx context.Context, tt time.Time, threadGroupID, fromTaskID dbUtils.SFID, limitCount int) ([]dbUtils.SFID, *errors.Error) {

	sql := `
select
	a.task_id
from
	t_thread_group_task as a, t_thread_group as c
where
	a.thread_group_id=c.id and a.thread_group_id=? and a.status=1 and a.task_id in (
	select
		b.id
	from
		t_task as b
	inner join
		t_task_plugin as c on b.id=c.task_id
	where
		b.id > ? and b.life_cycle=? and (c.next_query_time is null or c.next_query_time <= ?) and c.life_cycle=?
) and c.life_cycle=?
order by
	a.task_id asc
limit ?

`
	params := make([]interface{}, 0)
	params = append(params, threadGroupID, fromTaskID, enums.TaskLifeCycleActive, tt, enums.LifeCycleActive, enums.ThreadGroupLifeCycleActive, limitCount)

	results := make([]dbUtils.SFID, 0)
	conn := db.EwDB.NewConnection(ctx)
	err := conn.Raw(sql, params...).Pluck("task_id", &results).Error
	if err != nil {
		return results, errors.NewDBError(err.Error())
	}
	return results, nil
}
