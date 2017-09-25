package leetcode

import (
	"AlgorithmWithGo/myLibs"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func ConstructMaximumBinaryTree(nums []int) *myLibs.TreeNode {
	return constructMaximumBinaryTree(nums)
}

func constructMaximumBinaryTree(nums []int) *myLibs.TreeNode {
	return makeMaxBinaryTree(nums)
}

func makeMaxBinaryTree(nums []int) *myLibs.TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &myLibs.TreeNode{nums[0], nil, nil}
	}

	maximumIndex, maximumNum := 0, math.MinInt64
	for i, n := range nums {
		if maximumNum < n {
			maximumNum = n
			maximumIndex = i
		}
	}
	leftTree, rightTree := constructMaximumBinaryTree(nums[:maximumIndex]), constructMaximumBinaryTree(nums[maximumIndex+1:])

	return &myLibs.TreeNode{maximumNum, leftTree, rightTree}
}
