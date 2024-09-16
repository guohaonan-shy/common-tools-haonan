package hash_table

/*
	In extreme cases, this solution will degrade to O(N^2), like x, x+1, x+2...x+y.
	However, x+1 ~ x+y, these elements who have val - 1 in sequence is not what we want, not the longest. => if val - 1 exists, we skip.
	Only num whose value-1 doesn't exist in sequence is what we want to calculate and find.
*/
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
