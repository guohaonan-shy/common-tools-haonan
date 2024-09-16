package linked_list

type LinkedNode struct {
	Key, Val int
	Next     *LinkedNode
	Prev     *LinkedNode
}

func NewLinkedNode(key, val int) *LinkedNode {
	return &LinkedNode{
		Key: key,
		Val: val,
	}
}

func (node *LinkedNode) tail() *LinkedNode {
	cur := node
	for ; cur.Next != nil; cur = cur.Next {
	}
	return cur
}

func (node *LinkedNode) search(key int) *LinkedNode {
	cur := node
	for ; cur != nil && cur.Key != key; cur = cur.Next {

	}
	return cur
}

type LRUCacheIns struct {
	Capacity   int
	data       map[int]int
	linkedList *LinkedNode
}

func Constructor(capacity int) *LRUCacheIns {
	return &LRUCacheIns{
		Capacity:   capacity,
		data:       make(map[int]int, capacity),
		linkedList: nil,
	}
}

func (lru *LRUCacheIns) Get(key int) int {

	val, ok := lru.data[key]
	if !ok {
		return -1
	}

	// move the retrieved element into the head
	target := lru.linkedList.search(key)
	prev, next := target.Prev, target.Next
	// 1. delete the connection
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}
	// 2. change the connection of target to head
	if len(lru.data) == 1 {
		return val
	}

	second := lru.linkedList
	target.Prev = second.Prev
	target.Next = second
	second.Prev = target
	lru.linkedList = target
	return val
}

func (lru *LRUCacheIns) Put(key, val int) {
	// no elements:
	// 1. init
	// 2. after a series of operations, there is no data currently.
	if lru.linkedList == nil {
		lru.linkedList = NewLinkedNode(key, val)
		lru.data[key] = val
		return
	}

	// if key is in the data, we don't need to remove the tail node
	if _, ok := lru.data[key]; ok {
		target := lru.linkedList.search(key)
		target.Val = val
		prev, next := target.Prev, target.Next
		if prev != nil {
			prev.Next = nil
		}
		if next != nil {
			next.Prev = nil
		}

		second := lru.linkedList
		if second != target { // not head, avoid self reference
			target.Next = second
			second.Prev = target
			lru.linkedList = target
		}
		lru.data[key] = val
		return
	}

	// insert a kv that doesn't in the list
	// insert element in the head of list,
	// but there is a corner case: we need to remove the tail if reach the capacity
	currentLength := len(lru.data)
	if currentLength == lru.Capacity { // need to remove tail
		//
		tail := lru.linkedList.tail() // previous node become the tail, o(n)
		delete(lru.data, tail.Key)
		preTail := tail.Prev
		if preTail != nil {
			preTail.Next = nil
		} else {
			// capacity == 1
			lru.linkedList = nil
		}
		tail.Prev = nil
	}
	insertedNode := NewLinkedNode(key, val)
	secondNode := lru.linkedList
	insertedNode.Next = secondNode
	if secondNode != nil {
		secondNode.Prev = insertedNode
	}
	lru.data[key] = val
	lru.linkedList = insertedNode
	return
}
