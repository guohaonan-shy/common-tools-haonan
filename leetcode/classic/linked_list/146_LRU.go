package linked_list

type DoubleLinkedNode struct {
	Key  int
	Val  int
	Prev *DoubleLinkedNode
	Next *DoubleLinkedNode
}

type LRUCache struct {
	Capacity int
	Length   int
	Head     *DoubleLinkedNode
	Tail     *DoubleLinkedNode
	Dict     map[int]*DoubleLinkedNode
}

func Constructor(capacity int) LRUCache {
	dummy := &DoubleLinkedNode{
		Key:  -1,
		Val:  -1,
		Prev: nil,
		Next: nil,
	}
	return LRUCache{
		Capacity: capacity,
		Length:   0,
		Dict:     make(map[int]*DoubleLinkedNode, capacity),
		Head:     dummy,
		Tail:     dummy,
	}
}

func (this *LRUCache) Get(key int) int {
	if target, ok := this.Dict[key]; !ok {
		return -1
	} else {
		targetVal := target.Val
		// update queue
		if this.Head.Next == target { // get head, no need for move
			return targetVal
		}

		// 至少两个节点
		target.Prev.Next = target.Next

		if target == this.Tail { // update tail，更新tail位置
			this.Tail = this.Tail.Prev
		} else {
			target.Next.Prev = target.Prev
		}
		target.Prev = this.Head
		target.Next = this.Head.Next

		this.Head.Next.Prev = target
		this.Head.Next = target
		return targetVal
	}
}

func (this *LRUCache) Put(key int, value int) {
	if target, ok := this.Dict[key]; !ok {
		node := &DoubleLinkedNode{
			Key:  key,
			Val:  value,
			Prev: this.Head,
			Next: this.Head.Next,
		}

		// 非第一次插入和第一次插入
		if this.Length > 0 { // this.head.next != nil
			this.Head.Next.Prev = node
		} else {
			this.Tail = node
		}
		this.Head.Next = node
		this.Dict[key] = node
		// 插入完成

		// 考虑evict
		if this.Length == this.Capacity {
			deleteKey := this.Tail.Key
			this.Tail = this.Tail.Prev
			this.Tail.Next = nil
			delete(this.Dict, deleteKey)
		} else {
			this.Length++
		}
	} else { // 三种情况，更新队首，队尾，对中间
		target.Val = value
		if this.Head.Next == target { // update head
			return
		}

		// 至少两个节点
		target.Prev.Next = target.Next

		if target == this.Tail { // update tail，更新tail位置
			this.Tail = this.Tail.Prev
		} else {
			target.Next.Prev = target.Prev
		}
		target.Prev = this.Head
		target.Next = this.Head.Next

		this.Head.Next.Prev = target
		this.Head.Next = target
	}
}
