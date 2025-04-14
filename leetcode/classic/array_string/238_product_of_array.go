package array_string

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))

	for i := range nums {
		leftVal := 1
		if i > 0 {
			leftVal = left[i-1] * nums[i-1]
		}

		left[i] = leftVal
	}

	for i := len(nums) - 1; i >= 0; i-- {
		rightVal := 1
		if i < len(nums)-1 {
			rightVal = right[i+1] * nums[i+1]
		}

		right[i] = rightVal
	}

	res := make([]int, len(nums))
	for i := range res {
		res[i] = left[i] * right[i]
	}
	return res
}
