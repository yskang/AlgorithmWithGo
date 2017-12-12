package leetcode

import "AlgorithmWithGo/myLibs"

func RobIII(root *myLibs.TreeNode) int {
	return robIII(root)
}

func robIII_(root *myLibs.TreeNode) int {
	return robWithOrWithout(root)
}

func robWithoutRoot(r *myLibs.TreeNode) int {
	if r == nil {
		return 0
	}
	return robWithOrWithout(r.Left) + robWithOrWithout(r.Right)
}

func robWithOrWithout(r *myLibs.TreeNode) int {
	if r == nil {
		return 0
	}
	return getMaxRob(robWithoutRoot(r.Left) + robWithoutRoot(r.Right) + r.Val, robWithoutRoot(r))
}

func robIII(root *myLibs.TreeNode) int {
	robWithoutRoot := func(r *myLibs.TreeNode) int {
		if r == nil {
			return 0
		}
		return robWithOrWithout(r.Left) + robWithOrWithout(r.Right)
	}
	robWithOrWithout := func(r *myLibs.TreeNode) int {
		if r == nil {
			return 0
		}
		return getMaxRob(robWithoutRoot(r.Left) + robWithoutRoot(r.Right) + r.Val, robWithoutRoot(r))
	}

	return robWithOrWithout(root)
}

func getMaxRob(a, b int) int {
	if a < b {
		return b
	}
	return a
}