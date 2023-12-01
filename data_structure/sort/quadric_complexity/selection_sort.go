package compared_based

import "math"

func findIndexOfMin(list []int) int {
	index := 0
	minValue := math.MaxInt32
	for i, value := range list {
		if value < minValue {
			minValue = value
			index = i
		}
	}
	return index
}

func SelectionSort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		minIndex := findIndexOfMin(list[i+1:])
		minIndex += i + 1

		if list[i] > list[minIndex] {
			temp := list[minIndex]
			list[minIndex] = list[i]
			list[i] = temp
		}

	}
}
