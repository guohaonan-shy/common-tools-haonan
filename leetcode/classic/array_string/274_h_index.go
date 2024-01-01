package array_string

import (
	"sort"
)

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	h_value := 0
	for i, v := range citations {
		value := i + 1
		if value <= v {
			h_value = value
		} else {
			break
		}
	}
	return h_value
}
