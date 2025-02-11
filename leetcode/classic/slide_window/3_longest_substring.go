package slide_window

func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	window := make(map[byte]int, 0)
	maxSubstring := 0
	for end < len(s) {
		var (
			index = -1
			ok    = false
		)

		if index, ok = window[s[end]]; ok {
			if index >= start { // map内元素不在窗口内
				start = index + 1
			}
		}
		maxSubstring = max(maxSubstring, end-start+1)
		window[s[end]] = end
		end++
	}
	return maxSubstring
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstringV2(s string) int {
	left, right := 0, 0
	maxLength := 0
	subString := make(map[byte]struct{}, 0)
	for ; right < len(s); right++ {
		_, exist := subString[s[right]]
		for ; left <= right && exist; left++ {
			delete(subString, s[left])
			_, exist = subString[s[right]]
		}
		subString[s[right]] = struct{}{}
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

// longest repeat character
func lengthOfLongestSubstringV3(s string) int {
	if len(s) == 0 {
		return 0
	}

	target := s[0]
	left, right := 0, 0
	maxLength := 0

	for right < len(s) {
		for ; right < len(s) && s[right] == target; right++ {
			if right-left+1 > maxLength {
				maxLength = right - left + 1
			}
		}

		if right == len(s) {
			break
		}

		target = s[right]
		left = right
	}
	return maxLength
}
