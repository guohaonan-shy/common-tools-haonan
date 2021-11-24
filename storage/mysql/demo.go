package mysql

import (
	"fmt"
	"github.com/ghn980421/common-tools-haonan/errorh"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	DB *DBWrapper
)

type MiniDemo struct {
	a int64
	b int64
	c int64
}

func (m *MiniDemo) BeforeCreate(tx *gorm.DB) error {
	conditions := &MiniDemo{
		a: 1,
		b: 2,
	}

	des := make([]*MiniDemo, 0)
	err := tx.Where(conditions).Find(des).Error

	if err != nil {
		logrus.Errorf("after execute failed, error: %s", err)
		return errorh.WrapError(errorh.Errno_Mysql_Select_Failed, fmt.Sprintf("select failed, err:%s", err))
	}

	// business logic
	//
	//

	return nil
}

// HookUse gorm begin transaction firstly, and execute before hook func, then execute create/save/update/delete and after hook func
// Because these atomic operations execute by default transaction, if any error occur in the process, the transaction will rollback
func HookUse() error {
	DB := DBObjectInit("")

	m := &MiniDemo{}

	// GORM perform single create, update, delete operations in transactions by default to ensure database data integrity
	// You can disable it by setting `SkipDefaultTransaction` to true
	err := DB.GetDBInstance().Create(m).Error
	if err != nil {
		logrus.Errorf("before execute failed, error: %s", err)
		return err
	}

	return nil
}

func TransactionUse() error {
	DB := DBObjectInit("")

	m := &MiniDemo{}

	upsertFunc := func(tx *gorm.DB) error {
		err := DB.GetDBInstance().Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "a"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"status":1}),
		}).Create(m).Error
		
		if err != nil {
			return errorh.WrapError(errorh.Errno_Mysql_Upsert_Failed, fmt.Sprintf("upsert failed, err: %s", err))	
		}

		return nil
	}
	
	err := DB.GetDBInstance().Transaction(upsertFunc)
	
	if err != nil {
		return err
	}

	return nil
}