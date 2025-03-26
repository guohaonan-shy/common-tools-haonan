package array_string

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

// 虽然是双层循环，但是遍历的元素数目是N；假设起点固定为0，则经过多次移动之后重新回到位置0，遍历了a圈共遍历了b个元素，满足aN=bK，即aN为N和K的最小公倍数，lcm(N,K)
// b = lcm(N,K)/K；要遍历完N个元素，需要N/b = KN/lcm(N,K) =gcd(N,K)
func rotate_cycle(nums []int, k int) {
	start, count := 0, gcd(len(nums), k)
	for ; start < count; start++ {
		value, pos := nums[start], start
		for ok := true; ok; ok = start != pos {
			next := (pos + k) % len(nums)
			value, nums[next] = nums[next], value
			pos = next
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func rotate_reverse(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func rotateV2(nums []int, k int) {

	k = k % len(nums)

	maxDivider := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	cycle := maxDivider(len(nums), k)

	for start := 0; start < cycle; start++ {
		insertedElement := nums[start]
		cur := start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % len(nums)
			nums[next], insertedElement = insertedElement, nums[next]
			cur = next
		}
	}
	return
}
