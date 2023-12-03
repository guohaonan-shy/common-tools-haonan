package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Queue(t *testing.T) {
	t.Run("queue_case1", func(t *testing.T) {
		queue1 := NewQueue[int]()
		queue1.Insert(1)
		queue1.Insert(2)
		queue1.Insert(3)
		queue1.Insert(4)
		assert.Equal(t, []int{1, 2, 3, 4}, queue1.data)
		queue1.Insert(5)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 0, 0, 0}, queue1.data)
	})

	t.Run("queue_case2", func(t *testing.T) {
		queue1 := NewQueue[int]()
		queue1.Insert(1)
		queue1.Insert(2)
		queue1.Insert(3)
		queue1.Insert(4)
		assert.Equal(t, []int{1, 2, 3, 4}, queue1.data)
		// pop from front
		element := queue1.Pop()
		assert.Equal(t, 1, element)
		assert.Equal(t, []int{0, 2, 3, 4}, queue1.data)
		assert.Equal(t, 1, queue1.front)
		// insert in pop position
		queue1.Insert(5)
		assert.Equal(t, []int{5, 2, 3, 4}, queue1.data)
		assert.Equal(t, 0, queue1.end)
		// trigger extend
		queue1.Insert(6)
		assert.Equal(t, []int{2, 3, 4, 5, 6, 0, 0, 0}, queue1.data)
	})

	t.Run("queue_case3", func(t *testing.T) {
		queue3 := NewQueue[int]()
		element := queue3.Pop()
		assert.Equal(t, 0, element)
		//
		queue3.Insert(1)
		queue3.Insert(2)
		queue3.Pop()
		assert.Equal(t, []int{0, 2, 0, 0}, queue3.data)
		assert.Equal(t, 1, queue3.front)
		assert.Equal(t, 1, queue3.end)
	})
}
