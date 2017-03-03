package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	midIndex := int(len(nums)/2)
	return &TreeNode{nums[midIndex], sortedArrayToBST(nums[:midIndex]), sortedArrayToBST(nums[midIndex+1:])}
}
