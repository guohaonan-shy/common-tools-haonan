package slice

import "reflect"

func Contains(item interface{}, array interface{}) bool {
	if reflect.TypeOf(array).Kind() == reflect.Slice || reflect.TypeOf(array).Kind() == reflect.Array {
		arrayValue := reflect.ValueOf(array)
		for i := 0; i < arrayValue.Len(); i++ {
			if reflect.DeepEqual(item, arrayValue.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

