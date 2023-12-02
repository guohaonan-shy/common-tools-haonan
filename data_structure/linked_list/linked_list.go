package linked_list

type LinkedNode struct {
	item int
	next *LinkedNode
}

func NewLinkedNode(value int) *LinkedNode {
	return &LinkedNode{
		item: value,
	}
}

type LinkedList struct {
	head, tail *LinkedNode
	length     int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Insert(i int, value int) {

	node := NewLinkedNode(value)

	if i == 0 { // insert into head
		nextNode := ll.head
		node.next = nextNode
		ll.head = node
		if ll.length == 0 { // insert into an empty list
			ll.tail = node
		}
	} else if i > 0 && i < ll.length {
		cur := ll.head
		for j := 0; j < i-1; j++ {
			cur = cur.next
		}

		nextNode := cur.next
		cur.next = node
		node.next = nextNode
	} else {
		cur := ll.tail
		cur.next = node
		ll.tail = node
	}
	ll.length++
}

func (ll *LinkedList) Remove(i int) {
	if i == 0 {
		head := ll.head
		temp := head.next
		head.next = nil
		ll.head = temp
	} else if i == ll.length-1 {
		cur := ll.head
		for j := 0; j < ll.length-2; j++ {
			cur = cur.next
		}
		cur.next = nil
		ll.tail = cur
	} else {
		cur := ll.head
		for j := 0; j < i-1; j++ {
			cur = cur.next
		}
		temp := cur.next
		cur.next = cur.next.next
		temp.next = nil
	}
	ll.length--
}
