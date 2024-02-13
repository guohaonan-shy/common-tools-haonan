package two_pointer

func largestRectangleArea(heights []int) int {
	idx := maxIdx(heights)

	maxRec := 0
	for _, id := range idx {
		left, right := id, id
		length := heights[id]
		for {
			var leftVal, rightVal int
			if left > 0 {
				leftVal = heights[left-1]
			}
			if right < len(heights)-1 {
				rightVal = heights[right+1]
			}

			if leftVal < rightVal {
				if right < len(heights)-1 {
					right++
				}
				maxRec = max(min(length, rightVal)*(right-left+1), maxRec)
				length = min(length, rightVal)
			} else if leftVal > rightVal {
				if left > 0 {
					left--
				}
				maxRec = max(min(length, leftVal)*(right-left+1), maxRec)
				length = min(length, leftVal)
			} else {
				if left > 0 {
					left--
				}
				if right < len(heights)-1 {
					right++
				}
				maxRec = max(min(length, leftVal)*(right-left+1), maxRec)
				length = min(length, leftVal)
			}

			if left <= 0 && right >= len(heights)-1 {
				break
			}
		}

	}
	return maxRec
}

func maxIdx(nums []int) []int {
	maxVal := -1
	idx := make([]int, 0)
	for i, val := range nums {

		if val > maxVal {
			idx = []int{i}
			maxVal = val
			continue
		}

		if val == maxVal {
			idx = append(idx, i)
		}

	}
	return idx
}
