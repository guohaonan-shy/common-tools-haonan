package array_string

// 当某个位置x能抵达，意味着在x之前的位置，至少存在一个y，y可达且y+nums[y] >= x
func canJump(nums []int) bool {
	max_distance := 0
	for pos, distance := range nums {
		if pos <= max_distance { // can reach
			max_distance = max(max_distance, pos+distance)
		} else {
			break
		}
	}
	return max_distance >= len(nums)-1
}

func canJumpV2(nums []int) bool {
	maxPos := 0

	for idx, step := range nums {
		// current position 'idx' is in the available area
		if idx <= maxPos {
			maxPos = max(maxPos, idx+step)
		} else {
			break
		}
	}
	return maxPos >= len(nums)-1
}
