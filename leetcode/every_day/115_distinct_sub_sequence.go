package every_day

func numDistinct(s string, t string) int {
	dp := make([][]int, len(t)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}

	for lenOfT := 1; lenOfT <= len(t); lenOfT++ {
		cnt := 0
		for lenOfS := lenOfT; lenOfS <= len(s); lenOfS++ {
			if s[lenOfS-1] == t[lenOfT-1] {
				if lenOfT == 1 {
					// 初始化起始的子序列
					cnt += 1
				} else {
					cnt += dp[lenOfT-1][lenOfS-1] // dp[lenOfT-1][lenOfS-1] => t[:len(t)-1] 作为子序列在s[:len(s)-1]内有几个
					// why 累加？=> 回归问题本质，子序列个数是序列s内包含的前缀子序列t[:len(t)-1]的个数 * s[len(t)-1:len(s)-1]内同t[len(t)-1]相同的个数
				}
				dp[lenOfT][lenOfS] = cnt
			} else {
				dp[lenOfT][lenOfS] = dp[lenOfT][lenOfS-1]
			}
		}
	}
	return dp[len(t)][len(s)]
}
