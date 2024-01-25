package dynammic_programming

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		if i == 0 {
			for j := 0; j < len(word2)+1; j++ {
				dp[0][j] = j
			}
		}

		dp[i][0] = i
	}
	// i,j 表示word1[:i] -> word2[:j]的编辑距离，比如dp[1][1]，长度为1的word1 -> 长度为1的word2的最小编辑距离
	// take word1 = "horse", word2 = "ros" as a case
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1) // dp[i-1][j]为hors -> ros的编辑距离，在这个基础上只需要删除horse末尾的e即可完成变换；同理，dp[i][j-1]为horse -> ro，在这个基础上，之后多一步添加s即可
			// 此处要注意，当使用replace将word1[i-1]替换成word2[j-1]时，要注意，如果相等，则不需要替换
			replace := dp[i-1][j-1] + 1
			if word1[i-1] == word2[j-1] {
				replace -= 1
			}

			dp[i][j] = min(replace, dp[i][j])
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
