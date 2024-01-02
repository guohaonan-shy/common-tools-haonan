package slide_window

import "math"

func minWindow(s string, t string) string {
	dict := make(map[byte]int, 0)
	for i := range t {
		dict[s[i]]++
	}

	cnt := make(map[byte]int, 0)

	check := func() bool {
		for key, value := range dict {
			if value > cnt[key] {
				return false
			}
		}
		return true
	}
	ansL, ansR := -1, -1
	minLength := math.MaxInt32
	for left, right := 0, 0; right < len(s); right++ {
		if _, ok := dict[s[right]]; ok {
			cnt[s[right]]++
		}
		for check() && left <= right {
			if minLength > right-left+1 {
				minLength = right - left + 1
				ansL, ansR = left, right
			}

			if _, ok := dict[s[left]]; ok {
				cnt[s[left]]--
			}
			left++
		}
	}

	if ansR == -1 {
		return ""
	} else {
		return s[ansL : ansR+1]
	}
}
