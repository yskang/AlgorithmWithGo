package leetcode

import (
	"AlgorithmWithGo/myLibs"
)

func AddOneRowToTree(root *myLibs.TreeNode, v int, d int) *myLibs.TreeNode {
	return addOneRow(root, v, d)
}

func addOneRow(root *myLibs.TreeNode, v int, d int) *myLibs.TreeNode {
	row := make([]*myLibs.TreeNode, 0)
	row = append(row, root)

	for i := 1 ; i < d - 1 ; i++ {
		newRow := make([]*myLibs.TreeNode, 0)
		for _, node := range row {
			if node != nil {
				newRow = append(newRow, []*myLibs.TreeNode{node.Left, node.Right}...)
			}
		}
		row = newRow
	}

	if d == 1 {
		root = &myLibs.TreeNode{v, root, nil}
	} else {
		for _, node := range row {
			if node != nil {
				if node.Left != nil {
					node.Left = &myLibs.TreeNode{v, node.Left, nil}
				} else {
					node.Left = &myLibs.TreeNode{v, nil, nil}
				}

				if node.Right != nil {
					node.Right = &myLibs.TreeNode{v, nil, node.Right}
				} else {
					node.Right = &myLibs.TreeNode{v, nil, nil}
				}
			}
		}
	}

	return root
}
