package sort

func QuickSort(list []int) []int {
	if len(list) == 1 {
		return list
	}

	pos := Partition(list)
	if pos > 0 {
		QuickSort(list[0:pos])
	}

	if pos < len(list)-1 {
		QuickSort(list[pos+1:])
	}

	return list
}

// Partition time complexity is O(N), because this function is a single loop and
func Partition(list []int) int {
	pivot := list[0]
	pos := 0
	for i := 1; i < len(list); i++ {
		if list[i] < pivot {
			pos++
			temp := list[i]
			list[i] = list[pos]
			list[pos] = temp
		}
	}

	// swap pivot with pos
	temp := list[pos]
	list[pos] = pivot
	list[0] = temp
	return pos
}
