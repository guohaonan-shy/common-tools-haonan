package two_pointer

func nextPermutation(nums []int) {
	cur := len(nums) - 1
	for ; cur > 0; cur-- {
		if nums[cur-1] < nums[cur] {
			break
		}
	}

	if cur == 0 {
		reverse(nums)
		return
	}

	/*
		nums[cur:] is decreasing, we need to find the minimum which is greater than nums[cur-1]
	*/

	target := len(nums) - 1
	for ; target >= cur; target-- {
		if nums[target] > nums[cur-1] {
			break
		}
	}

	nums[cur-1], nums[target] = nums[target], nums[cur-1]

	reverse(nums[cur:])
	return
}

func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return
}
