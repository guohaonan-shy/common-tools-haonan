package param_check

import (
	"github.com/sirupsen/logrus"
	"reflect"
)

func IsParamValid(params ...interface{}) bool {
	res := true
	for _, param := range params {
		switch v := reflect.ValueOf(param); v.Kind() {
		case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			res = res && v.IsZero()
			if !res {
				logrus.Warnf("string or int param is invalid, param type is: %v", v.Kind())
				return res
			}
		case reflect.Ptr:
			res = res && v.IsNil()
			if !res {
				logrus.Warnf("ptr param is invalid!")
				return res
			}
		case reflect.Array, reflect.Slice, reflect.Chan, reflect.Map:
			res = res && !v.IsNil() && v.Len() != 0
			if !res {
				logrus.Warnf("Slice param is invalid!")
				return res
			}
		case reflect.Invalid:
			res = false
			if !res {
				logrus.Warnf("Invalid param is invalid!")
				return res
			}
		case reflect.Bool:
			res = res && v.Bool()
			return res
		default:
			logrus.Infof("unhandled kind %s", v.Kind())
		}
	}
	return res
}
