package sort

import (
	"github.com/v2pro/plz/test/testify/assert"
	"math/rand"
	"testing"
)

func Test_BubbleSort(t *testing.T) {
	test_case := []int{5, 4, 3, 2, 1}
	BubbleSort(test_case)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, test_case)
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {

		test_case := make([]int, 0, 1000)
		for i := 0; i < 1000; i++ {
			test_case = append(test_case, rand.Intn(1000))
		}

		b.Run("bubble_sort", func(b *testing.B) {
			BubbleSort(test_case)
		})
		b.Run("bubble_sort_early_terminate", func(b *testing.B) {
			BubbleSort_EarlyTerminate(test_case)
		})
	}
}
