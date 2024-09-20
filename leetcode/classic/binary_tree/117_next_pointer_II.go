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

func connectV2(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := make([]*Node, 0)
	nextQueue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		for i := 0; i < len(queue); i++ {
			if i == len(queue)-1 {
				queue[i].Next = nil
			} else {
				queue[i].Next = queue[i+1]
			}

			if queue[i].Left != nil {
				nextQueue = append(nextQueue, queue[i].Left)
			}

			if queue[i].Right != nil {
				nextQueue = append(nextQueue, queue[i].Right)
			}
		}

		queue = nextQueue
		nextQueue = []*Node{}
	}
	return root
}
