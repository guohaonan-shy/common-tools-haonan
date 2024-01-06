package hash_table

func longestConsecutive(nums []int) int {
	dict := make(map[int]bool, 0)
	maxLength := 0
	for _, value := range nums { // 集合
		dict[value] = true
	}

	// 双层循环，但是时间复杂度是N
	for key := range dict {
		if !dict[key-1] { // 连续序列的开头一定是不连续的，即element-1不存在
			temp := key
			length := 0
			for ; dict[temp]; temp++ {
				length++
			}
			maxLength = max(maxLength, length)
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
