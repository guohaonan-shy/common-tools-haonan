package binary_search

// [1,3,5,6] 2
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] == target {
		return left
	} else if nums[left] > target {
		return left
	} else {
		return left + 1
	}
}
