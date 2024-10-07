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

/*
	I spent a lot of time in how to fix in dynamic programming, but I found there is no difference with the above iteration

	Given the total trips, as days increase from one day to upper limit of days that we can travel, the total trips that we can travel is increasing

	eg. If we take one day, we can take 1 trip;
		If we take three days we can take 5 trip ...

	Therefore, our goal is to find the minimum days that we used to take the give trips. We can use binary search to find the minimum day count to meet the given target.

	Another reason why we consider binary search is that if we specified the day, the trips we can take can be calculated.
	If we can reach the target, we can reduce the days to find the potential minimum.
	If not, we can give more days for exploration.
*/

func minimumTime_standard(time []int, totalTrips int) int64 {
	maxBus := 0
	for _, timeCost := range time {
		maxBus = max(maxBus, timeCost)
	}
	left, right := int64(1), int64(maxBus*totalTrips) // upper limit: the max value of time * total trips

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
