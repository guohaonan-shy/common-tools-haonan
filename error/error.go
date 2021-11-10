package error

import (
	"github.com/ghn980421/common-tools-haonan/json"
	"github.com/sirupsen/logrus"
)

type CommonError struct {
	Errno  int64
	ErrMsg string
}

func (e *CommonError) Error() string {
	str, err := json.StructBeautify(e)
	if err != nil {
		logrus.Errorf("error failed, err: %v", err)
		return ""
	}
	return str
}

func Wrap(errno int64, msg string) error {
	return &CommonError{
		Errno:  errno,
		ErrMsg: msg,
	}
}
