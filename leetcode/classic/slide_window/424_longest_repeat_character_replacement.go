package slide_window

func characterReplacement(s string, k int) int {
	left, right := 0, 0
	set := make(map[byte]int, 0)
	maxLength := 0

	for ; right < len(s); right++ {
		rightChar := s[right]
		set[rightChar]++

		for ; left <= right && getMode(set) > k; left++ {
			leftChar := s[left]
			set[leftChar]--
			if set[leftChar] == 0 {
				delete(set, leftChar)
			}
		}

		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

/*
	time complexity is high if we need to iterate set to calculate the count of elements except mode => a data structure that maintain the sorted list with counts => heap
*/
func getMode(set map[byte]int) int {
	maxCnt, totalCnt := 0, 0
	for _, cnt := range set {
		if cnt > maxCnt {
			maxCnt = cnt
		}
		totalCnt += cnt
	}
	return totalCnt - maxCnt
}
