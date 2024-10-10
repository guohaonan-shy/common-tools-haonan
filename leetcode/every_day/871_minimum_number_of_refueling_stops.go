package every_day

type Position struct {
	Cur   int
	Gas   int
	Stops int
}

/*
	BFS timeout => time complexity is O(n^n)
*/

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	globalMinStops := -1

	queue := make([]*Position, 0)
	queue = append(queue, &Position{
		Cur:   0,
		Gas:   startFuel,
		Stops: 0,
	})

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		farthestPos := cur.Cur + cur.Gas

		if globalMinStops != -1 && cur.Stops >= globalMinStops {
			break
		}

		if cur.Cur+cur.Gas >= target {
			if globalMinStops == -1 {
				globalMinStops = cur.Stops
			} else {
				globalMinStops = min(globalMinStops, cur.Stops)
			}

		}
		for i := len(stations) - 1; i >= 0; i-- {
			station := stations[i]
			stationPos, stationFuel := station[0], station[1]

			if stationPos > farthestPos {
				continue
			}

			if stationPos <= cur.Cur {
				break
			}

			queue = append(queue, &Position{
				Cur:   stationPos,
				Gas:   cur.Gas - (stationPos - cur.Cur) + stationFuel,
				Stops: cur.Stops + 1,
			})
		}
	}
	return globalMinStops
}

/*
Observing the stations, we find that the stations is ordered by miles away from the start point => we consider DP
dp[i] means the farthest miles when stop times == i; eg. the longest miles when stop times is 4 is dependent with stop times == 3

time complexity => O(n^2)
*/
func minRefuelStopDP(target int, startFuel int, stations [][]int) int {
	dp := make([]int, len(stations)+1)
	dp[0] = startFuel

	for i, station := range stations {
		for j := i; j >= 0; j-- {
			if dp[j] >= station[0] {
				dp[j+1] = max(dp[j+1], dp[j]+station[1])
			}
		}
	}

	for i := range dp {
		if dp[i] >= target {
			return i
		}
	}
	return -1
}
