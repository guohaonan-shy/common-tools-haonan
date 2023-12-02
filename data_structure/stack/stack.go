package stack

import "strconv"

func PostfixExpression(input []string) int {
	stack := make([]int, 0)
	for _, str := range input {
		if str == "+" || str == "-" || str == "*" || str == "/" {
			para1 := stack[len(stack)-1]
			para2 := stack[len(stack)-2]

			stack = stack[:len(stack)-1]
			stack = stack[:len(stack)-1]

			result := 0
			switch str {
			case "+":
				result = para1 + para2
			case "-":
				result = para1 - para2
			case "*":
				result = para1 * para2
			case "/":
				result = para1 / para2
			}

			stack = append(stack, result)

		} else {
			value, _ := strconv.Atoi(str)
			stack = append(stack, value)
		}
	}

	return stack[len(stack)-1]
}
