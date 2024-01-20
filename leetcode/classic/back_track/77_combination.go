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
