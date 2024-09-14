package slide_window

import "math"

func minWindow(s string, t string) string {
	dict := make(map[byte]int, 0)
	for i := range t {
		dict[t[i]]++
	}

	cnt := make(map[byte]int, 0)

	diff := len(t)
	ansL, ansR := -1, -1
	minLength := math.MaxInt32
	for left, right := 0, 0; right < len(s); right++ {
		if _, ok := dict[s[right]]; !ok { // 非t中的元素， 不处理
			continue
		}

		cnt[s[right]]++
		if cnt[s[right]] <= dict[s[right]] { //右移主要是为了减少diff，此时新增元素是否满足最低要求，即t中的元素要求
			diff--
		}
		for diff == 0 { // diff满足，不断左移，减少非必要元素，直到不满足条件
			if _, ok := dict[s[left]]; !ok { // 非t中的元素， 不处理
				left++ // 下一个
				continue
			}

			if minLength > right-left+1 { // 最小迭代
				minLength = right - left + 1
				ansL, ansR = left, right
			}

			cnt[s[left]]--
			if cnt[s[left]] < dict[s[left]] { // 同上
				diff++
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

func minWindowV2(s string, t string) string {
	left, right := 0, 0
	resL, resR := -1, -1
	minLength := math.MaxInt32
	dict := make(map[byte]int, 0)
	for i := range t {
		dict[t[i]] += 1
	}
	cnt := make(map[byte]int, 0)
	diff := len(t)

	res := ""

	for ; right < len(s); right++ {
		if _, ok := dict[s[right]]; !ok {
			continue
		}

		cnt[s[right]]++
		if cnt[s[right]] <= dict[s[right]] {
			diff--
		}

		for ; diff == 0; left++ {
			if _, ok := dict[s[left]]; !ok {
				continue
			}
			if minLength > right-left+1 {
				minLength = right - left + 1
				resL, resR = left, right+1
			}
			cnt[s[left]]--
			/*
				corner case: cnt[byte] > dict[byte] => cnt[byte] = dict[byte]; this case doesn't produce diff
			*/
			if cnt[s[left]] < dict[s[left]] {
				diff++
			}
		}
	}
	if resL != -1 {
		res = s[resL:resR]
	}
	return res
}
