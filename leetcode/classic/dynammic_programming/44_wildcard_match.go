package dynammic_programming

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true

	if p[0] == '*' {
		for i := 0; i <= len(s); i++ {
			dp[i][1] = true
		}
	}

	for j := 1; j <= len(p); j++ {
		for i := 1; i <= len(s); i++ {
			if p[j-1] == '*' {

				for length := i; length >= 0; length-- {
					if dp[length][j-1] {
						dp[i][j] = true
						// pre = true
						break
					}
				}
			} else if p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] && p[j-1] == s[i-1]
			}

		}
	}
	return dp[len(s)][len(p)]
}

func isMatchV2(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true
	/*
		case:
		1. when the pattern is empty string => only if target string is empty, match success
		2. when current pattern byte is '*' => 长度从0到current string的字符串都可被匹配
		3. when current pattern byte is '?' 或 字母 => 目标字符串长度至少为1且符合对应条件
	*/
	for patternLength := 1; patternLength <= len(p); patternLength++ {
		for stringLength := 0; stringLength <= len(s); stringLength++ {

			if stringLength == 0 && p[patternLength-1] != '*' {
				continue
			}

			if p[patternLength-1] == '?' {
				dp[stringLength][patternLength] = dp[stringLength-1][patternLength-1]
			} else if p[patternLength-1] == '*' {
				res := false
				for i := 0; i <= stringLength; i++ {
					if dp[i][patternLength-1] {
						res = true
						break
					}
				}

				dp[stringLength][patternLength] = res
			} else {
				dp[stringLength][patternLength] = dp[stringLength-1][patternLength-1] && s[stringLength-1] == p[patternLength-1]
			}

		}

	}
	return dp[len(s)][len(p)]
}
