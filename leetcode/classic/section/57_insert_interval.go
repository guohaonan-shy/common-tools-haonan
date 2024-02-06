package section

func insert(intervals [][]int, newInterval []int) [][]int {
	// insert interval first

	pointer := 0
	for ; pointer < len(intervals); pointer++ {
		comp := intervals[pointer]
		if comp[0] > newInterval[0] {
			break
		}
	}

	// insert
	temp := make([][]int, len(intervals)-pointer)
	copy(temp, intervals[pointer:])
	intervals = append(append(intervals[:pointer], newInterval), temp...)

	// then merge interval
	slow, fast := 0, 1
	for ; fast < len(intervals); fast++ {
		pre, cur := intervals[slow], intervals[fast]
		if pre[1] < cur[0] {
			slow++
			intervals[slow] = intervals[fast]
		} else {
			intervals[slow][1] = max(intervals[slow][1], intervals[fast][1])
		}
	}
	return intervals[:slow+1]
}
