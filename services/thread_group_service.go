package services

import (
	"context"
	"elder-wand/db"
	"elder-wand/enums"
	"elder-wand/errors"
	"elder-wand/models"
	"github.com/jinzhu/gorm"
)

type threadGroupService struct {
}

var ThreadGroupService threadGroupService

func (t *threadGroupService) Create(conn *gorm.DB, threadGroup models.ThreadGroup) *errors.Error {
	err := conn.Create(&threadGroup).Error
	if err != nil {
		return errors.NewDBError(err.Error())
	}
	return nil
}

func (t *threadGroupService) GetActiveThreadGroups(conn *gorm.DB) ([]models.ThreadGroup, *errors.Error) {

	groups := make([]models.ThreadGroup, 0)
	err := conn.Raw("select * from t_thread_group where life_cycle=? ", enums.ThreadGroupLifeCycleActive).Scan(&groups).Error
	if err != nil {
		return groups, errors.NewDBError(err.Error())
	}
	return groups, nil
}

func (t *threadGroupService) GetActiveThreadGroupsByNodeID(ctx context.Context, nodeID int) ([]models.ThreadGroup, *errors.Error) {

	conn := db.ClearingDB.NewConnection(ctx)
	groups := make([]models.ThreadGroup, 0)
	err := conn.Raw("select * from t_thread_group where ew_node=? and life_cycle=? ", nodeID, enums.ThreadGroupLifeCycleActive).Scan(&groups).Error
	if err != nil {
		return groups, errors.NewDBError(err.Error())
	}
	return groups, nil
}
