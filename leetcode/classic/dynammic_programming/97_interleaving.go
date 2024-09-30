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

func isInterleaveV2(s1 string, s2 string, s3 string) bool {

	if len(s1) == 0 { // process the corner case: s1 == 0
		return s2 == s3
	}

	if len(s2) == 0 { // process the corner case: s2 == 0
		return s1 == s3
	}
	// both of above process steps has already solved the case where all strings are empty
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true

	for length := 1; length <= len(s3); length++ {
		for lengthS1 := 0; lengthS1 <= len(s1) && lengthS1 <= length; lengthS1++ {
			// filter the wrong condition that s2 pointer is greater than s2 length
			// Don't use this in the loop conditions: lengthS1 <= len(s1) && lengthS1 <= length && """length-lengthS1 <= len(s2)"""
			// because we iterate s1 in increasing order, the order of s2 is reverse so that we will miss some cases when length - lengthS1 <= len(s2)
			if length-lengthS1 > len(s2) {
				continue
			}

			if lengthS1 == 0 {
				dp[0][length] = s2[:length] == s3[:length]
				continue
			}

			if lengthS1 == length {
				dp[length][0] = s1[:length] == s3[:length]
				continue
			}

			dp[lengthS1][length-lengthS1] = (dp[lengthS1-1][length-lengthS1] && s1[lengthS1-1] == s3[length-1]) ||
				(dp[lengthS1][length-lengthS1-1] && s2[length-lengthS1-1] == s3[length-1])
		}
	}
	return dp[len(s1)][len(s2)]
}
