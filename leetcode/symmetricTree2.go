package leetcode


func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return checkSymmetric(left.Left, right.Right) && checkSymmetric(left.Right, right.Left)
}