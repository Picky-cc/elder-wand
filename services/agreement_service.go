package services

import (
	"context"
	"elder-wand/db"
	"elder-wand/enums"
	"elder-wand/errors"
	"elder-wand/models"
	"elder-wand/utils/dbUtils"
	"github.com/jinzhu/gorm"
)

type agreementService struct {
}

var AgreementService agreementService

func (a *agreementService) Create(conn *gorm.DB, agreement models.Agreement) *errors.Error {
	err := conn.Create(&agreement).Error
	if err != nil {
		return errors.NewDBError(err.Error())
	}
	return nil
}

func (a *agreementService) GetAgreementByProjectAndAgreementType(ctx context.Context, projectID dbUtils.SFID, aType enums.AgreementType) (*models.Agreement, *errors.Error) {
	agreement := models.Agreement{}
	conn := db.ClearingDB.NewConnection(ctx)

	sql := ` select * from t_clearing_agreement where clearing_project_id = ? and agreement_type = ? `
	conn = conn.Raw(sql, projectID, aType).Find(&agreement)
	if conn.RecordNotFound() {
		return nil, nil
	} else if conn.Error != nil {
		return nil, errors.NewDBError(conn.Error.Error())
	}
	return &agreement, nil
}

func (a *agreementService) GetAgreementByID(ctx context.Context, ID dbUtils.SFID) (*models.Agreement, *errors.Error) {
	agreement := models.Agreement{}
	conn := db.ClearingDB.NewConnection(ctx)

	sql := ` select * from t_clearing_agreement where id=? `
	conn = conn.Raw(sql, ID).Find(&agreement)
	if conn.RecordNotFound() {
		return nil, nil
	} else if conn.Error != nil {
		return nil, errors.NewDBError(conn.Error.Error())
	}
	return &agreement, nil
}

func (a *agreementService) GetActiveAgreementByID(ctx context.Context, ID dbUtils.SFID, lifeCycle enums.AgreementLifeCycle) (*models.Agreement, *errors.Error) {
	agreement := models.Agreement{}
	conn := db.ClearingDB.NewConnection(ctx)

	sql := ` select * from t_clearing_agreement where id=? and life_cycle = ? `
	conn = conn.Raw(sql, ID, lifeCycle).Find(&agreement)
	if conn.RecordNotFound() {
		return nil, errors.NewNotFoundErrorf("Agreement %d not exist", ID)
	} else if conn.Error != nil {
		return nil, errors.NewDBError(conn.Error.Error())
	}
	return &agreement, nil
}

func (a *agreementService) GetActiveAgreement(conn *gorm.DB) ([]models.Agreement, *errors.Error) {
	list := make([]models.Agreement, 0)

	sql := ` select * from t_clearing_agreement where life_cycle = ? `
	err := conn.Raw(sql, enums.AgreementLifeCycleActive).Find(&list).Error
	if err != nil {
		return list, errors.NewDBError(conn.Error.Error())
	}
	return list, nil
}
