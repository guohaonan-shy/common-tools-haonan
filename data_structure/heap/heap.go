package heap

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |

		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |

		~float32 | ~float64 |

		~string
}

type Heap[T ordered] struct {
	isMax bool
	data  []T
}

func NewEmptyHeap[T ordered]() *Heap[T] {
	return &Heap[T]{
		isMax: false, // default minRoot
		data:  make([]T, 0),
	}
}

func Heapify[T ordered](isMax bool, values []T) *Heap[T] {
	heap := &Heap[T]{
		isMax: isMax,
	}

	for _, value := range values {
		heap.Insert(value)
	}
	return heap
}

func FasterHeapify[T ordered](isMax bool, values []T) *Heap[T] {
	heap := &Heap[T]{
		isMax: isMax,
		data:  values,
	}

	last := len(values) - 1
	parent := (last - 1) / 2
	for i := parent; i >= 0; i-- {
		cur := i
		for cur<<1+1 < len(heap.data) || cur<<1+2 < len(heap.data) {
			var (
				left, right int
			)

			if cur<<1+2 < len(heap.data) {
				right = cur<<1 + 2
			}
			left = cur<<1 + 1

			if heap.isLeft(cur, left, right) {
				temp := heap.data[left]
				heap.data[left] = heap.data[cur]
				heap.data[cur] = temp
				cur = left
			} else {
				temp := heap.data[right]
				heap.data[right] = heap.data[cur]
				heap.data[cur] = temp
				cur = right
			}
		}
	}
	return heap
}

func (heap *Heap[T]) Insert(value T) {
	heap.data = append(heap.data, value)
	cur := len(heap.data) - 1
	for cur > 0 {
		parent := (cur - 1) / 2

		if heap.isMax {
			if heap.data[parent] < heap.data[cur] {
				temp := heap.data[parent]
				heap.data[parent] = heap.data[cur]
				heap.data[cur] = temp
			}
		} else {
			if heap.data[parent] > heap.data[cur] {
				temp := heap.data[parent]
				heap.data[parent] = heap.data[cur]
				heap.data[cur] = temp
			}
		}

		cur = parent
	}
}

func (heap *Heap[T]) Pop() T {
	result := heap.data[0]
	heap.data[0] = heap.data[len(heap.data)-1]
	heap.data = heap.data[0 : len(heap.data)-1]

	cur := 0
	for cur<<1+1 < len(heap.data) || cur<<1+2 < len(heap.data) {
		var (
			left, right int
		)

		if cur<<1+2 < len(heap.data) {
			right = cur<<1 + 2
		}
		left = cur<<1 + 1

		if heap.isLeft(cur, left, right) {
			temp := heap.data[left]
			heap.data[left] = heap.data[cur]
			heap.data[cur] = temp
			cur = left
		} else {
			temp := heap.data[right]
			heap.data[right] = heap.data[cur]
			heap.data[cur] = temp
			cur = right
		}
	}

	return result
}

func (heap *Heap[T]) isLeft(cur, left, right int) bool {
	res := false
	if right == 0 {
		if heap.isMax {
			if heap.data[left] > heap.data[cur] {
				res = true
			}
		} else {
			if heap.data[left] < heap.data[cur] {
				res = true
			}
		}
	} else {
		if heap.isMax {
			if heap.data[left] > heap.data[cur] && heap.data[left] > heap.data[right] {
				res = true
			}
		} else {
			if heap.data[left] < heap.data[cur] && heap.data[left] < heap.data[right] {
				res = true
			}
		}
	}
	return res
}
