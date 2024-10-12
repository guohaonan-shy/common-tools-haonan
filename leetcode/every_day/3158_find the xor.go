package every_day

func duplicateNumbersXOR(nums []int) int {
	set := make(map[int]int, 0)
	res := 0
	for _, num := range nums {
		set[num] += 1
		if set[num] == 2 {
			res ^= num
		}
	}
	return res
}
