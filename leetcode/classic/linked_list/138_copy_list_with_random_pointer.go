package linked_list

type NodeWithRandomPointer struct {
	Val    int
	Next   *NodeWithRandomPointer
	Random *NodeWithRandomPointer
}

func copyRandomList(head *NodeWithRandomPointer) *NodeWithRandomPointer {
	copyHead, mapping := deepcopyList(head, map[*NodeWithRandomPointer]*NodeWithRandomPointer{})
	p1, p2 := head, copyHead
	for p1 != nil && p2 != nil {
		random := mapping[p1.Random]
		p2.Random = random
		p1 = p1.Next
		p2 = p2.Next
	}
	return copyHead
}

func deepcopyList(node *NodeWithRandomPointer, mapping map[*NodeWithRandomPointer]*NodeWithRandomPointer) (*NodeWithRandomPointer, map[*NodeWithRandomPointer]*NodeWithRandomPointer) {
	copyNode := &NodeWithRandomPointer{
		Val: node.Val,
	}
	mapping[node] = copyNode
	if node == nil {
		return nil, mapping
	}

	copyNode.Next, mapping = deepcopyList(node.Next, mapping)
	return copyNode, mapping
}
