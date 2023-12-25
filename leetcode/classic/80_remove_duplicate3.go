package classic

func removeDuplicates_atMostTwice(nums []int) int {
	duplicates, slow := 1, 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast-1] == nums[fast] {
			if duplicates < 2 {
				nums[slow] = nums[fast]
				duplicates += 1
				slow++
			}
		} else {
			nums[slow] = nums[fast]
			duplicates = 1
			slow++
		}
	}
	return slow
}

func removeDuplicates_atMostTwice_standard(nums []int) int {
	if len(nums) <= 2 {
		return 2
	}
	slow := 2
	for fast := 2; fast < len(nums); fast++ { // 1，1，1，2，2，3
		if nums[slow-2] != nums[fast] {
			// 这行比较难理解，用例子来说明，本意是当前遍历的元素值最多可能和当前待插入位置的上上个元素相等，即保证重复元素最多出现“两次”
			// eg1: [2,2,2,2], 初始slow = 2，当fast = 2 or 3时，均满足nums[fast] == nums[slow-2]，即本次fast位置的元素出现次数大于2，需要往后遍历寻找到满足条件的元素
			// eg2: [2,3,3,3], 初始slow = 2，当fast == 2时， nums[fast] != nums[slow-2]，即fast这个位置的元素满足最多出现两次的条件，插入slow
			// nums[slow-1] 和 nums[slow-2]可能一样也可能不一样，如果一样，即该判断条件可解读为slow这个位置是否能插入fast元素，而要插入则需要满足只出现两次；
			// 如果不一样，则fast肯定满足该条件；fast插入slow的位置，最多只与slow-1位置的元素相同
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
