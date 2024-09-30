package dynammic_programming

func longestPalindrome(s string) string {

	maxLength := 1
	left, right := 0, 1
	dp := make([][]bool, len(s))

	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for length := 1; length <= len(s); length++ {
		for start := 0; start+length <= len(s); start++ {
			if length == 1 {
				dp[start][start] = true

			} else if length == 2 {
				dp[start][start+1] = s[start] == s[start+1]
			} else {
				dp[start][start+length-1] = dp[start+1][start+length-2] && s[start] == s[start+length-1]
			}

			if dp[start][start+length-1] {
				if length > maxLength {
					maxLength = length
					left, right = start, start+length
				}
			}
		}
	}
	return s[left:right]
}
