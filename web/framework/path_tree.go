package framework

import "strings"

type TreeNode struct {
	path      string
	pattern   string
	children  []*TreeNode
	isDynamic bool
}

// layer means the number of parts which the current insert has already operated at
func (node *TreeNode) insert(path string, parts []string, layer int) {
	// last part in path, no child
	if layer == len(parts) {
		node.path = path
		node.pattern = parts[layer-1]
		return
	}

	// normal
	waitToMatchPart := parts[layer]
	child := node.matchChild(waitToMatchPart)
	if child == nil { // no child match
		child = &TreeNode{
			pattern:   waitToMatchPart,
			isDynamic: waitToMatchPart[0] == ':' || waitToMatchPart[0] == '*',
		}
		node.children = append(node.children, child)
	}
	child.insert(path, parts, layer+1)
}

func (node *TreeNode) matchChild(part string) *TreeNode {
	for _, child := range node.children {
		if child.pattern == part || child.isDynamic == true {
			return child
		}
	}
	return nil
}

func (node *TreeNode) matchChildren(part string) []*TreeNode {
	children := make([]*TreeNode, 0, len(node.children))
	for _, child := range node.children {
		if child.pattern == part || child.isDynamic == true {
			children = append(children, child)
		}
	}
	return children
}

func (node *TreeNode) search(parts []string, layer int) *TreeNode {
	if len(parts) == layer || strings.HasPrefix(node.pattern, "*") {
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
