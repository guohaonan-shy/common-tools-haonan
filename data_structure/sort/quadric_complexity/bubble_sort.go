package quadric_complexity

// BubbleSort non-decreasing sort
func BubbleSort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
	}
	return
}

func BubbleSort_EarlyTerminate(list []int) {
	for i := 0; i < len(list)-1; i++ {
		swap := false
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
				if !swap {
					swap = true
				}
			}
		}
		if !swap {
			break
		}
	}
	return
}
