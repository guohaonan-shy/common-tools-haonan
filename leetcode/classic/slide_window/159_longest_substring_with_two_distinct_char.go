package slide_window

func lengthOfLongestSubstringTwoDistinct(s string) int {

	left, right := 0, 0
	longest := 1
	cnt := make(map[byte]int, 0)
	diff := 2

	for ; right < len(s); right++ {
		// 窗口右移动
		if _, ok := cnt[s[right]]; !ok {
			cnt[s[right]]++
			diff--
		} else {
			cnt[s[right]]++
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

func lengthOfLongestSubstringTwoDistinctV2(s string) int {
	set := make(map[byte]int, 0)
	left, right := 0, 0
	maxLength := 0

	for ; right < len(s); right++ { // regard right end as the iteration index
		char := s[right]
		set[char]++
		for ; left <= right && len(set) > 2; left++ {
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
