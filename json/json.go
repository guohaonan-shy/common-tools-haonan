package json

import (
	"bytes"
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/decoder"
	"github.com/sirupsen/logrus"
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
