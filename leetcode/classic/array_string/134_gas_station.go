package array_string

func canCompleteCircuit(gas []int, cost []int) int {
	temp := make([]int, len(gas)) // temp[i]表示i作为起点，抵达i+1时的剩余油量
	total := 0
	for i := range gas {
		temp[i] = gas[i] - cost[i]
		total += temp[i]
	}

	if total < 0 { // 总油量小于总消耗
		return -1
	}

	// total >=0 保证肯定有解
	curSum := 0
	ans := 0
	for i := range temp { // 从第一个点为起点，开始前进
		curSum += temp[i]
		if curSum < 0 {
			// 当抵达某个点之后，剩余油量为负值；因为累计开始，肯定是从剩余为正数的第一个站点开始的，所以从起点x到当前点x之间的所有点都不会做为起点（减去x这个正数，值更是负的）
			// 则下一个起点为i之后的第一个正值
			// 这一步还会跳过temp[j] < 0 作为起点的情况，不需要单独维护一个是正的数组，
			curSum = 0
			ans = i + 1
		}
	}
	return ans
}
