package back_track

func combinationSum(candidates []int, target int) [][]int {
	temp := make([]int, 0)
	res := make([][]int, 0)

	var dfs func(target int, idx int)
	dfs = func(target int, idx int) {
		if target == 0 {
			ans := make([]int, len(temp))
			copy(ans, temp)
			res = append(res, ans)
			return
		}

		for i, candidate := range candidates {

			if target-candidate < 0 || i < idx {
				continue
			}

			temp = append(temp, candidate)
			dfs(target-candidate, i)
			temp = temp[:len(temp)-1]

		}

	}
	dfs(target, 0)
	return res
}
