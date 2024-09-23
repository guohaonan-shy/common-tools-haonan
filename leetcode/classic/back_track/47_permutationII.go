package back_track

func permuteUnique(nums []int) [][]int {
	k := len(nums)
	temp := make([]int, 0)
	res := make([][]int, 0)
	state := make([]bool, k)

	var dfs func()
	dfs = func() {

		if len(temp) == k {
			ele := make([]int, k)
			copy(ele, temp)
			res = append(res, ele)
			return
		}

		hash := make(map[int]bool, 0)

		for i := 0; i < k; i++ {
			if !state[i] && !hash[nums[i]] {
				state[i] = true
				hash[nums[i]] = true
				temp = append(temp, nums[i])
				dfs()
				temp = temp[:len(temp)-1]
				state[i] = false
			}
		}
	}
	dfs()
	return res
}

func permuteUniqueV2(nums []int) [][]int {
	res := make([][]int, 0)
	tempPermutation := make([]int, len(nums))

	var dfs func(cur int)
	dfs = func(cur int) {
		if len(tempPermutation) == len(nums) {
			permutation := make([]int, len(nums))
			copy(permutation, tempPermutation)
			res = append(res, permutation)
			return
		}

		for start := cur; start < len(nums); start++ {
			if start > cur && nums[start] == nums[start-1] {
				continue
			}
			tempPermutation = append(tempPermutation, nums[start])
			dfs(start + 1)
			tempPermutation = tempPermutation[:len(tempPermutation)-1]
		}
		return
	}
	dfs(0)
	return res
}
