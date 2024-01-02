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
