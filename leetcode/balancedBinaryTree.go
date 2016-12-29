package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBalanced(root.Left) && isBalanced(root.Right) && math.Abs(checkDepth(root.Left) - checkDepth(root.Right)) < 2
}

func checkDepth(root *TreeNode) float64 {
	if root == nil {
		return 0
	}

	return 1 + math.Max(checkDepth(root.Left), checkDepth(root.Right))
}