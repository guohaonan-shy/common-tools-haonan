package classic

// 需要比较左右两边元素，因此左右两次遍历
func trap(height []int) int {
	left_max := height[0]
	left := make([]int, len(height))
	for i := 1; i < len(height); i++ {
		if left_max > height[i] {
			left[i] = left_max - height[i]
		} else {
			left_max = max(left_max, height[i])
		}
	}

	right_max := height[len(height)-1]
	right := make([]int, len(height))
	for i := len(height) - 2; i >= 0; i-- {
		if right_max > height[i] {
			right[i] = right_max - height[i]
		} else {
			right_max = max(right_max, height[i])
		}
	}

	total := 0
	for i := 0; i < len(height); i++ {
		total += min(left[i], right[i])
	}
	return total
}
