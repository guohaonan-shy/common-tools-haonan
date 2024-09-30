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

func minDistanceV2(word1, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for lengthWord1 := 0; lengthWord1 <= len(word1); lengthWord1++ {
		for lengthWord2 := 0; lengthWord2 <= len(word2); lengthWord2++ {
			if lengthWord1 == 0 && lengthWord2 == 0 {
				continue
			}

			if lengthWord1 == 0 {
				dp[0][lengthWord2] = dp[0][lengthWord2-1] + 1
				continue
			}

			if lengthWord2 == 0 {
				dp[lengthWord1][0] = dp[lengthWord1-1][0] + 1
				continue
			}

			// for normal cases
			// it is the same between delete and add a character so that two words can be the same => we can just consider add a character in one word and replace characters.
			// eg. word1 delete one character to make two words same == word2 add one character to make two words same

			// case1: add a character in word1 or add in word2
			minVal := min(dp[lengthWord1-1][lengthWord2], dp[lengthWord1][lengthWord2-1]) + 1
			// case2: replace
			replace := dp[lengthWord1-1][lengthWord2-1] + 1
			if word1[lengthWord1-1] == word2[lengthWord2-1] {
				replace -= 1
			}
			dp[lengthWord1][lengthWord2] = min(minVal, replace)
		}
	}
	return dp[len(word1)][len(word2)]
}
