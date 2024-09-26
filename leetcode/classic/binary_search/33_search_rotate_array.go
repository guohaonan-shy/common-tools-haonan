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

func rotateSearch_noDistinct(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {

		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] <= nums[right] {

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
	return -1

}

func searchRotateArray(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {

		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// left side ascend
		if nums[left] < nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid
			} else {
				// if target is not in the left ascend sequence, it must be on the right side
				// because nums is made by rotating an ascending array at a specified index, if left sequence is ascending, the elements in right side must be greater than or smaller than left side.
				left = mid + 1
			}
		} else {
			if nums[mid+1] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	if nums[left] == target {
		return left
	}

	return -1
}
