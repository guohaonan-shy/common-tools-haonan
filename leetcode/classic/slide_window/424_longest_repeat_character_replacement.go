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

/*
First, I want to emphasize that 'maxCnt' is max counts from 0 idx to current one rather than max counts in window. The solution ton maintain a max count in window is above.
*/
func characterReplacementV2(s string, k int) int {
	left, right := 0, 0
	set := make(map[byte]int, 0)
	maxCnt := 0

	for ; right < len(s); right++ {
		rightChar := s[right]
		set[rightChar]++
		maxCnt = max(maxCnt, set[rightChar])
		/*
			Here, we maintain a historical max counts;
			When we left move the window, the max count in the window has only two results:
				1. maintain unchanged, because the element we removed is not the mode in window or the max count in this window (might be other elements) is the same with previous one
				2. decreased, the mode changes to another element whose count is lower than previous one or the max counts of this element is still the largest but decreased

			In addition, we expect to find the longest substring for requirements. The first window we found must start from 0 index, so 'right-left+1(window size) - maxcnt <= k' is the length what we expected
			Therefore, if we want to find a larger window size, we have to find a maxcnt that is greater than current one, window_size = maxcnt+k. The larger maxcnt is, the larger window size is.

			When maxcnt become greater, we can reduce the time to move left side. Or we just maintain the window size as the previous longest one.
			We don't need to focus on the window whose size is lower than current maximum.
		*/
		if right-left+1-maxCnt > k {
			leftChar := s[left]
			set[leftChar]--
			left++
		}
	}
	return len(s) - left
}
