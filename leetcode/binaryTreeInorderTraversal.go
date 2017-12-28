package leetcode

import (
	"github.com/yskang/leetcodeUtil/treeNode"
)

func InorderTraversal(root *treeNode.TreeNode) []int {
	return inorderTraversal(root)
}

func inorderTraversal(root *treeNode.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(inorder(root.Left), root.Val), inorder(root.Right)...)
}

func inorder(node *treeNode.TreeNode) []int {
	if node == nil {
		return []int{}
	}

	return append(append(inorder(node.Left), node.Val), inorder(node.Right)...)
}
