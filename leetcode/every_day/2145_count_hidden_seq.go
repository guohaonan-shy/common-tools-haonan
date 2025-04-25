package every_day

func numberOfArrays(differences []int, lower int, upper int) int {
	prev := 0
	left, right := lower, upper
	for i := 0; i < len(differences); i++ {

		if left > right {
			break
		}

		prev = differences[i] + prev

		right = min(right, upper-prev)
		left = max(left, lower-prev)
	}

	if left > right {
		return 0
	}
	return right - left + 1
}
