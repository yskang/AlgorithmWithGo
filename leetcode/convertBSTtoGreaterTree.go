package leetcode

import (
	"AlgorithmWithGo/myLibs"
)

var sum int

func ConvertBST(root *myLibs.TreeNode) *myLibs.TreeNode {
	return convertBST(root)
}

func convertBST(root *myLibs.TreeNode) *myLibs.TreeNode {
	sum = 0
	convertTree(root)
	return root
}

func convertTree(node *myLibs.TreeNode) {
	if node == nil {
		return
	}
	convertTree(node.Right)
	node.Val += sum
	sum = node.Val
	convertTree(node.Left)
}