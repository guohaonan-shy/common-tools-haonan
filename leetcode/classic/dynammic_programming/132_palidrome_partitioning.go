package dynammic_programming

import "math"

func minCut(s string) int {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s))
		if i == 0 || i == 1 {
			for j := 0; j < len(s); j++ {
				dp[i][j] = true
			}
		}
	}
	// preprocess: query if any substring is a palidrome => 2-D dynamic programming
	for length := 2; length <= len(s); length++ {
		for idx := 0; idx <= len(s)-length; idx++ {
			dp[length][idx] = dp[length-2][idx+1] && (s[idx] == s[idx+length-1])
		}
	}

	res := make([]int, len(s)+1)
	// res[i] => 长度为i的字符串需要至少几次partition才能变成palidrome => 1-D dynamic programming
	// 任意一个字符串都可以切割成palindrome => 切割成一个一个单字符
	for length := 2; length <= len(s); length++ {
		if dp[length][0] {
			continue
		}

		res[length] = math.MaxInt32
		for idx := 1; idx < length; idx++ {
			if dp[length-idx][idx] && res[idx]+1 < res[length] {
				res[length] = res[idx] + 1
			}
		}
	}
	return res[len(s)]
}
