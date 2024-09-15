package slide_window

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	left, right := 0, 0
	cnt := make(map[byte]int, 0)
	diff := k
	longest := 0

	for ; right < len(s); right++ {
		if _, ok := cnt[s[right]]; ok {
			cnt[s[right]]++
		} else {
			cnt[s[right]]++
			diff--
		}

		if diff >= 0 && right-left+1 > longest {
			longest = right - left + 1
		}

		for diff < 0 {
			cnt[s[left]]--
			if cnt[s[left]] == 0 {
				delete(cnt, s[left])
				diff++
			}

			left++
		}
	}
	return longest
}

func lengthOfLongestSubstringKDistinctV2(s string, k int) int {
	left, right := 0, 0
	maxLength := 0
	set := make(map[byte]int, 0)

	for ; right < len(s); right++ {
		rightChar := s[right]
		set[rightChar]++

		for ; left <= right && len(set) > k; left++ {
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
