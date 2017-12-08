package leetcode

import (
	"AlgorithmWithGo/myLibs"
)

func IsSubtree(s *myLibs.TreeNode, t *myLibs.TreeNode) bool {
	return isSubtree(s, t)
}

func isSubtree(s *myLibs.TreeNode, t *myLibs.TreeNode) bool {
	for c := range getNextNode(s) {
		if compTree(c, t) {
			return true
		}
	}

	return false
}

func getNextNode(s *myLibs.TreeNode) <-chan *myLibs.TreeNode {
	c := make(chan *myLibs.TreeNode)
	go func() {
		var rc func(*myLibs.TreeNode)
		rc = func(n *myLibs.TreeNode) {

			if n == nil {
				return
			}

			c <- n

			rc(n.Left)
			rc(n.Right)
		}
		rc(s)
		close(c)
	}()
	return c
}

func compTree(a *myLibs.TreeNode, b *myLibs.TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil && b != nil {
		return false
	} else if a != nil && b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	if compTree(a.Left, b.Left) == false {
		return false
	}
	if compTree(a.Right, b.Right) == false {
		return false
	}

	return true
}