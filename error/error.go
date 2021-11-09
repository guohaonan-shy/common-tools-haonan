package error

import (
	"github.com/ghn980421/common-tools-haonan/json"
)

type CommonError struct {
	Errno  int64
	ErrMsg string
}

func (e *CommonError) Error() string {
	return json.StructBeautify(e)
}

func Wrap(errno int64, msg string) error {
	return &CommonError{
		Errno:  errno,
		ErrMsg: msg,
	}
}
