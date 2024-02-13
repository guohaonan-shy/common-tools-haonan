package array_string

func sortColors_Nmemory(nums []int) {
	colorDict := make(map[int]int, 0)
	for _, color := range nums {
		colorDict[color] += 1
	}

	idx := 0
	zero := colorDict[0]
	for ; idx < zero; idx++ {
		nums[idx] = 0
	}

	first := colorDict[1]
	for ; idx < zero+first; idx++ {
		nums[idx] = 1
	}

	second := colorDict[2]
	for ; idx < zero+first+second; idx++ {
		nums[idx] = 2
	}
	return
}

func sortColors_twoIter(nums []int) {
	// two iteration
	left, right := 0, len(nums)-1

	// right 待指，下一个2插入的位置
	for i := 0; i < right; i++ {
		for right > 0 && nums[right] == 2 {
			right--
		}
		// 此时得到了下一个2待插入的位置
		// 如果right超出边界，或者此时遍历指针已经在待插入2的右边，中断
		if right < 0 || i > right {
			break
		}

		// 否则交换位置
		if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}
	}

	// i一定是大于等于left的，left只有在发生交换时，才会递增
	for i := 0; i <= right; i++ {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
	return
}

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	for i := 0; i < right; i++ {
		for ; i < right && nums[i] == 2; right-- {
			nums[i], nums[right] = nums[right], nums[i]
		}

		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
}
