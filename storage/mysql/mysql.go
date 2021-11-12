package mysql

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBWrapper struct {
	db *gorm.DB
	config *gorm.Config
}

// DBObjectInit
// This function only offers the basic option for db connection, if you want some specific options, please use session change setting
func DBObjectInit(dsn string) *DBWrapper {
	option := &gorm.Config{}
	DB, err := gorm.Open(sqlite.Open(dsn), option)
	if err != nil {
		logrus.Errorf("gorm.Open failed, err: %v", err)
		panic(err)
	}
	return &DBWrapper{
		db: DB,
		config: option,
	}
}

func (dw *DBWrapper)GetDBInstance() *gorm.DB {
	return dw.db
}