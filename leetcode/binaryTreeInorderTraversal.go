package leetcode

import (
	"github.com/yskang/leetcodeUtil/leetData"
)

func InorderTraversal(root *leetData.TreeNode) []int {
	return inorderTraversal(root)
}

func inorderTraversal(root *leetData.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(inorder(root.Left), root.Val), inorder(root.Right)...)
}

func inorder(node *leetData.TreeNode) []int {
	if node == nil {
		return []int{}
	}

	return append(append(inorder(node.Left), node.Val), inorder(node.Right)...)
}
