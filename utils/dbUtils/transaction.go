package dbUtils

import (
	"elder-wand/errors"

	"github.com/jinzhu/gorm"
)

func RunInTransaction(conn *gorm.DB, fn func(*gorm.DB) *errors.Error) *errors.Error {
	tx := conn.Begin()
	if tx.Error != nil {
		return errors.NewDBError(tx.Error.Error())
	}
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return errors.NewDBError(err.Error())
	}
	return nil
}
