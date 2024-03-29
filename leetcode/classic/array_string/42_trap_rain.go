package array_string

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

func trap_stack(height []int) int {
	stack := make([]int, 0)
	total := 0
	for i := 0; i < len(height); i++ {

		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]

			water := (i - 1 - left) * (min(height[left], height[i]) - height[top])
			total += water
		}
		stack = append(stack, i)
	}
	return total
}
