package array_string

func strStr(haystack string, needle string) int {
	start, j := 0, 0

	for ; start < len(haystack)-len(needle)+1; start++ {
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

/*
	this practice is relatively simple, but there are several problems or pits that need to consider:
	1. stop condition for outer loop:
		1.1 when remaining elements of haystack is less than needle, we don't need to iterate forward
		1.2 the ith (len(haystack)-len(needle)) needs to handle, like the corner case: haystack = "a", needle = "a"
	2. if haystack[i]==needle[j], we need to plus both i and j; if not we stop this internal loop and continue next iteration in outer loop
*/

func strStrV2(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		cur := i
		for j := 0; j < len(needle); j++ {
			if haystack[cur] != needle[j] {
				break
			} else {
				if j == len(needle)-1 {
					return i
				}
				cur++
			}
		}
	}
	return -1
}
