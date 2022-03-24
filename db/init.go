package db

import (
	"elder-wand/enums"
	"elder-wand/errors"
	"elder-wand/models"
	"elder-wand/settings"
	"elder-wand/utils/dbUtils"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

func initDefaultThreadGroup(db *gorm.DB) {

	threadGroupList, err := getActiveThreadGroups(db)
	if err != nil {
		panic(err)
	}
	if len(threadGroupList) > 0 {
		return
	}
	trans := db.Begin()
	threadGroup := models.ThreadGroup{
		BaseModel: models.BaseModel{
			ID:      dbUtils.GenerateID(),
			Created: time.Now(),
			Updated: time.Now(),
		},
	}
	threadGroup.Name = "默认线程组"
	threadGroup.ThreadCount = 10
	threadGroup.SleepSeconds = 20
	threadGroup.LifeCycle = enums.ThreadGroupLifeCycleActive
	threadGroup.EwNode = fmt.Sprint(settings.Config.EwNodeID)

	err = createThreadGroup(trans, threadGroup)
	if err != nil {
		trans.Rollback()
		panic(err)
	}
	agreements, err := getActiveAgreement(trans)
	if err != nil {
		trans.Rollback()
		panic(err)
	}
	for _, agreement := range agreements {
		group := models.ThreadGroupTask{
			BaseModel: models.BaseModel{
				ID:      dbUtils.GenerateID(),
				Created: time.Now(),
				Updated: time.Now(),
			},
			ThreadGroupID: threadGroup.ID,
			AgreementID:   agreement.ID,
			Status:        enums.ThreadTaskStatusWaiting,
		}

		err = createThreadGroupTask(trans, group)
		if err != nil {
			trans.Rollback()
			panic(err)
		}
	}

	trans.Commit()
}

func getActiveThreadGroups(conn *gorm.DB) ([]models.ThreadGroup, *errors.Error) {

	groups := make([]models.ThreadGroup, 0)
	err := conn.Raw("select * from t_clearing_thread_group where life_cycle=? ", enums.ThreadGroupLifeCycleActive).Scan(&groups).Error
	if err != nil {
		return groups, errors.NewDBError(err.Error())
	}
	return groups, nil
}

func createThreadGroup(conn *gorm.DB, threadGroup models.ThreadGroup) *errors.Error {
	err := conn.Create(&threadGroup).Error
	if err != nil {
		return errors.NewDBError(err.Error())
	}
	return nil
}

func getActiveAgreement(conn *gorm.DB) ([]models.Agreement, *errors.Error) {
	list := make([]models.Agreement, 0)

	sql := ` select * from t_clearing_agreement where life_cycle = ? `
	err := conn.Raw(sql, enums.AgreementLifeCycleActive).Scan(&list).Error
	if err != nil {
		return list, errors.NewDBError(conn.Error.Error())
	}
	return list, nil
}

func createThreadGroupTask(conn *gorm.DB, threadGroupTask models.ThreadGroupTask) *errors.Error {
	err := conn.Create(&threadGroupTask).Error
	if err != nil {
		return errors.NewDBError(err.Error())
	}
	return nil
}
