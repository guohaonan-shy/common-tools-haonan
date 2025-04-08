package array_string

func jump(nums []int) int {
	max_pos, end := 0, 0
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		max_pos = max(max_pos, i+nums[i])
		if i == end { // 目前以遍历到上一步能走的最远距离，需要跨到下一个最大距离
			end = max_pos
			step++
		}
	}
	return step
}

// pre-condition: all test cases are generated such that we can reach the final idx
func jumpV2(nums []int) int {
	maxPos, end := 0, 0
	cnt := 0

	for idx, step := range nums[:len(nums)-1] {
		maxPos = max(maxPos, idx+step)
		if idx == end {
			cnt++
			end = maxPos
		}
	}
	return cnt
}
