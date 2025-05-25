package dynammic_programming

func partition(s string) [][]string {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s))
		if i == 0 || i == 1 {
			for j := 0; j < len(s); j++ {
				dp[i][j] = true
			}
		}
	}

	for length := 2; length <= len(s); length++ {
		for idx := 0; idx <= len(s)-length; idx++ {
			dp[length][idx] = dp[length-2][idx+1] && (s[idx] == s[idx+length-1])
		}
	}

	var dfs func(cur int) // cur => s[cur]以及之前的substring作为一个完整的string，针对cur之后的元素进行切割
	res := make([][]string, 0)
	temp := make([]string, 0)

	dfs = func(cur int) {
		if cur == len(s) {
			ans := make([]string, len(temp))
			copy(ans, temp)
			res = append(res, ans)
			return
		}

		for i := cur; i < len(s); i++ {
			if dp[i-cur+1][cur] {
				temp = append(temp, s[cur:i])
				dfs(i + 1)
				temp = temp[:len(temp)-1]
			}
		}
	}
	dfs(0)
	return res
}
