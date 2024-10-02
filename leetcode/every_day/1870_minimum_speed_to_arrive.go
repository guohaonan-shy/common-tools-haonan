package every_day

import "math"

/*
	the most extreme case: we need to take one train to run the longest distance in the shortest time => we run 10^5 distance in 1 hour => the max minimum speed is 10^5
*/

func minSpeedOnTime(dist []int, hour float64) int {
	//minSpeed, maxSpeed := 1, math.MaxInt32
	minSpeed, maxSpeed := 1, 10000000 // because hour has at most two digit, the shortest drive period is 10^5/0.01 = 10^7
	globalMinSpeed := maxSpeed
	for minSpeed <= maxSpeed {

		mid := (minSpeed + maxSpeed) / 2

		localHour := 0.0

		for i := range dist {
			precisionHour := float64(dist[i]) / float64(mid)
			if localHour+precisionHour > hour {
				localHour += precisionHour
				break
			}

			if i == len(dist)-1 {
				localHour += precisionHour
			} else {
				localHour += math.Ceil(precisionHour)
			}
		}
		if localHour <= hour {
			maxSpeed = mid - 1
			globalMinSpeed = min(globalMinSpeed, mid)
		} else {
			minSpeed = mid + 1
		}
	}
	// be compatible with 32-bit and 64-bit both, because we use maximum of int32 as our right pointer
	//if minSpeed < 0 || minSpeed == math.MaxInt32+1 {
	//	globalMinSpeed = -1
	//}

	if minSpeed == 10000001 {
		globalMinSpeed = -1
	}
	return globalMinSpeed
}
