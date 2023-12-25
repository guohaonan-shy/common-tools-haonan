package classic

// 类似deque
func rotate(nums []int, k int) {
	//if len(nums) < k {
	//	return
	//}
	pivot := k % len(nums)
	temp := make([]int, len(nums))
	copy(temp, nums)
	for index := pivot + 1; index < len(nums); index++ {
		new_index := (index + k) % len(nums)
		nums[new_index] = temp[index]
	}

	for index := 0; index < pivot+1; index++ {
		new_index := (index + k) % len(nums)
		nums[new_index] = temp[index]
	}

	//tail = (tail + k) % len(nums)
	//
	//temp := make([]int, tail+1)
	//copy(temp, nums[0:tail+1])
	//
	//for i := head; head < len(nums); head++ {
	//	nums[i+head-moves] = value
	//}
	//
	//for i, value := range temp {
	//	nums[i+k] = value
	//}

	return
}
