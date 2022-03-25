package models

import (
	"elder-wand/enums"
	"elder-wand/utils/dbUtils"
	"time"
)

type TaskPlugin struct {
	BaseModel
	TaskID        dbUtils.SFID          //任务ID
	ServicePlugin enums.ServicePlugin   //任务类型
	NextQueryTime *time.Time            //下次查询时间
	LastQueryTime time.Time             //最后一次查询时间
	Interval      int                   //间隔时间，单位秒
	LifeCycle     enums.CommonLifeCycle //生命周期
	Params        string                //配置参数
}

func (*TaskPlugin) TableName() string {
	return "t_task_plugin"
}
