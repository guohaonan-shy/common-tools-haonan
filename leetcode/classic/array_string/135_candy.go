package array_string

// 思路点：分数高的学生获取的糖的数目一定要比相邻的多 => 左右两边两次遍历
func candy(ratings []int) int {
	left, right := make([]int, len(ratings)), make([]int, len(ratings))

	left[0], right[len(ratings)-1] = 1, 1

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
	}

	total_num := 0
	for i := range ratings {
		num := compare(left[i], right[i])
		total_num += num
	}

	return total_num
}

func compare(a, b int) int {

	if a < b {
		return b
	} else {
		return a
	}

}

/*
now, we optimize the space complexity based on time complexity O(n)
*/
func candyV2(ratings []int) int {
	pre, res := 1, 1
	complement := 0
	startOfDecr := 0
	for i := 1; i < len(ratings); i++ {
		// increase => pre+1
		if ratings[i] >= ratings[i-1] {
			startOfDecr = 0
			if ratings[i] == ratings[i-1] {
				pre = 1 // rating equals, this child just need to allocate one candy for basic requirements
			} else {
				pre += 1 // when rating strictly increases, this child needs to allocate for pre+1 candies
			}
			res += pre
			complement = 0
		} else {
			// strictly decrease
			if startOfDecr == 0 {
				startOfDecr = pre // this value is used to check whether we need to complement one candy for the start of decreased child, or the end of increased child
			}
			complement += 1 // complement for previous elements in strictly decreasing order
			res += complement
			if startOfDecr == complement {
				res += 1 // complement 1 for the start of decreased child
				startOfDecr += 1
			}
			pre = 1
		}
	}
	return res
}
