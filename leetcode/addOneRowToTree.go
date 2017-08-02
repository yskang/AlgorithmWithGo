package leetcode

import (
	"AlgorithmWithGo/myLibs"
)

func AddOneRowToTree(root *myLibs.TreeNode, v int, d int) *myLibs.TreeNode {
	return addOneRow(root, v, d)
}

func addOneRow(root *myLibs.TreeNode, v int, d int) *myLibs.TreeNode {
	return root
}
