package json

import huge "github.com/dablelv/go-huge-util"

// StructBeautify ident handle
func StructBeautify(i interface{}) string {
	s, _ := huge.ToIndentJSON(i)
	return s
}
