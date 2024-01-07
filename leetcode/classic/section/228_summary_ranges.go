package section

import "strconv"

func summaryRanges(nums []int) []string {

	res := make([]string, 0)
	left, right := 0, 0
	for right < len(nums) {
		if right-left != nums[right]-nums[left] {
			var str string
			if right-1 == left {
				str = strconv.Itoa(nums[left])
			} else {
				str = strconv.Itoa(nums[left]) + "->" + strconv.Itoa(nums[right-1])
			}
			res = append(res, str)
			left = right
		} else {
			right++
		}
	}

	if right-1 == left {
		res = append(res, strconv.Itoa(nums[left]))
	} else {
		res = append(res, strconv.Itoa(nums[left])+"->"+strconv.Itoa(nums[right-1]))
	}

	return res
}
