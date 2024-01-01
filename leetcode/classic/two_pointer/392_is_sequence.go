package two_pointer

func isSubsequence(s string, t string) bool {
	haystack, needle := 0, 0
	for ; haystack < len(t) && needle < len(s); haystack++ {
		if t[haystack] == s[needle] {
			needle++
		}
	}

	if needle == len(s) {
		return true
	}
	return false
}
