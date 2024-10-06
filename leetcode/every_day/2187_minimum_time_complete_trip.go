package every_day

/*
timeout
*/
func minimumTime(time []int, totalTrips int) int64 {
	// according to the constraints, the maximum expected day is 10^7
	cur := 1
	for ; cur <= 10000000; cur++ {
		local := 0
		for _, timeCost := range time {
			local += cur / timeCost
			if local >= totalTrips {
				return int64(cur)
			}
		}
	}
	return int64(cur - 1)
}

func minimumTime_standard(time []int, totalTrips int) int64 {
	maxBus := 0
	for _, timeCost := range time {
		maxBus = max(maxBus, timeCost)
	}
	left, right := int64(1), int64(maxBus*totalTrips)

	for left < right {
		mid := (left + right) / 2
		var cnt int64 = 0
		for _, timeCost := range time {
			cnt += mid / int64(timeCost)
			if cnt >= int64(totalTrips) {
				break
			}
		}
		if cnt < int64(totalTrips) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
