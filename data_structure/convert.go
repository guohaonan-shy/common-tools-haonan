package data_structure

import (
	"context"
	"fmt"
	"github.com/ghn980421/common-tools-haonan/errorh"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"reflect"
)

type Converter interface {
	Convert(ctx context.Context, a interface{}, b interface{}) error
}

func CommonConvert(ctx context.Context, input interface{}, output interface{}) error {
	err := mapstructure.Decode(input, output)
	if err != nil {
		logrus.Errorf("Decode Error:%s, Please Check Your Input:%v", err, input)
		return errorh.WrapError(errorh.Errno_MapStructure_Decode_failed, fmt.Sprintf("Decode Failed, err:%s", err ))
	}

	// execute converter.convert
	if ok := reflect.TypeOf(output).Implements(reflect.TypeOf((*Converter)(nil)).Elem()); !ok {
		return nil
	}

	err = output.(Converter).Convert(ctx, input, output)
	if err != nil {
		logrus.Errorf("Implement func execute failed, err:%s", err)
		return errorh.WrapError(errorh.Error_Interface_Implement_Execute_failed, fmt.Sprintf("Implement func execute failed, err:%s", err))
	}

	return nil
}
