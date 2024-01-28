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
