package back_track

func generateParenthesis(n int) []string {
	temp := []byte{}
	res := make([]string, 0)

	var dfs func(left, right int)
	dfs = func(left, right int) {

		if right == 0 {
			ans := make([]byte, len(temp))
			copy(ans, temp)
			res = append(res, string(ans))
			return
		}

		if left == right {
			temp = append(temp, '(')
			dfs(left-1, right)
			temp = temp[:len(temp)-1]
		}

		if left < right {
			// left and right
			if left != 0 {
				temp = append(temp, '(')
				dfs(left-1, right)
				temp = temp[:len(temp)-1]
			}

			if right != 0 {
				temp = append(temp, ')')
				dfs(left, right-1)
				temp = temp[:len(temp)-1]
			}
		}

	}
	dfs(n, n)
	return res
}
