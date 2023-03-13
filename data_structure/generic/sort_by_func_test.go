package generic

import (
	"encoding/json"
	"testing"
)

type BizType int32

const (
	BizType_Unknown BizType = 0
	BizType_A       BizType = 1
	BizType_B       BizType = 2
	BizType_C       BizType = 3
	BizType_D       BizType = 4
	BizType_E       BizType = 5
	BizType_NB      BizType = 99
)

type ComparableInstance struct {
	BizType BizType `json:"biz_type"`
}

func NewComparableInstance(biz BizType) *ComparableInstance {
	return &ComparableInstance{
		BizType: biz,
	}
}

func SortFn(a, b *ComparableInstance) bool {
	if b.BizType == BizType_NB {
		return true
	}

	return a.BizType >= b.BizType
}

func Test_SortByFunc(t *testing.T) {
	// int64, string, struct that all fields are comparable
	sortList := make([]*ComparableInstance, 0)

	for i := 0; i < 6; i++ {
		instance := &ComparableInstance{}
		if i == 5 {
			instance.BizType = BizType_NB
		} else {
			instance.BizType = BizType(i + 1)
		}
		sortList = append(sortList, instance)
	}

	bytesIn, _ := json.Marshal(sortList)
	t.Logf("Before Sort: %s", bytesIn)

	res := SortByFunc[*ComparableInstance](sortList, SortFn)

	bytesOut, _ := json.Marshal(res)
	t.Logf("After Sort: %s", bytesOut)

}
