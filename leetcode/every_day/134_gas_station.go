package every_day

func canCompleteCircuit(gas []int, cost []int) int {

	totalGas, totalCost := sum(gas), sum(cost)
	if totalGas < totalCost {
		return -1
	}
	// the below part we have confirmed this case has solution, we just iterate and find it
	curGas := 0
	startPoint := 0
	/*
		current gas volume in tank:
		1. if we can go to the next location, we accumulate gas in the tank
		2. if we can't, we cannot start from previous location. So start from next point and reset the gas tank to 0
	*/
	for i := range gas {
		localGas, nextCost := gas[i], cost[i]
		if curGas+localGas >= nextCost {
			curGas += localGas - nextCost
		} else {
			curGas = 0
			startPoint = i + 1
		}
	}
	return startPoint
}

func sum(list []int) int {
	res := 0
	for i := range list {
		res += list[i]
	}
	return res
}
