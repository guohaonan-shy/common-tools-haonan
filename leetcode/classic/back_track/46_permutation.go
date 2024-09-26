package back_track

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	ans := make([][]int, 0)
	for idx := range nums {
		temp := make([]int, len(nums)-1)

		copy(temp[:idx], nums[:idx])
		copy(temp[idx:], nums[idx+1:])

		choices := permute(temp)

		an := make([][]int, 0)
		for _, choice := range choices {
			an = append(an, append([]int{nums[idx]}, choice...))
		}

		ans = append(ans, an...)
	}
	return ans
}

func permute_standard(nums []int) [][]int {
	res := make([][]int, 0)
	temp := make([]int, 0)
	k := len(nums)
	state := make([]bool, len(nums))

	var dfs func(int)
	dfs = func(idx int) {
		if idx > k {
			ans := make([]int, k)
			copy(ans, temp)
			res = append(res, ans)
			return
		}

		for i, num := range nums {
			if !state[i] {
				state[i] = true
				temp = append(temp, num)
				dfs(idx + 1)
				temp = temp[:len(temp)-1]
				state[i] = false
			}
		}
	}

	dfs(1)
	return res
}

func permuteV2(nums []int) [][]int {
	res := make([][]int, 0)
	tempPermutation := make([]int, 0)
	reached := make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(tempPermutation) == len(nums) {
			permutation := make([]int, len(nums))
			copy(permutation, tempPermutation)
			res = append(res, permutation)
			return
		}

		for i, num := range nums {
			if !reached[i] {
				tempPermutation = append(tempPermutation, num)
				reached[i] = true
				dfs()
				reached[i] = false
				tempPermutation = tempPermutation[0 : len(tempPermutation)-1]
			}
		}
		return
	}
	dfs()
	return res
}
