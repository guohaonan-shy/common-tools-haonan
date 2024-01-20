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
