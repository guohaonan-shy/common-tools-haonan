package binary_search

func search(nums []int, target int) int {
	return binarySearchInRotateArray(nums, 0, len(nums)-1, target)
}

func binarySearchInRotateArray(nums []int, left, right int, target int) int {
	if left > right {
		return -1
	}

	mid := (left + right + 1) / 2

	if nums[mid] == target {
		return mid
	}

	// improved
	// 每次确认完mid之后，mid及mid以左或者mid及mid以右必有一个序列是有序的，之后判断有序队列的range是否能scope target
	if nums[left] < nums[mid] { // left is sorted
		if nums[left] <= target && target < nums[mid] { // 这里有个细节，如果目标值找不到的话，到最后，mid == target 或者 target == right，因此此处用mid比较，而非mid-1或者mid+1
			return binarySearchInRotateArray(nums, left, mid-1, target)
		} else {
			return binarySearchInRotateArray(nums, mid+1, right, target)
		}
	} else { // right is sorted
		if nums[mid] < target && target <= nums[right] {
			return binarySearchInRotateArray(nums, mid+1, right, target)
		} else {
			return binarySearchInRotateArray(nums, left, mid-1, target)
		}
	}

	//search1, search2 := binarySearchInRotateArray(nums, left, mid-1, target), binarySearchInRotateArray(nums, mid+1, right, target)
	//if search1 == -1 && search2 == -1 {
	//	return -1
	//}
	//
	//if search1 != -1 {
	//	return search1
	//} else {
	//	return search2
	//}
}
