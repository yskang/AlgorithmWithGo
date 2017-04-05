package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	treeNote := make([]int, 0)
	depth := 0

	treeNote = append(treeNote, root.Val)

	getChildren(root.Left, depth+1, treeNote)
	getChildren(root.Right, depth+1, treeNote)

	fmt.Println(treeNote)

	return treeNote
}

func getChildren(root *TreeNode, depth int, treeNote []int) {
	if len(treeNote) <= depth {
		fmt.Println("append", treeNote, "to", root.Val)
		treeNote = append(treeNote, root.Val)
		fmt.Println(treeNote)
	} else if treeNote[depth] < root.Val {
		fmt.Println("update")
		treeNote[depth] = root.Val
	}

	if root.Left == nil && root.Right == nil {
		return
	}

	if root.Left != nil && root.Right == nil {
		getChildren(root.Left, depth+1, treeNote)
		return
	}

	if root.Left == nil && root.Right != nil {
		getChildren(root.Right, depth+1, treeNote)
		return
	}

	getChildren(root.Left, depth+1, treeNote)
	getChildren(root.Right, depth+1, treeNote)
}
