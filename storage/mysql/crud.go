package mysql

import (
	"fmt"
	"github.com/ghn980421/common-tools-haonan/errorh"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (dw *DBWrapper) GetConn(model interface{}) (*gorm.DB, error) {
	db := dw.GetDBInstance().Model(model)

	if db.Error != nil {
		return nil, errorh.WrapError(errorh.Errno_Mysql_Conn_Failed, fmt.Sprintf("GetDBInstance failed, err: %v", db.Error))
	}
	return db, nil
}

func (dw *DBWrapper) Create(row interface{}) error {
	db, err := dw.GetConn(row)
	if err != nil {
		return err
	}

	if err = db.Create(row).Error; err != nil {
		return errorh.WrapError(errorh.Errno_Mysql_Create_Failed, fmt.Sprintf("Create failed, err: %v", err))
	}
	return nil
}

func (dw *DBWrapper) Update(model interface{}, update map[string]interface{}) error  {
	db, err := dw.GetConn(model)
	if err != nil {
		return err
	}

	if err := db.Updates(update).Error; err != nil {
		return errorh.WrapError(errorh.Errno_Mysql_Update_Failed, fmt.Sprintf("Update failed, err: %v", err))
	}

	return nil
}

func (dw *DBWrapper) Upsert(create interface{}, update map[string]interface{}, ignore bool) error {

	db, err := dw.GetConn(create)
	if err != nil {
		return err
	}

	conflict := clause.OnConflict{
		DoNothing: ignore,
		DoUpdates: clause.Assignments(update),
	} // conflict handle option

	if err = db.Clauses(conflict).Create(create).Error; err != nil {
		return errorh.WrapError(errorh.Errno_Mysql_Upsert_Failed, fmt.Sprintf("Upsert failed, err: %v", err))
	}

	return nil
}
