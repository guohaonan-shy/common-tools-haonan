package classic

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	res := true
	for left < right {
		left_is_letter := isAlphaNumeric(rune(s[left]))
		if !left_is_letter {
			left++
			continue
		}

		right_is_letter := isAlphaNumeric(rune(s[right]))
		if !right_is_letter {
			right--
			continue
		}

		left_element := s[left]
		if s[left] >= 65 && s[left] <= 90 {
			left_element = s[left] + 32
		}

		right_element := s[right]
		if s[right] >= 65 && s[right] <= 90 {
			right_element = s[right] + 32
		}

		if left_element == right_element {
			left++
			right--
		} else {
			res = false
			break
		}
	}
	return res
}

func isAlphaNumeric(s rune) bool {
	if s >= 65 && s <= 90 {
		return true
	}

	if s >= 97 && s <= 122 {
		return true
	}

	if s >= 48 && s <= 57 {
		return true
	}

	return false
}
