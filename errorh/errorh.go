package errorh

import (
	"github.com/ghn980421/common-tools-haonan/json"
)

const (
	Errno_Success = 0

	Errno_Mysql_Conn_Failed   = 1000
	Errno_Mysql_Create_Failed = 1001
	Errno_Mysql_Select_Failed = 1002
	Errno_Mysql_Upsert_Failed = 1003
)

type CommonError struct {
	Errno  int64
	ErrMsg string

	err error
}

func (e *CommonError) Error() string {
	str, _ := json.StructBeautify(e)
	return str
}

func (e *CommonError) GetErrno() int64 {
	return e.Errno
}

func (e *CommonError) GetErrMsg() string {
	return e.ErrMsg
}

func WrapError(errno int64, msg string) error {
	return &CommonError{
		Errno:  errno,
		ErrMsg: msg,
	}
}
