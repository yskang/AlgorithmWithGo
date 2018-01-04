package leetcode

import (
	"github.com/yskang/leetcodeUtil/leetData"
)

//BuildTree is a solution of a problem "105. Construct Binary Tree from Preorder and Inorder Traversal" in leetcode
func BuildTree(preorder []int, inorder []int) *leetData.TreeNode {
	return buildTree(preorder, inorder)
}

func buildTree(preorder []int, inorder []int) *leetData.TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	left, right := make([]int, 0), make([]int, 0)
	for i, e := range inorder {
		if e == preorder[0] {
			left = append(left, inorder[:i]...)
			right = append(right, inorder[i+1:]...)
			break
		}
	}

	return &leetData.TreeNode{preorder[0], buildTree(preorder[1:1+len(left)], left), buildTree(preorder[1+len(left):], right)}
}
