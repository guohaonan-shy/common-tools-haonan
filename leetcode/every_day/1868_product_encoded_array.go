package every_day

func findRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int {
	left, right := 0, 0
	res := [][]int{}
	preProduct := 0
	for left < len(encoded1) && right < len(encoded2) {
		subArray1, subArray2 := encoded1[left], encoded2[right]
		product := subArray1[0] * subArray2[0]

		var localLength = 0

		if subArray1[1] > subArray2[1] {
			localLength = subArray2[1]
			subArray1[1] -= subArray2[1]
			right++
		} else if subArray1[1] < subArray2[1] {
			localLength = subArray1[1]
			subArray2[1] -= subArray1[1]
			left++
		} else {
			localLength = subArray1[1]
			left++
			right++
		}

		if product == preProduct { // consecutive
			res[len(res)-1][1] += localLength
		} else {
			res = append(res, []int{product, localLength})
		}

		preProduct = product
	}
	return res
}
