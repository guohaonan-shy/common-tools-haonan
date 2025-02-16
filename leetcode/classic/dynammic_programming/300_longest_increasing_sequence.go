package dynammic_programming

func longestIncreasingSequence(nums []int) int {

	dp := make([]int, len(nums))
	// dp[i] is the maximum length of strictly increasing subsequence that ends with index i
	// if maximum subsequence that doesn't end with last element, 'dp[len(nums)-1]' is not what we want.
	// maximum among dp[0:len(nums)] is what we want
	maxLength := 0
	for end := 0; end < len(nums); end++ {
		localMax := 1
		for start := 0; start < end; start++ {
			if nums[start] < nums[end] {
				localMax = max(localMax, dp[start]+1)
			}
		}
		dp[end] = localMax
		maxLength = max(maxLength, localMax)
	}
	return maxLength
}

func longestIncreasingSequence_binarySearch(nums []int) int {
	list := make([]int, len(nums)+1) // list[i] means that the minimum of last element in the increasing sequence, which the length is i

	insetedIdx := 1 // this is the position that our target element, 'num', will be inserted in
	for _, num := range nums {
		left, right := 1, insetedIdx
		for left < right {

			mid := (left + right) / 2

			if list[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		list[left] = num

		if left == insetedIdx {
			insetedIdx++
		}
	}
	return insetedIdx - 1
}
