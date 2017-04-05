package leetcode

import (
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	paths := []string{}

	for _, path := range (append(binaryTreePaths(root.Left), binaryTreePaths(root.Right)...)) {
		paths = append(paths, strconv.Itoa(root.Val) + "->" + path)
	}

	return paths
}