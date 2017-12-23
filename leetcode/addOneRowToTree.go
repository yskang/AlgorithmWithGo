package leetcode

import (
	"AlgorithmWithGo/myLibs"
)

// AddOneRowToTree add One Row to Tree
func AddOneRowToTree(root *myLibs.TreeNode, v int, d int) *myLibs.TreeNode {
	return addOneRow(root, v, d)
}

func addOneRow(root *myLibs.TreeNode, v int, d int) *myLibs.TreeNode {
	row := make([]*myLibs.TreeNode, 0)
	row = append(row, root)

	for i := 1; i < d-1; i++ {
		newRow := make([]*myLibs.TreeNode, 0)
		for _, node := range row {
			if node != nil {
				newRow = append(newRow, []*myLibs.TreeNode{node.Left, node.Right}...)
			}
		}
		row = newRow
	}

	if d == 1 {
		root = &myLibs.TreeNode{Val: v, Left: root, Right: nil}
	} else {
		for _, node := range row {
			if node != nil {
				if node.Left != nil {
					node.Left = &myLibs.TreeNode{Val: v, Left: node.Left, Right: nil}
				} else {
					node.Left = &myLibs.TreeNode{Val: v, Left: nil, Right: nil}
				}

				if node.Right != nil {
					node.Right = &myLibs.TreeNode{Val: v, Left: nil, Right: node.Right}
				} else {
					node.Right = &myLibs.TreeNode{Val: v, Left: nil, Right: nil}
				}
			}
		}
	}

	return root
}
