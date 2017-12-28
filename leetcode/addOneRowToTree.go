package leetcode

import (
	"github.com/yskang/leetcodeUtil/leetData"
)

// AddOneRowToTree add One Row to Tree
func AddOneRowToTree(root *leetData.TreeNode, v int, d int) *leetData.TreeNode {
	return addOneRow(root, v, d)
}

func addOneRow(root *leetData.TreeNode, v int, d int) *leetData.TreeNode {
	row := make([]*leetData.TreeNode, 0)
	row = append(row, root)

	for i := 1; i < d-1; i++ {
		newRow := make([]*leetData.TreeNode, 0)
		for _, node := range row {
			if node != nil {
				newRow = append(newRow, []*leetData.TreeNode{node.Left, node.Right}...)
			}
		}
		row = newRow
	}

	if d == 1 {
		root = &leetData.TreeNode{Val: v, Left: root, Right: nil}
	} else {
		for _, node := range row {
			if node != nil {
				if node.Left != nil {
					node.Left = &leetData.TreeNode{Val: v, Left: node.Left, Right: nil}
				} else {
					node.Left = &leetData.TreeNode{Val: v, Left: nil, Right: nil}
				}

				if node.Right != nil {
					node.Right = &leetData.TreeNode{Val: v, Left: nil, Right: node.Right}
				} else {
					node.Right = &leetData.TreeNode{Val: v, Left: nil, Right: nil}
				}
			}
		}
	}

	return root
}
