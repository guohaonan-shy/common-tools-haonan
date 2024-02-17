package stack

func largestRectangleArea(heights []int) int {
	idx := maxIdx(heights)

	maxRec := 0
	for _, id := range idx {
		left, right := id, id
		length := heights[id]
		for {
			var leftVal, rightVal int
			if left > 0 {
				leftVal = heights[left-1]
			}
			if right < len(heights)-1 {
				rightVal = heights[right+1]
			}

			if leftVal < rightVal {
				if right < len(heights)-1 {
					right++
				}
				maxRec = max(min(length, rightVal)*(right-left+1), maxRec)
				length = min(length, rightVal)
			} else if leftVal > rightVal {
				if left > 0 {
					left--
				}
				maxRec = max(min(length, leftVal)*(right-left+1), maxRec)
				length = min(length, leftVal)
			} else {
				if left > 0 {
					left--
				}
				if right < len(heights)-1 {
					right++
				}
				maxRec = max(min(length, leftVal)*(right-left+1), maxRec)
				length = min(length, leftVal)
			}

			if left <= 0 && right >= len(heights)-1 {
				break
			}
		}

	}
	return maxRec
}

func maxIdx(nums []int) []int {
	maxVal := -1
	idx := make([]int, 0)
	for i, val := range nums {

		if val > maxVal {
			idx = []int{i}
			maxVal = val
			continue
		}

		if val == maxVal {
			idx = append(idx, i)
		}

	}
	return idx
}

// 要想复杂度为N，大概思路就是，height[i]为长的最大面积从左到右进行遍历
// 如果要以height[i]为长，宽就是左右两边距离i最近的小于height[i]的元素为边界
func largestRectangleArea_standard(heights []int) int {

	// 首先先从左往右遍历
	stack := make([]int, 0)
	left := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		// pivot := height[i]
		// 从栈顶索引指向的元素开始往左，栈顶索引指向的元素为height[i-1]，那么stack[-1]和stack[-2]之间的索引指向的值一定是大height[stack[-1]]的
		// 即该栈实际指向的元素是从左往右的递增序列
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0)
	right := make([]int, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			right[i] = len(heights)
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	maxVal := -1
	for i := 0; i < len(heights); i++ {
		maxVal = max(maxVal, (right[i]-left[i]-1)*heights[i])
	}
	return maxVal
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
