package stack

func isValid(s string) bool {
	stack := make([]byte, 0)

	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (left == '(' && s[i] == ')') || (left == '[' && s[i] == ']') || (left == '{' && s[i] == '}') {
				continue
			}
			return false
		}
	}
	return true
}
