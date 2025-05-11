package dynammic_programming

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)

	for length := 1; length <= len(s); length++ {

		if s[length-1] != '0' {
			if length == 1 {
				dp[length] = 1
			} else {
				dp[length] += dp[length-1]
			}
		}

		if length >= 2 && isValid(s[length-2:]) {
			if length == 2 {
				dp[length] += 1
			} else {
				dp[length] += dp[length-2]
			}
		}
	}
	return dp[len(s)]
}

func isValid(str string) bool {
	if str[0] == '0' || str[0] > '2' {
		return false
	}

	if str[0] == '2' && ('7' <= str[1] && str[1] <= '9') {
		return false
	}
	return true
}
