package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	nodes := make([]int, 0)

	addNode(root, 0, &nodes)

	return nodes
}

func addNode(root *TreeNode, depth int, nodes *[]int) {
	if root == nil {
		return
	}

	if len(*nodes) == 0 || len(*nodes) <= depth {
		*nodes = append(*nodes, root.Val)
	} else if (*nodes)[depth] < root.Val {
		(*nodes)[depth] = root.Val
	}

	if root.Left == nil && root.Right == nil {
		return
	}

	addNode(root.Left, depth+1, nodes)
	addNode(root.Right, depth+1, nodes)
}