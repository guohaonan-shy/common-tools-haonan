package dynammic_programming

/*
This method is ok but consume much memory when given string is too long
*/
func longestValidParentheses(s string) int {

	dp := make([][][]byte, len(s)+1)

	for i := range dp {
		dp[i] = make([][]byte, len(s))
		for j := range dp[i] {
			dp[i][j] = make([]byte, 0)
		}
	}

	res := 0

	for length := 1; length <= len(s); length++ {
		for start := 0; start+length <= len(s); start++ {
			cur := s[start+length-1]

			if cur == ')' && len(dp[length-1][start]) > 0 {
				curStack := dp[length-1][start]
				if curStack[len(curStack)-1] == '(' {
					curStack = curStack[0 : len(curStack)-1]
					dp[length][start] = curStack
					if len(curStack) == 0 {
						res = max(res, length)
					}
					continue
				}
			}

			dp[length][start] = append(dp[length-1][start], cur)
		}
	}
	return res
}

func longestValidParentheses_standard(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp := make([]int, len(s)) // 以idx为结尾的合法的子字符串长度
	res := 0
	for end := 1; end < len(s); end++ {
		if s[end] == '(' {
			continue
		}

		if s[end-1] == '(' { // '()' can add 2
			if end-2 >= 0 {
				dp[end] = dp[end-2] + 2
			} else {
				// corner case: end index == 1, which means current length is 2
				dp[end] = 2
			}
		} else {
			// '...))'
			// the matched index exists && the matched index's element is '('
			if end-dp[end-1]-1 >= 0 && s[end-dp[end-1]-1] == '(' {
				if end-dp[end-1]-2 >= 0 {
					dp[end] = dp[end-1] + 2 + dp[end-dp[end-1]-2]
				} else {
					dp[end] = dp[end-1] + 2
				}
			}
		}
		res = max(res, dp[end])
	}
	return res
}
