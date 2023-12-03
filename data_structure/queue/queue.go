package queue

type Queue[T any] struct {
	front, end       int
	length, capacity int
	data             []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		front:    0,
		end:      -1,
		capacity: 4,
		length:   0,
		data:     make([]T, 4, 4),
	}
}

func (queue *Queue[T]) Insert(value T) {
	if queue.length == queue.capacity { // expand and copy
		newData := make([]T, queue.capacity*2, queue.capacity*2)
		cur := queue.front
		for i := 0; i < queue.length; i++ {
			newData[i] = queue.data[cur]
			if cur == queue.length-1 {
				cur = 0
			} else {
				cur++
			}
		}

		newData[queue.length] = value

		queue.front, queue.end = 0, queue.length+1
		queue.capacity = queue.capacity * 2
		queue.data = newData
		queue.length = len(newData)
	} else {
		var end int
		if queue.end == queue.capacity-1 {
			end = (queue.end + 1) % queue.capacity
		} else {
			end = queue.end + 1
		}
		queue.length += 1
		queue.data[end] = value
		queue.end = end
	}
}

func (queue *Queue[T]) Pop() T {
	if queue.length == 0 {
		return *new(T)
	}

	var res T
	res = queue.data[queue.front]
	temp := new(T)
	queue.data[queue.front] = *temp
	if queue.front < queue.end {
		queue.front++
	} else if queue.front > queue.end {
		queue.front = (queue.front + 1) % queue.capacity
	}
	queue.length--
	return res
}
