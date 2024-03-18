package red_black_tree

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
	"reflect"
)

func constructRBT(keys []int) {
	// 创建一个新的红黑树
	t := redblacktree.NewWithIntComparator()

	// 向红黑树中插入一些元素
	for _, key := range keys {
		t.Put(key, "")
	}

	// 遍历红黑树
	postOrderTraversal(t.Root)
}

func getNodeColor(node *redblacktree.Node) string {
	value := reflect.ValueOf(node).Elem().FieldByName("color")
	isBlack := value.Bool()
	if isBlack {
		return "Black"
	} else {
		return "Red"
	}
}

func postOrderTraversal(node *redblacktree.Node) {
	if node == nil {
		return
	}
	postOrderTraversal(node.Left)
	postOrderTraversal(node.Right)
	fmt.Printf("Key: %v, Color: %s\n", node.Key, getNodeColor(node))
}
