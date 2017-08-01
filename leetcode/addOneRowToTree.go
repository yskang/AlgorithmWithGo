package leetcode

import (
	"AlgorithmWithGo/myLibs"
	"fmt"
)

func AddOneRowToTree(root *myLibs.TreeNode, v int, d int) *myLibs.TreeNode {
	return addOneRow(root, v, d)
}

func addOneRow(root *myLibs.TreeNode, v int, d int) *myLibs.TreeNode {
	fmt.Println(myLibs.PrintTreeNode(root))
	return root
}
