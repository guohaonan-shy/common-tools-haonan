package section

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] <= points[j][0]
	})

	//left := 0
	conjunction := 0
	preBallon := points[0]
	shoot := 1
	for right := 1; right < len(points); right++ {
		interval := make([]int, 2)

		if preBallon[1] < points[right][0] {
			interval[0] = preBallon[0]
			interval[1] = interval[0] - 1
		} else {
			interval[0] = max(preBallon[0], points[right][0])
			interval[1] = min(preBallon[1], points[right][1])
		}

		conjunction = interval[1] - interval[0] + 1
		if conjunction == 0 {
			shoot++
			preBallon = points[right]
		} else {
			preBallon = interval
		}
	}
	return shoot
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
