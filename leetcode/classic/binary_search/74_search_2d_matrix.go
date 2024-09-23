package binary_search

func searchMatrix(matrix [][]int, target int) bool {
	columnNums := len(matrix[0])
	for _, row := range matrix {
		idx := binarySearch(row, 0, columnNums-1, target)
		if idx < columnNums {
			if row[idx] == target {
				return true
			} else {
				return false
			}
		}
	}
	return false

}

func binarySearch(nums []int, left, right int, target int) int {
	if left > right {
		return left
	}

	mid := (left + right + 1) / 2
	if nums[mid] == target {
		return mid
	} else if target < nums[mid] {
		return binarySearch(nums, left, mid-1, target)
	} else {
		return binarySearch(nums, mid+1, right, target)
	}
}

func searchMatrixV2(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		find, idx := binarySearchV2(row, 0, len(row)-1, target)
		if find {
			return true
		}

		if idx >= len(row) {
			continue
		}

		if idx == 0 {
			return false
		}
	}
	return false
}

func binarySearchV2(nums []int, left, right, target int) (bool, int) {
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true, mid
		} else if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if nums[left] == target {
		return true, left
	} else if nums[left] < target {
		return false, left + 1
	} else {
		return false, left
	}
}
