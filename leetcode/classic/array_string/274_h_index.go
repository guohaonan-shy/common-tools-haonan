package array_string

import (
	"sort"
)

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	h_value := 0
	for i, v := range citations {
		value := i + 1
		if value <= v { // at least i+1 paper's citation is more than i+1, because i+1 < current citation
			h_value = value
		} else {
			break
		}
	}
	return h_value
}
