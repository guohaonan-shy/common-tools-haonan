package section

import "strconv"

func summaryRanges(nums []int) []string {

	res := make([]string, 0)
	right := 0

	for right < len(nums) {
		left := right
		for right++; right < len(nums) && nums[right-1]+1 == nums[right]; right++ {

		}

		str := strconv.Itoa(nums[left])
		if left < right-1 {
			str += "->" + strconv.Itoa(nums[right-1])
		}
		res = append(res, str)
	}

	return res
}
