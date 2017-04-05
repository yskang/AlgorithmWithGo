package leetcode

import (
	"strconv"
)

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lefts := printTreeLeft(root.Left)
	rights := printTreeRight(root.Right)

	for i := 0 ; i < len(lefts) ; i++ {
		if lefts[i] != rights[i] {
			return false
		}
	}

	return true
}

func printTreeLeft(root *TreeNode) []string {
	if root == nil {
		return []string{"null"}
	}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	returnList := []string{strconv.Itoa(root.Val)}

	for _, elem := range(printTreeLeft(root.Left)) {
		returnList = append(returnList, elem)
	}

	for _, elem := range(printTreeLeft(root.Right)) {
		returnList = append(returnList, elem)
	}

	return returnList
}

func printTreeRight(root *TreeNode) []string {
	if root == nil {
		return []string{"null"}
	}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	returnList := []string{strconv.Itoa(root.Val)}

	for _, elem := range(printTreeRight(root.Right)) {
		returnList = append(returnList, elem)
	}

	for _, elem := range(printTreeRight(root.Left)) {
		returnList = append(returnList, elem)
	}

	return returnList
}
