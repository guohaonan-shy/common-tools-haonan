package back_track

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	temp := []int{}
	res := make([][]int, 0)
	var dfs func(cur int)

	dfs = func(cur int) {
		ans := make([]int, len(temp))
		copy(ans, temp)
		res = append(res, ans)

		for i := cur; i < len(nums); i++ {
			if i > cur && nums[i-1] == nums[i] {
				continue
			}
			temp = append(temp, nums[i])
			dfs(i + 1)
			temp = temp[:len(temp)-1]
		}
		return
	}

	dfs(0)
	return res
}
