package every_day

import "math"

/*
	first, we expect to find the minimum value of a sub-array (a group of continuous elements) => sliding window
*/
func minimumDifference(nums []int, k int) int {
	globalMin := math.MaxInt32

	left, right := 0, 0
	subValue := 0
	for ; right < len(nums); right++ {
		subValue |= nums[right]
		/*
			all elements in the array are positive => as right side of window moves, bitwise result is non-strictly increasing
			if subValue <= k, we can further move window's right side to find the potential minimum
		*/
		if subValue <= k {
			globalMin = min(globalMin, k-subValue)
			continue
		}

		/*
			if not, we record the minimum of the value that subValue is greater then k
		*/
		globalMin = min(globalMin, subValue-k)
		/*
			move left side forward
		*/
		for ; left < right; left++ {
			localVal := 0
			// i<=k => because we want to calculate the next subValue, we have to include the right side
			for i := left + 1; i <= right; i++ {
				localVal |= nums[i]
			}
			subValue = localVal
			globalMin = min(globalMin, abs(subValue-k))
			if subValue <= k {
				break
			}
		}
	}

	return globalMin
}
