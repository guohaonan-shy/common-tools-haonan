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
	return cur.Prev
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
	// set dummy in the head and tail of the linked-list
	dummyHead, dummyTail := &LinkedNode{Key: -1, Val: -1}, &LinkedNode{Key: -1, Val: -1}
	dummyHead.Next = dummyTail
	dummyTail.Prev = dummyHead

	return &LRUCacheIns{
		Capacity:   capacity,
		data:       make(map[int]int, capacity),
		linkedList: dummyHead,
	}
}

/*
Get() needs to move the retrieved kv node into the head of linked list
*/
func (lru *LRUCacheIns) Get(key int) int {

	val, ok := lru.data[key]
	if !ok {
		return -1
	}

	// move the retrieved element into the head
	target := lru.linkedList.search(key)
	prev, next := target.Prev, target.Next
	// because of the existence of two dummy nodes, target.prev && target.next must be non-nil
	// 1. delete the connection
	prev.Next = next
	next.Prev = prev

	// 2. insert 'target' into the head
	oldHead := lru.linkedList.Next
	target.Next = oldHead
	oldHead.Prev = target

	lru.linkedList.Next = target
	target.Prev = lru.linkedList
	return val
}

/*
	Put() needs to insert the node in the head of linked-list;
	And if the capacity is full, we have to remove the tail of this linked-list
*/

func (lru *LRUCacheIns) Put(key, val int) {
	// no elements:
	// 1. init
	// 2. after a series of operations, there is no data currently

	// if key is in the data, we don't need to remove the tail node => just insert to the head
	if _, ok := lru.data[key]; ok {
		target := lru.linkedList.search(key)
		target.Val = val
		prev, next := target.Prev, target.Next

		prev.Next = next
		next.Prev = prev

		oldHead := lru.linkedList.Next
		lru.linkedList.Next = target
		target.Next = oldHead
		oldHead.Prev = target
		target.Prev = lru.linkedList

		lru.data[key] = val
		return
	}

	// insert a kv that doesn't in the list
	// insert element in the head of list,
	// but there is a corner case: we need to remove the tail if reach the capacity

	if lru.Capacity == len(lru.data) {
		tail := lru.linkedList.tail()

		prev, next := tail.Prev, tail.Next
		prev.Next = next
		next.Prev = prev

		delete(lru.data, tail.Key)
	}

	insertedNode := &LinkedNode{
		Key: key,
		Val: val,
	}
	oldHead := lru.linkedList.Next

	lru.linkedList.Next = insertedNode
	insertedNode.Next = oldHead
	oldHead.Prev = insertedNode
	insertedNode.Prev = lru.linkedList

	lru.data[key] = val
	return
}
