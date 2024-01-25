package dynammic_programming

func isInterleave(s1 string, s2 string, s3 string) bool {

	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	dp[0][0] = true

	for length := 1; length <= len(s3); length++ {
		for i := 0; i <= length; i++ {
			col := length - i
			if i > len(s1) || col > len(s2) {
				continue
			}

			if i-1 >= 0 {
				dp[i][col] = dp[i][col] || (dp[i-1][col] && s3[length-1] == s1[i-1])
			}

			if length-i-1 >= 0 {
				dp[i][col] = dp[i][col] || (dp[i][col-1] && s3[length-1] == s2[col-1])
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
