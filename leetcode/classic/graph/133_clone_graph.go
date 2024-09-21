package graph

type Node struct {
	Val       int
	Neighbors []*Node
}

var reached = make(map[*Node]*Node, 0)

func cloneGraph(node *Node) *Node {
	cloneNode, ok := reached[node]
	if ok {
		return cloneNode
	}

	cloneNode = &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	reached[node] = cloneNode

	for _, n := range node.Neighbors {
		cur := n
		cloneNeighbor := cloneGraph(cur)
		cloneNode.Neighbors = append(cloneNode.Neighbors, cloneNeighbor)
	}
	return cloneNode
}
