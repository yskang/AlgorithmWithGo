package leetcode

import "AlgorithmWithGo/myLibs"

func MergeTrees(t1 *myLibs.TreeNode, t2 *myLibs.TreeNode) *myLibs.TreeNode {
	return mergeTrees(t1, t2)
}

func mergeTrees(t1 *myLibs.TreeNode, t2 *myLibs.TreeNode) *myLibs.TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	} else if t1 == nil || t2 == nil {
		if t1 == nil {
			return t2
		}
		return t1
	}

	t1.Val += t2.Val

	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)

	return t1
}

