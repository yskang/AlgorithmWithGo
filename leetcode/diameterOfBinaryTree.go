package leetcode

import (
	"AlgorithmWithGo/myLibs"
)

func DiameterOfBinaryTree(root *myLibs.TreeNode) int {
	return diameterOfBinaryTree(root)
}

func diameterOfBinaryTree(root *myLibs.TreeNode) int {
	if root == nil {
		return 0
	}

	maxDepth := 0
	pTree(root, &maxDepth)

	return maxDepth
}

func pTree(root *myLibs.TreeNode, maxDepth *int) {
	if root == nil {
		return
	}

	d := getDepth(root)

	if d > *maxDepth {
		*maxDepth = d
	}

	pTree(root.Left, maxDepth)
	pTree(root.Right, maxDepth)
}

func getDepth(root *myLibs.TreeNode) int {
	if root == nil {
		return 0
	}

	left := depthTree(root.Left, 0)
	right := depthTree(root.Right, 0)

	return left + right
}

func depthTree(root *myLibs.TreeNode, curr int) int {
	if root == nil {
		return curr
	}

	left := depthTree(root.Left, curr+1)
	right := depthTree(root.Right, curr+1)

	if left > right {
		return left
	}

	return right
}

