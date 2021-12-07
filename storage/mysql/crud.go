package mysql

import (
	"fmt"
	"github.com/ghn980421/common-tools-haonan/errorh"
	"gorm.io/gorm/clause"
)

func (dw *DBWrapper) Create(row interface{}) error {
	db := dw.GetDBInstance()

	if db.Error != nil {
		return errorh.WrapError(errorh.Errno_Mysql_Conn_Failed, fmt.Sprintf("GetDBInstance failed, err: %v", db.Error))
	}

	if err := db.Model(row).Create(row); err != nil {
		return errorh.WrapError(errorh.Errno_Mysql_Create_Failed, fmt.Sprintf("Create failed, err: %v", err))
	}
	return nil
}

func (dw *DBWrapper) Update(model interface{}, update map[string]interface{}) error  {
	db := dw.GetDBInstance().Model(model)

	if db.Error != nil {
		return errorh.WrapError(errorh.Errno_Mysql_Conn_Failed, fmt.Sprintf("GetDBInstance failed, err: %v", db.Error))
	}

	if err := db.Model(model).Updates(update); err != nil {
		return errorh.WrapError(errorh.Errno_Mysql_Update_Failed, fmt.Sprintf("Update failed, err: %v", err))
	}

	return nil
}
