package binary_search

func searchRange(nums []int, target int) []int {
	left := binarySearch_Range(nums, 0, len(nums)-1, target, false)
	if left == -1 {
		return []int{-1, -1}
	}
	return []int{left, binarySearch_Range(nums, 0, len(nums)-1, target, true)}
}

func binarySearch_Range(nums []int, left, right int, target int, last bool) int {

	for left <= right {
		mid := (left + right + 1) / 2
		if nums[mid] == target {
			if last {
				if mid == len(nums)-1 || nums[mid+1] != target {
					return mid
				} else {
					left = mid + 1
				}
			} else {
				if mid == 0 || nums[mid-1] != target {
					return mid
				} else {
					right = mid - 1
				}
			}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func searchRangeV2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	return []int{findFirst(nums, target), findLast(nums, target)}
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if nums[left] != target {
		return -1
	}

	return left
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right + 1) / 2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}

	if nums[left] != target {
		return -1
	}
	return left
}
