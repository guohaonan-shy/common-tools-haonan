package slide_window

func longestValidParentheses(s string) int {

	longest := 0
	stack := make([]int, 0)

	for i := range s {

		if s[i] == ')' {
			if len(stack) == 0 {
				continue
			} else {
				idx := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				longest = max(longest, i-idx+1)
			}
		} else {
			stack = append(stack, i)
		}
	}

	return longest
}
