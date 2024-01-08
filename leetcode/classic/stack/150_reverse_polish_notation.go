package stack

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, str := range tokens {
		if str != "+" && str != "-" && str != "*" && str != "/" {
			integer, _ := strconv.Atoi(str)
			stack = append(stack, integer)
		} else {
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var res int = 0
			switch str {
			case "+":
				res = left + right
			case "-":
				res = left - right
			case "*":
				res = left * right
			case "/":
				res = left / right
			}
			stack = append(stack, res)
		}
	}
	return stack[0]
}
