package classic

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
