package generic

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ContainInt64(t *testing.T) {
	res := Contains[int64](int64(1), []int64{1, 2, 3})
	res2 := Contains[int64](int64(1123121), []int64{1, 2, 3})
	assert.Equal(t, true, res)
	assert.Equal(t, false, res2)
}

func Test_ContainString(t *testing.T) {
	res := Contains("abc", []string{"abc", "def", "efs"})
	res2 := Contains("a", []string{"b", "c", "d"})
	assert.Equal(t, true, res)
	assert.Equal(t, false, res2)
}

func Test_ContainByFunc(t *testing.T) {
	// int64, string, struct that all fields are comparable
	fn := func(value int64) bool {
		if value == 1 {
			return true
		}
		return false
	}

	res1 := ContainsByFunc(fn, []int64{1, 2, 3})
	res2 := ContainsByFunc(fn, []int64{2, 3, 4})
	assert.Equal(t, true, res1, "should contain 1, list:%d", []int64{2, 2, 3})
	assert.Equal(t, false, res2, "shouldn't contain 1")

	// struct

	type testStruct struct {
		value string
		age   int64
	}

	s := []*testStruct{
		{
			value: "yes",
			age:   10,
		},
		{
			value: "no",
			age:   21,
		},
		{
			value: "ghn",
			age:   32,
		},
	}

	fn1 := func(object *testStruct) bool {
		if object.value == "ghn" {
			return true
		}
		return false
	}

	res3 := ContainsByFunc(fn1, s)
	assert.Equal(t, true, res3, "should contain ghn")
}

func Test_FilterMap(t *testing.T) {

	type BizModelEntity struct {
		EntityId     int64 `json:"entity_id"`
		IsNeedFilter bool  `json:"is_need_filter"`
	}

	in := map[int64]*BizModelEntity{
		123: &BizModelEntity{EntityId: 123, IsNeedFilter: true},
		234: &BizModelEntity{EntityId: 234, IsNeedFilter: false},
		345: &BizModelEntity{EntityId: 345, IsNeedFilter: false},
		456: &BizModelEntity{EntityId: 456, IsNeedFilter: true},
		457: &BizModelEntity{EntityId: 457, IsNeedFilter: true},
	}

	fn := func(value *BizModelEntity) bool { // 业务函数
		if value.IsNeedFilter == true {
			return true
		}
		return false
	}

	bytesIn, _ := json.Marshal(in)
	t.Logf("Before Filter: %s", bytesIn)

	res := FilterByFunc[int64, *BizModelEntity](in, fn)

	bytesOut, _ := json.Marshal(res)
	t.Logf("After Filter: %s", bytesOut)

	assert.Equal(t, map[int64]*BizModelEntity{
		234: &BizModelEntity{EntityId: 234, IsNeedFilter: false},
		345: &BizModelEntity{EntityId: 345, IsNeedFilter: false},
	}, res)
}
