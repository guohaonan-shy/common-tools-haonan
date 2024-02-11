package array_string

import "strings"

func isNumber(s string) bool {
	index := -1
	for i := range s {
		if s[i] == 'e' || s[i] == 'E' {
			index = i
			break
		}
	}

	if index != -1 {
		return (isDecimal(s[:index]) || isInteger(s[:index], false)) && isInteger(s[index+1:], true)
	} else {
		return isInteger(s, false) || isDecimal(s)
	}
}

func isDecimal(s string) bool {
	if s == "" {
		return false
	}

	dotIdx := strings.Index(s, ".")

	if dotIdx == -1 {
		return false
	}

	prev, after := s[:dotIdx], s[dotIdx+1:]

	if (len(prev) == 0 || (len(prev) == 1 && (prev[0] == '+' || prev[0] == '-'))) && (len(after) == 0 || (len(after) == 1 && (after[0] == '+' || after[0] == '-'))) {
		return false
	}

	if len(prev) != 0 && !isInteger(prev, false) {
		return false
	}

	if len(after) != 0 {
		isSign := after[0] == '+' || after[0] == '-'
		if isSign {
			return false
		}

		if !isInteger(after, true) {
			return false
		}
	}
	return true
}

func isInteger(s string, isAfter bool) bool {
	if s == "" {
		return false
	}
	isSign := s[0] == '+' || s[0] == '-'
	start := 0
	if isSign {
		start = 1
	}

	if isAfter && start >= len(s) {
		return false
	}

	for ; start < len(s); start++ {
		if s[start] == '+' || s[start] == '-' {
			return false
		}

		if s[start]-'0' < 0 || s[start]-'0' > 9 {
			return false
		}
	}
	return true
}
