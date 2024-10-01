package every_day

func minimumTravelCost(days []int, costs []int) int {
	lastDay := days[len(days)-1] // the latest day that we have to go out for a trip
	dp := make([]int, lastDay+1) // the minimum cost of no.i day
	dayIdx := 0
	for curDay := 1; curDay <= lastDay; curDay++ {
		nextTravelDay := days[dayIdx]
		/*
			 check if today is the date that we must go out
				1. if not, we can buy nothing;
				2. if yes, we can buy a ticket that cost least(we're greedy) to cover today into our travel plan:
					1. we buy a 1-day plan based on i-1 day's minimum cost.
					2. we buy a 7-day plan based on i-7 day's minimum cost(if i-7 < 0, we starts from 0).
					3. we buy a 30-day plan based on i-30 day's minimum cost(if i-30 < 0, we starts from 0).
		*/

		if curDay < nextTravelDay {
			dp[curDay] = dp[curDay-1]
			continue
		}

		dp[curDay] = dp[curDay-1] + min(costs[0], min(costs[1], costs[2])) // buy the cheapest ticket(greedy)
		dayIdx++                                                           // next necessary date

		if curDay-7 >= 0 {
			dp[curDay] = min(dp[curDay], dp[curDay-7]+min(costs[1], costs[2])) // if we buy 7 or 30-day plan 7 days ago
		} else {
			dp[curDay] = min(dp[curDay], dp[0]+min(costs[1], costs[2])) // choose the cheapest between 7-day plan and 30-day plan
		}

		if curDay-30 >= 0 {
			dp[curDay] = min(dp[curDay], dp[curDay-30]+costs[2])
		} else {
			dp[curDay] = min(dp[curDay], dp[0]+costs[2])
		}

	}
	return dp[lastDay]
}
