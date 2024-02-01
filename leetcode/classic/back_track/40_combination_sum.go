package back_track

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	temp := make([]int, 0)
	res := make([][]int, 0)

	var dfs func(goal int, idx int)
	dfs = func(goal int, idx int) {

		if goal == 0 {
			newItem := make([]int, len(temp))
			copy(newItem, temp)
			res = append(res, newItem)
			return
		}

		for i := idx; i < len(candidates); i++ {
			if goal-candidates[i] < 0 {
				break
			}
			temp = append(temp, candidates[i])
			dfs(goal-candidates[i], i+1)
			temp = temp[:len(temp)-1]
		}
	}

	dfs(target, 0)
	return res
}
