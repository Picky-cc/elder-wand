package models

import (
	"elder-wand/enums"
	"elder-wand/utils/dbUtils"
	"time"
)

// 资金池任务表
type AgreementDataService struct {
	BaseModel
	AgreementID   dbUtils.SFID          //清算池ID
	ServicePlugin enums.ServicePlugin   //任务类型
	NextQueryTime *time.Time            //下次查询时间
	LastQueryTime time.Time             //最后一次查询时间
	Interval      int                   //间隔时间，单位秒
	LifeCycle     enums.CommonLifeCycle //生命周期
	Params        string                //配置参数
}

func (*AgreementDataService) TableName() string {
	return "t_clearing_agreement_data_service"
}
