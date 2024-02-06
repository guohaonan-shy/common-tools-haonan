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
