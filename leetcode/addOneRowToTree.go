package leetcode

import "github.com/yskang/leetcodeUtil/treeNode"

// AddOneRowToTree add One Row to Tree
func AddOneRowToTree(root *treeNode.TreeNode, v int, d int) *treeNode.TreeNode {
	return addOneRow(root, v, d)
}

func addOneRow(root *treeNode.TreeNode, v int, d int) *treeNode.TreeNode {
	row := make([]*treeNode.TreeNode, 0)
	row = append(row, root)

	for i := 1; i < d-1; i++ {
		newRow := make([]*treeNode.TreeNode, 0)
		for _, node := range row {
			if node != nil {
				newRow = append(newRow, []*treeNode.TreeNode{node.Left, node.Right}...)
			}
		}
		row = newRow
	}

	if d == 1 {
		root = &treeNode.TreeNode{Val: v, Left: root, Right: nil}
	} else {
		for _, node := range row {
			if node != nil {
				if node.Left != nil {
					node.Left = &treeNode.TreeNode{Val: v, Left: node.Left, Right: nil}
				} else {
					node.Left = &treeNode.TreeNode{Val: v, Left: nil, Right: nil}
				}

				if node.Right != nil {
					node.Right = &treeNode.TreeNode{Val: v, Left: nil, Right: node.Right}
				} else {
					node.Right = &treeNode.TreeNode{Val: v, Left: nil, Right: nil}
				}
			}
		}
	}

	return root
}
