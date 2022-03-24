package services

import (
	"context"
	"elder-wand/db"
	"elder-wand/enums"
	"elder-wand/errors"
	"elder-wand/models"
	"elder-wand/utils/dbUtils"
	"github.com/jinzhu/gorm"
	"time"
)

type agreementDataService struct {
}

var AgreementDataService agreementDataService

func (a *agreementDataService) Create(conn *gorm.DB, agreementDataService models.AgreementDataService) *errors.Error {
	err := conn.Create(&agreementDataService).Error
	if err != nil {
		return errors.NewDBError(err.Error())
	}
	return nil
}

func (a *agreementDataService) GetDataServiceListByAgreement(ctx context.Context, agreementID dbUtils.SFID, lifeCycle enums.CommonLifeCycle, queryTime *time.Time) ([]models.AgreementDataService, *errors.Error) {
	results := make([]models.AgreementDataService, 0)

	sql := "select * from t_clearing_agreement_data_service where agreement_id=? and life_cycle=? and next_query_time <= ? "

	conn := db.ClearingDB.NewConnection(ctx)
	err := conn.Raw(sql, agreementID, lifeCycle, queryTime).Scan(&results).Error
	if err != nil {
		return nil, errors.NewDBError(err.Error())
	}
	return results, nil
}

func (a *agreementDataService) UpdateDataServiceQueryTime(conn *gorm.DB, serviceID dbUtils.SFID, queryTime time.Time) (int64, *errors.Error) {
	sql := "update t_clearing_agreement_data_service set last_query_time=?, next_query_time=date_add(?, interval `interval` second) where id=?"
	conn = conn.Exec(sql, queryTime, queryTime, serviceID)
	if conn.Error != nil {
		return 0, errors.NewDBError(conn.Error.Error())
	}
	return conn.RowsAffected, nil
}
