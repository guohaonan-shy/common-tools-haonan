package back_track

func combine(n int, k int) [][]int {
	return combineBackTrack(1, n, k)
}

func combineBackTrack(start, end int, k int) [][]int {

	if k == 1 {
		ans := make([][]int, 0)
		for i := start; i <= end; i++ {
			ans = append(ans, []int{i})
		}
		return ans
	}

	ans := make([][]int, 0)
	for j := start; j <= end-k+1; j++ {

		val := j

		options := combineBackTrack(j+1, end, k-1)

		for _, option := range options {
			temp := []int{val}
			temp = append(temp, option...)

			ans = append(ans, temp)
		}

	}
	return ans

}

func combine_standard(n, k int) [][]int {

	tempPath := make([]int, 0)
	res := make([][]int, 0)

	var dfs func(start, idx int)
	dfs = func(start, idx int) {

		if idx > k {
			ans := make([]int, len(tempPath))
			copy(ans, tempPath)
			res = append(res, ans)
			return
		}

		for i := start; i <= n; i++ {
			if n-i < k-idx {
				break
			}

			tempPath = append(tempPath, i)
			dfs(i+1, idx+1)
			tempPath = tempPath[:len(tempPath)-1]
		}
	}

	dfs(1, 1)
	return res
}

func combineV2(n, k int) [][]int {
	tempCombination := make([]int, 0)
	res := make([][]int, 0)

	var dfs func(cur int)

	dfs = func(cur int) {
		if len(tempCombination) == k {
			combination := make([]int, k)
			copy(combination, tempCombination)
			res = append(res, combination)
			return
		}

		for start := cur; start <= n && start <= n+1-k+len(tempCombination); start++ {
			tempCombination = append(tempCombination, start)
			dfs(start)
			tempCombination = tempCombination[0 : len(tempCombination)-1]
		}
		return
	}
	dfs(1)
	return res
}
