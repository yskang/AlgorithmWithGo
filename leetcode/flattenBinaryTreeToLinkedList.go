package leetcode

import (
	"github.com/yskang/leetcodeUtil/leetData"
)

// Flatten is a solution of a problem '114. Flatten Binary Tree to Linked List' of leetcode.
func Flatten(root *leetData.TreeNode) {
	flatten(root)
}

func flatten(root *leetData.TreeNode) {
	flattenTree(root)
}

func flattenTree(root *leetData.TreeNode) *leetData.TreeNode {
	if root == nil {
		return nil
	}

	if root.Left != nil && root.Right != nil {
		tmp := root.Right
		root.Right = root.Left
		root.Left = nil
		last := flattenTree(root.Right)
		last.Right = tmp
		return flattenTree(tmp)
	} else if root.Left != nil {
		root.Right = root.Left
		root.Left = nil
		return flattenTree(root.Right)
	} else if root.Right != nil {
		return flattenTree(root.Right)
	}
	return root
}
