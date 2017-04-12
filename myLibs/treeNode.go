package myLibs

import (
	"strings"
	"strconv"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func MakeTreeNode(inputString string) *TreeNode {
	if inputString == "" {
		return nil
	}

	root := new(TreeNode)
	rows := make([]*TreeNode, 0)

	inputStringsSplitedByComma := strings.Split(inputString, ",")

	inputStringsTrimmed := make([]string, 0)
	for _, s := range inputStringsSplitedByComma {
		inputStringsTrimmed = append(inputStringsTrimmed, strings.TrimSpace(s))
	}

	root.Val, _ = strconv.Atoi(inputStringsTrimmed[0])
	rows = append(rows, root)

	offset := 0
	for {
		offset += len(rows)
		if offset == len(inputStringsTrimmed) {
			break
		}
		temp := make([]*TreeNode, 0)
		for j:=0 ; j < len(rows) ; j++ {
			for k:= 0 ; k < 2 ; k++ {
				if offset + k + j*2 > len(inputStringsTrimmed) - 1 {
					return root
				}
				input := inputStringsTrimmed[offset + k + j*2]
				val, error := strconv.Atoi(input)
				node := new(TreeNode)

				if error == nil {
					temp = append(temp, node)
					node.Val = val
					if k % 2 == 0 {
						rows[j].Left = node
					} else {
						rows[j].Right = node
					}
				}
			}
		}
		rows = temp
	}

	return root
}

func CompareTreeNode(nodeA *TreeNode, nodeB *TreeNode) bool {
	if nodeA == nil && nodeB == nil {
		return true
	} else if nodeA == nil || nodeB == nil {
		return false
	}

	if nodeA.Val == nodeB.Val {
		return CompareTreeNode(nodeA.Left, nodeB.Left) && CompareTreeNode(nodeA.Right, nodeB.Right)
	}

	return false
}