package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Heap(t *testing.T) {
	t.Run("heap_case1", func(t *testing.T) {
		heap1 := NewEmptyHeap[int]()
		heap1.Insert(1)
		heap1.Insert(2)
		heap1.Insert(4)
		assert.Equal(t, []int{1, 2, 4}, heap1.data)
		heap1.Insert(3)
		assert.Equal(t, []int{1, 2, 4, 3}, heap1.data)
		heap1.Insert(7)
		heap1.Insert(10)
		heap1.Insert(8)
		heap1.Insert(20)
		heap1.Insert(21)
		assert.Equal(t, []int{1, 2, 4, 3, 7, 10, 8, 20, 21}, heap1.data)
		heap1.Insert(6)
		assert.Equal(t, []int{1, 2, 4, 3, 6, 10, 8, 20, 21, 7}, heap1.data)
	})

	t.Run("heap_case2", func(t *testing.T) {
		heap2 := NewEmptyHeap[int]()
		heap2.Insert(5)
		heap2.Insert(7)
		heap2.Insert(8)
		heap2.Insert(2)
		assert.Equal(t, []int{2, 5, 8, 7}, heap2.data)
		// pop
		element := heap2.Pop()
		assert.Equal(t, 2, element)
		assert.Equal(t, []int{5, 7, 8}, heap2.data)
		// insert
		heap2.Insert(1)
		assert.Equal(t, []int{1, 5, 8, 7}, heap2.data)
		heap2.Insert(2)
		assert.Equal(t, []int{1, 2, 8, 7, 5}, heap2.data)
		heap2.Insert(3)
		assert.Equal(t, []int{1, 2, 3, 7, 5, 8}, heap2.data)
		heap2.Insert(4)
		assert.Equal(t, []int{1, 2, 3, 7, 5, 8, 4}, heap2.data)
		heap2.Insert(7)
		assert.Equal(t, []int{1, 2, 3, 7, 5, 8, 4, 7}, heap2.data)
	})

	t.Run("heap_case3", func(t *testing.T) {
		heap3 := Heapify(false, []int{6, 5, 4, 3, 2, 1})
		assert.Equal(t, []int{1, 3, 2, 6, 4, 5}, heap3.data)
	})

	t.Run("heap_case4", func(t *testing.T) {
		heap4 := FasterHeapify(false, []int{6, 5, 4, 3, 2, 1})
		assert.Equal(t, []int{1, 2, 4, 3, 5, 6}, heap4.data)
	})
}
