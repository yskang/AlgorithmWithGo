package leetcode

import "AlgorithmWithGo/myLibs"

func InorderTraversal(root *myLibs.TreeNode) []int {
	return inorderTraversal(root)
}

func inorderTraversal(root *myLibs.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(inorder(root.Left), root.Val), inorder(root.Right)...)
}

func inorder(node *myLibs.TreeNode) []int {
	if node == nil {
		return []int{}
	}

	return append(append(inorder(node.Left), node.Val), inorder(node.Right)...)
}

