package json

import (
	"bytes"
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/decoder"
	"github.com/sirupsen/logrus"
	"reflect"
)

// StructBeautify ident handle
func StructBeautify(i interface{}) (string, error) {
	bs, err := sonic.Marshal(i)
	if err != nil {
		logrus.Errorf("sonic.Marshal failed, err: %v", err)
		return "", err
	}
	buf := new(bytes.Buffer)
	err = json.Indent(buf, bs, "", "\t")
	if err != nil {
		logrus.Errorf("json.Indent failed, err: %v", err)
		return "", err
	}
	return buf.String(), nil
}

func BytesJsonDecode(data string, obj interface{}) error {
	decodeObj := decoder.NewDecoder(data)
	decodeObj.UseNumber()
	return decodeObj.Decode(&obj)
}

func GetJSONString(object interface{}) string {
	if object == nil {
		logrus.Warnf("convert json string nil object")
		return ""
	}
	if v := reflect.ValueOf(object); v.Kind() == reflect.Ptr {
		if v.IsNil() {
			logrus.Warnf("convert json string nil ptr")
			return ""
		}
	}
	jsonBytes, err := sonic.Marshal(object)
	if err != nil {
		logrus.Error("convert json string err=%v", err)
		return "{}" //default empty json object string
	}
	return string(jsonBytes)
}
