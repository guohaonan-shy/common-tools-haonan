package stack

func isValid(s string) bool {
	stack := make([]byte, 0)

	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}

		if len(stack) == 0 {
			// right parentheses has no left one to match
			return false
		}

		left := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if (left == '(' && s[i] == ')') || (left == '[' && s[i] == ']') || (left == '{' && s[i] == '}') {
			// right parentheses has left one to match a complete parentheses, continue to match futher
			continue
		}
		return false
	}

	/*
		if left parentheses is more than right ones, the stack has remaining elements
		or
		all parentheses is matched
	*/
	return len(stack) == 0
}
