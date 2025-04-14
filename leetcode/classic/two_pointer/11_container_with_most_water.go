package two_pointer

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max_area := 0
	for left < right {
		max_area = max(max_area, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			// 因为左右双指针无论谁移动，矩阵的长肯定是减小的，要想面积增大，唯一的可能就是改变宽，宽增大的唯一方法就是左右指针所指元素的较小值（宽）增加
			// 当height[right]大时，即容器盛水面积由较小的height[left]决定
			// 假如right--，那么面积一定是减小的(height[right-1]无论相比于height[right]是减小的还是增大的，宽仍是height[left] 或者 是比height[left]更小的值);
			// left++是有可能增加面积的
			// 从left和right 两个指针本身考虑
			left++
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

func maxAreaV2(height []int) int {
	area := 0
	left, right := 0, len(height)-1

	for left < right {

		width := right - left
		area = max(area, width*min(height[left], height[right]))
		// decide iterative direction of the pointer
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}
