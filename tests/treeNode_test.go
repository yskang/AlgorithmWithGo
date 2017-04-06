package tests

import (
	"testing"
	"../myLibs"
	"fmt"
)

func TestEmptyTreeNode(t *testing.T) {
	root := myLibs.MakeTreeNode("1,2,3,4,5,null, 6, 1,2,3,4,5,6")

	treeMap := make(map[int][]int)
	getNodes(root, treeMap, 0)

	fmt.Println(treeMap)
}

func getNodes(node *myLibs.TreeNode, treeMap map[int][]int, depth int) {
	if node == nil {
		return
	}

	treeMap[depth] = append(treeMap[depth], node.Val)

	if node.Left == nil && node.Right == nil {
		return
	}

	getNodes(node.Left, treeMap, depth+1)
	getNodes(node.Right, treeMap, depth+1)
}

