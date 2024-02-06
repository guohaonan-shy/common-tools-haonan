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
