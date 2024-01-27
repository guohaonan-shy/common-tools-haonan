package every_day

type SparseVector struct {
	Content map[int]int // key: index of non-zero, value: non-zero value
}

func Constructor(nums []int) SparseVector {
	content := make(map[int]int, 0)
	for i, num := range nums {
		if num != 0 {
			content[i] = num
		}
	}

	return SparseVector{
		Content: content,
	}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	total := 0
	for idx, val := range vec.Content {
		if num, ok := this.Content[idx]; ok {
			total += val * num
		}
	}
	return total
}
