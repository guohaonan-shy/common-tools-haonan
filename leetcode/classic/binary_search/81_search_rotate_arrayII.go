package binary_search

// 最坏情况，所有元素都是重复的，并且不是target元素 时间复杂度为O(N)
func searchII(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}
		// distinct array可以保证，mid左右两边区间至少又一个是单调的
		// 但是存在重复元素的情况无法保证这一情况，即存在三种情况：1. 左边单调 2. 右边单调 3. 左右两边都不单调
		if nums[left] == nums[mid] && nums[mid] == nums[right] { // 两边都找不到单调区间，只能逐步逼近
			left++
			right--
		} else if nums[mid] <= nums[right] {

			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		}
	}
	return false
}
