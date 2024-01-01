package array_string

// nums是有序的，即对任意的i,j, i<=k<=j, 如果nums[i]==nums[j], 则一定nums[i] == nums[k] == nums[j]
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	// 1,1,2
	// 0,1,2,1,1,2,2,3,3,4

	slow := 1 // 下一个递增元素插入位置
	//prev := nums[0]
	for fast := 1; fast < len(nums); fast++ {
		//if nums[i] <= prev {
		//	continue
		//} else {
		//	nums[left] = nums[i]
		//	prev = nums[i]
		//	left++
		//}
		if nums[fast-1] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}

	}
	return slow
}
