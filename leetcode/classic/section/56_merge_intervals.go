package section

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//preInterval := intervals[0]
	//res := make([][]int, 0)

	slow, fast := 0, 1
	for ; fast < len(intervals); fast++ {
		pre, cur := intervals[slow], intervals[fast]
		if pre[1] < cur[0] {
			slow++
			intervals[slow] = intervals[fast]
		} else {
			pre[1] = max(pre[1], cur[1])
			intervals[slow] = pre
		}

	}
	//for i := 1; i < len(intervals); i++ {
	//	interval := make([]int, 2)
	//	if preInterval[1] >= intervals[i][0] {
	//		interval[0] = preInterval[0]
	//		if preInterval[1] >= intervals[i][1] {
	//			interval[1] = preInterval[1]
	//		} else {
	//			interval[1] = intervals[i][1]
	//		}
	//		preInterval = interval
	//	} else {
	//		res = append(res, preInterval)
	//
	//		interval[0], interval[1] = intervals[i-1][0], intervals[i-1][1]
	//		preInterval = intervals[i]
	//	}
	//}
	//res = append(res, preInterval)
	return intervals[:slow+1]
}
