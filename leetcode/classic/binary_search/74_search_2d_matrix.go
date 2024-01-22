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
