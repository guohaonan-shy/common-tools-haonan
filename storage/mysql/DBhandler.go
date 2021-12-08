package mysql

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBHandler struct {
	db       *gorm.DB
	err      error
	connFunc func() (*gorm.DB, error)
}

func (handler *DBHandler) ConnectDB(dsn string, option *gorm.Config) {
	conn, err := gorm.Open(sqlite.Open(dsn), option)
	if err != nil {
		logrus.Errorf("gorm.Open failed, err: %v", err)
	}

	handler.db = conn
	handler.err = err
	return
}

func (handler *DBHandler) GetConn() (*gorm.DB, error) {
	return handler.connFunc()
}
