package every_day

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutiveAnswerBySpecifiedAnswer(answerKey, k, 'T'), maxConsecutiveAnswerBySpecifiedAnswer(answerKey, k, 'F'))
}

func maxConsecutiveAnswerBySpecifiedAnswer(answerKey string, k int, target byte) int {
	if len(answerKey) == 0 {
		return 0
	}

	left, right := 0, 0
	globalMax := 0
	for left <= right && right < len(answerKey) {
		if answerKey[right] == target || (k > 0 && answerKey[right] != target) {
			if k > 0 && answerKey[right] != target {
				k--
			}
			globalMax = max(globalMax, right-left+1)
			right++
			continue
		}

		// need left move the window
		if answerKey[left] != target {
			k++
		}
		left++
	}
	return globalMax
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
