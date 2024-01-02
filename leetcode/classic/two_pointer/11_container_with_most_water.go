package two_pointer

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max_area := 0
	for left < right {
		max_area = max(max_area, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] { // 因为左右双指针无论谁移动，矩阵的长肯定是减小的，要想面积增大，唯一的可能就是改变宽，宽增大的唯一方法就是左右指针所指元素的较小值（宽）增加
			left++ // 假如right--，那么面积一定是减小的(height[right-1]如果仍然是偏大的，即宽仍是height[left]; 反之宽是比height[left]更小的值);left++面积可能会增加
		} else { // 反之
			right--
		}
	}
	return max_area
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
