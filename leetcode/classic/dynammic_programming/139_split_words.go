package dynammic_programming

func splitWord(word string, wordDict []string) bool {
	dict := make(map[string]bool, len(wordDict))
	dp := make([]bool, len(word)+1)
	for i := range wordDict {
		dict[wordDict[i]] = true
	}
	dp[0] = true
	// supposed dp[i] is the result whether previous i characters can split to words in dict: 0~i-1
	// because word.length >= 1, end start from 1
	for end := 1; end <= len(word); end++ {
		for start := 0; start < end; start++ { // dp[0] = true, it means that we only care about word[0:end] in wordDict when start = 0
			if dp[start] && dict[word[start:end]] { // dp[start] means if 0~start-1 can split; word[start:end] to check if this string is in dict => dp[end] 0~start-1 && start~end-1 has solution
				dp[end] = true
				break
			}
		}
	}
	return dp[len(word)]
}
