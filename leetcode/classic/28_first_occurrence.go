package classic

func strStr(haystack string, needle string) int {
	start, j := 0, 0

	for ; start < len(haystack)-len(needle); start++ {
		for i := start; i < start+len(needle) && j < len(needle); {
			if haystack[i] == needle[j] {
				j++
				i++
			} else {
				j = 0
				break
			}
		}

		if j == len(needle) {
			return start
		}

	}

	return -1
}
