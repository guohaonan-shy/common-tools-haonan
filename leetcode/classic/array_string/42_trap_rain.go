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

		// len(stack) > 0 => has top of stack; height[stack[len(stack)-1]] < height[i] => bottom exists
		/*
			one question: why must strictly increment here ?
			answer: not necessary, if we don't set strictly increment here, the stack is a strictly decreased; it's a question of left side or right side; no effect

			left => decrease; right increase
		*/
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// no storage if no decreased order on left side of ith
			if len(stack) == 0 {
				break
			}
			// the second smallest element in the decreased array
			left := stack[len(stack)-1]
			/*
				width = i-1-left
				height = min(height[left], height[i]) - height[top]
			*/
			water := (i - 1 - left) * (min(height[left], height[i]) - height[top])
			total += water
		}
		/*
			two scenarios for pushing stack:
			1. stack's length == 0
			2. decrease order; it means the top of stack is the local minimum of decreased order
		*/
		stack = append(stack, i)
	}
	return total
}
