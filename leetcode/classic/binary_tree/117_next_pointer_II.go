package binary_tree

type Node struct {
	Val         int
	Left, Right *Node
	Next        *Node
}

var nextStart *Node
var last *Node

func connect(root *Node) *Node {

	start := root
	for start != nil {

		for cur := start; cur != nil; cur = cur.Next {
			connectHandle(cur.Left)
			connectHandle(cur.Right)
		}
		start = nextStart
		nextStart = nil
		last = nil
	}
	return root
}

func connectHandle(node *Node) {
	if node == nil {
		return
	}

	if nextStart == nil {
		nextStart = node
	}

	if last != nil {
		last.Next = node
	}
	last = node
}
