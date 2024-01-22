package binary_search

func findMin(nums []int) int {
	return binarySearchMinRotate(nums, 0, len(nums)-1)
}

func binarySearchMinRotate(nums []int, left, right int) int {

	for left < right {
		// 单调
		//if nums[left] <= nums[right] {
		//	return nums[left]
		//}

		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}

	}

	return nums[left]
}
