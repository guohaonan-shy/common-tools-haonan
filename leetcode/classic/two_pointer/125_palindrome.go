package two_pointer

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	res := true
	for left < right {
		// find an alphanumeric byte; because non-alphanumeric byte will be removed, we don't need to compare them
		if !isAlphaNumeric(s[left]) {
			left++
			continue
		}

		for !isAlphaNumeric(s[right]) {
			right--
			continue
		}

		// numeric number can be compared directly
		left_element, right_element := s[left], s[right]
		// here, there is a corner case: number 0 and character 'P'
		if left_element == right_element || (left_element >= 65 && right_element >= 65 && (left_element-right_element == 32 || right_element-left_element == 32)) {
			left++
			right--
		} else {
			res = false
			break
		}
	}
	return res
}

func isAlphaNumeric(s byte) bool {
	/*
		A-Z 的 ASCII码: 65 - 90
	*/
	if s >= 65 && s <= 90 {
		return true
	}

	/*
		a-z 的 ASCII码: 97 - 122
	*/
	if s >= 97 && s <= 122 {
		return true
	}

	/*
		48 ～ 57 是 0-9 的ASCII码
	*/
	if s >= 48 && s <= 57 { // 数字0-9的ASCII码
		return true
	}

	return false
}
