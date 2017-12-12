package leetcode

import "AlgorithmWithGo/myLibs"

func RobIII(root *myLibs.TreeNode) int {
	return robIII(root)
}

func robIII(root *myLibs.TreeNode) int {
	withRoot, withoutRoot := depthFirst(root)
	return getMaxRob(withRoot, withoutRoot)
}

func depthFirst(r *myLibs.TreeNode) (int, int) {
	if r == nil {
		return 0, 0
	}
	leftWithRoot, leftWithoutRoot := depthFirst(r.Left)
	rightWithRoot, rightWithoutRoot := depthFirst(r.Right)
	withRoot := leftWithoutRoot + rightWithoutRoot + r.Val
	withoutRoot := getMaxRob(leftWithoutRoot, leftWithRoot) + getMaxRob(rightWithRoot, rightWithoutRoot)
	return withRoot, withoutRoot
}


func robIII__(root *myLibs.TreeNode) int {
	return getMaxRob(robWithoutRoot(root), robWithRoot(root))
}

func robWithoutRoot(r *myLibs.TreeNode) int {
	if r == nil {
		return 0
	}
	return robIII(r.Left) + robIII(r.Right)
}

func robWithRoot(r *myLibs.TreeNode) int {
	if r == nil {
		return 0
	}
	return r.Val + robWithoutRoot(r.Left) + robWithoutRoot(r.Right)
}

func robWithOrWithout(r *myLibs.TreeNode) int {
	if r == nil {
		return 0
	}
	return getMaxRob(robWithoutRoot(r.Left) + robWithoutRoot(r.Right) + r.Val, robWithoutRoot(r))
}

func robIII_(root *myLibs.TreeNode) int {
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