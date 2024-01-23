package binary_search

func findMin(nums []int) int {
	return binarySearchMinRotate(nums, 0, len(nums)-1)
}

func binarySearchMinRotate(nums []int, left, right int) int {

	for left < right { // 至少有两个元素
		// 单调
		//if nums[left] <= nums[right] {
		//	return nums[left]
		//}

		mid := (left + right) / 2 // 靠左倾斜
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
		// 按照靠左倾斜计算，则会出现当left = right -1时，会出现自己比较自己
		//if nums[mid] >= nums[left] {
		//	left = mid + 1
		//} else {
		//	right = mid
		//}

	}

	return nums[(left+right)/2]
}
