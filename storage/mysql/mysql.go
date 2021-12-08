package mysql

import (
	"gorm.io/gorm"
)

type DBWrapper struct {
	handler *DBHandler
	config  *gorm.Config
}

// DBObjectInit
// This function only offers the basic option for db connection, if you want some specific options, please use session change setting
func DBObjectInit(dsn string) *DBWrapper {

	option := &gorm.Config{
		SkipDefaultTransaction: false,
	}

	dbHandler := &DBHandler{}
	dbHandler.ConnectDB(dsn, option)

	if dbHandler.err != nil {
		panic(dbHandler.err)
	}

	return &DBWrapper{
		handler: dbHandler,
		config:  option,
	}
}

func (dw *DBWrapper) GetDBHandler() *DBHandler {
	return dw.handler
}

func (dw *DBWrapper) SetDBHandler(handler *DBHandler) {
	dw.handler = handler
}
