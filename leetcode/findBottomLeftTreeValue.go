package leetcode

import "sort"

type TreeNode struct {
Val int
Left *TreeNode
Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	treeMark := make(map[string]int)
	markElements(root, treeMark, "1")
	// fmt.Println(treeMark)

	maxMarkLength := 0
	for mark, _ := range treeMark {
		if len(mark) > maxMarkLength {
			maxMarkLength = len(mark)
		}
	}

	maxMarks := make([]string, 0)
	for mark, _ := range treeMark {
		if len(mark) == maxMarkLength {
			maxMarks = append(maxMarks, mark)
		}
	}

	sort.Strings(maxMarks)

	return treeMark[maxMarks[len(maxMarks)-1]]
}

func markElements(root *TreeNode, treeMark map[string]int, currentMark string) {
	treeMark[currentMark] = root.Val

	if root.Left == nil && root.Right == nil {
		return
	}

	if root.Left != nil && root.Right == nil {
		markElements(root.Left, treeMark, currentMark + "1")
		return
	}

	if root.Left == nil && root.Right != nil {
		markElements(root.Right, treeMark, currentMark + "0")
		return
	}

	markElements(root.Left, treeMark, currentMark + "1")
	markElements(root.Right, treeMark, currentMark + "0")
}
