package framework

type TreeNode struct {
	path     string
	pattern  string
	children []*TreeNode
}

func (node *TreeNode) insert(path string, parts []string, layer int) {
	// last part in path, no child
	if layer == len(parts) {
		node.path = path
		node.pattern = parts[layer-1]
		return
	}

	// normal
	child := node.matchChild(parts[layer])
	if child == nil { // no child match
		child = &TreeNode{
			pattern: parts[layer],
		}
		node.children = append(node.children, child)
	}
	child.insert(path, parts, layer+1)
}

func (node *TreeNode) matchChild(part string) *TreeNode {
	for _, child := range node.children {
		if child.pattern == part {
			return child
		}
	}
	return nil
}

func (node *TreeNode) matchChildren(part string) []*TreeNode {
	children := make([]*TreeNode, 0, len(node.children))
	for _, child := range node.children {
		if child.pattern == part {
			children = append(children, child)
		}
	}
	return children
}

func (node *TreeNode) search(parts []string, layer int) *TreeNode {
	if len(parts) == layer {
		return node
	}

	part := parts[layer]
	children := node.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, layer+1)
		if result != nil {
			return result
		}
	}

	return nil
}
