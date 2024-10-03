package every_day

import "math"

/*
the method of dfs is timeout => dynamic programming
*/
func minCost(maxTime int, edges [][]int, passingFees []int) int {
	// analysis the graph
	adjacenyMatrix := make([][]int, len(passingFees))
	for i := range adjacenyMatrix {
		adjacenyMatrix[i] = make([]int, len(passingFees))
	}

	for _, edge := range edges {
		left, right := edge[0], edge[1]
		weight := edge[2]

		if adjacenyMatrix[left][right] == 0 {
			adjacenyMatrix[left][right] = weight
		} else {
			adjacenyMatrix[left][right] = min(adjacenyMatrix[left][right], weight)
		}

		if adjacenyMatrix[right][left] == 0 {
			adjacenyMatrix[right][left] = weight
		} else {
			adjacenyMatrix[right][left] = min(adjacenyMatrix[right][left], weight)
		}
	}
	adjacenyList := make(map[int][]int, 0)
	for start := range adjacenyMatrix {
		adjacenyList[start] = []int{}
		for end := range adjacenyMatrix[start] {
			if adjacenyMatrix[start][end] > 0 {
				adjacenyList[start] = append(adjacenyList[start], end)
			}
		}
	}

	var (
		dfs           func(idx int, curTime int, cost int)
		globalMinCost = -1
		reached       = make(map[int]struct{}, 0)
	)

	dfs = func(idx int, curTime int, cost int) {
		if curTime > maxTime {
			return
		}

		if idx == len(passingFees)-1 {
			if globalMinCost == -1 || (cost < globalMinCost) {
				globalMinCost = cost
			}
			return
		}

		for _, next := range adjacenyList[idx] {
			if _, ok := reached[next]; ok {
				continue
			}
			reached[idx] = struct{}{}
			dfs(next, curTime+adjacenyMatrix[idx][next], cost+passingFees[next])
			delete(reached, idx)
		}
		return
	}
	dfs(0, 0, passingFees[0])
	return globalMinCost
}

func minCostDP(maxTime int, edges [][]int, passingFees []int) int {
	// analysis the graph
	adjacenyMatrix := make([][]int, len(passingFees))
	for i := range adjacenyMatrix {
		adjacenyMatrix[i] = make([]int, len(passingFees))
	}

	for _, edge := range edges {
		left, right := edge[0], edge[1]
		weight := edge[2]

		if adjacenyMatrix[left][right] == 0 {
			adjacenyMatrix[left][right] = weight
		} else {
			adjacenyMatrix[left][right] = min(adjacenyMatrix[left][right], weight)
		}

		if adjacenyMatrix[right][left] == 0 {
			adjacenyMatrix[right][left] = weight
		} else {
			adjacenyMatrix[right][left] = min(adjacenyMatrix[right][left], weight)
		}
	}
	adjacenyList := make(map[int][]int, 0)
	for start := range adjacenyMatrix {
		adjacenyList[start] = []int{}
		for end := range adjacenyMatrix[start] {
			if adjacenyMatrix[start][end] > 0 {
				adjacenyList[start] = append(adjacenyList[start], end)
			}
		}
	}

	dp := make([][]int, len(passingFees))
	for i := range dp {
		dp[i] = make([]int, maxTime+1)
	}

	// inititalization
	nexts := adjacenyList[0]

	for _, next := range nexts {
		time := adjacenyMatrix[0][next]
		if time <= maxTime {
			dp[next][0+time] = passingFees[0] + passingFees[next]
		}
	}

	for i := 1; i < maxTime; i++ {
		for point := 0; point < len(passingFees)-1; point++ {
			if dp[point][i] == 0 {
				continue
			}

			nextDes := adjacenyList[point]
			for _, next := range nextDes {
				time := i + adjacenyMatrix[point][next]
				if time > maxTime {
					continue
				}

				if dp[next][time] == 0 {
					dp[next][time] = dp[point][i] + passingFees[next]
				} else {
					dp[next][time] = min(dp[next][time], dp[point][i]+passingFees[next])
				}

			}
		}
	}
	globalMinCost := math.MaxInt32
	for _, cost := range dp[len(passingFees)-1] {
		if cost == 0 {
			continue
		}
		globalMinCost = min(globalMinCost, cost)
	}
	if globalMinCost == math.MaxInt32 {
		globalMinCost = -1
	}
	return globalMinCost
}
