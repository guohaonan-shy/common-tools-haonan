package logarithm_complexity

func MergeSort(list []int) []int {
	if len(list) == 1 {
		return list
	}

	// divide and merge
	// divide
	mid := len(list) / 2
	left := MergeSort(list[0:mid])
	right := MergeSort(list[mid:])
	// merge
	result := Merge(left, right)
	return result
}

func Merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	if i == len(left) {
		result = append(result, right[j:]...)
	}

	if j == len(right) {
		result = append(result, left[i:]...)
	}

	return result
}
