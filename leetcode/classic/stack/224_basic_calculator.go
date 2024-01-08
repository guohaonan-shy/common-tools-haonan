package stack

func calculate(s string) int {
	numberStack := make([][]int, 1)
	numberStack[0] = make([]int, 0)
	tokenStack := make([][]string, 1)
	tokenStack[0] = make([]string, 0)
	layer := 0

	for i := 0; i < len(s); {
		str := string(s[i])
		if str == " " {
			i++
			continue
		}
		if str >= "0" && str <= "9" {
			tempStack := make([]int, 0)
			for ; i < len(s) && string(s[i]) >= "0" && string(s[i]) <= "9"; i++ {
				if string(s[i]) == " " {
					continue
				}
				tempStack = append(tempStack, int(s[i]-'0'))
			}
			digit := 1
			number := 0
			for j := len(tempStack) - 1; j >= 0; j-- {
				number += digit * tempStack[j]
				digit *= 10
			}

			numberStack[layer] = append(numberStack[layer], number)

		} else {

			if str == "(" {
				layer++
				numberStack = append(numberStack, make([]int, 0))
				tokenStack = append(tokenStack, make([]string, 0))
			} else if str == ")" {
				res := numberStack[layer][0]
				numberStack[layer] = numberStack[layer][:0]
				layer--
				numberStack[layer] = append(numberStack[layer], res)
			} else {
				tokenStack[layer] = append(tokenStack[layer], str)
			}
			i++
		}

		for len(tokenStack[layer]) > 0 && (tokenStack[layer][len(tokenStack[layer])-1] == "+" || tokenStack[layer][len(tokenStack[layer])-1] == "-") && len(numberStack[layer]) >= 2 {
			left, right := numberStack[layer][len(numberStack[layer])-2], numberStack[layer][len(numberStack[layer])-1]
			numberStack[layer] = numberStack[layer][:len(numberStack[layer])-2]
			token := tokenStack[layer][len(tokenStack[layer])-1]
			tokenStack[layer] = tokenStack[layer][:len(tokenStack[layer])-1]
			var res = 0
			switch token {
			case "+":
				res = left + right
			case "-":
				res = left - right
			}

			numberStack[layer] = append(numberStack[layer], res)
		}

	}

	return numberStack[0][0]
}

func calculate_standard(s string) (ans int) {
	ops := []int{1} // 默认为+
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1] // 影响当前位置符号功能的点：最近的一个左括号前的符号（如果整个字符串没有括号，默认最外层就是一个括号）
			i++
		case '-':
			sign = -ops[len(ops)-1] // 最近一个是负号，相当于没有括号取反
			i++
		case '(':
			ops = append(ops, sign) // 最近
			i++
		case ')':
			ops = ops[:len(ops)-1] // 去掉对应左括号前的符号
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return
}
