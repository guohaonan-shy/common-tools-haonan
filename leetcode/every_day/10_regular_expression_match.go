package every_day

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true
	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			str := p[j-1]
			if str-'a' < 26 {
				if i != 0 {
					dp[i][j] = dp[i-1][j-1] && s[i-1] == p[j-1]
				}
			} else if str == '.' {
				if i != 0 {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				dp[i][j] = dp[i][j-2] // x* = ""
				// 如果x*代替多个x，那么末尾元素必须是x，否则无法代替
				if i != 0 && (s[i-1] == p[j-2] || p[j-2] == '.') { //i != 0 代表有所匹配
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	return dp[len(s)][len(p)]
}
