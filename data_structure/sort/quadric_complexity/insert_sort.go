package compared_based

func InsertSort(list []int) {
	for i := 1; i < len(list); i++ {
		pivot := list[i]
		j := i - 1
		for ; j > -1; j-- {
			if list[j] > pivot {
				list[j+1] = list[j]
			} else {
				break
			}
		}

		list[j+1] = pivot

	}
}
