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

	inputStringsSplitByComma := strings.Split(inputString, ",")

	inputStringsTrimmed := make([]string, 0)
	for _, s := range inputStringsSplitByComma {
		inputStringsTrimmed = append(inputStringsTrimmed, strings.TrimSpace(s))
	}

	root.Val, _ = strconv.Atoi(inputStringsTrimmed[0])
	rows = append(rows, root)

	offset := 1

	for {
		count := 0
		if offset == len(inputStringsTrimmed) {
			break
		}
		temp := make([]*TreeNode, 0)
		for j := 0 ; j < len(rows) ; j++ {
			for k := 0 ; k < 2 ; k++ {
				if offset + k + j * 2 > len(inputStringsTrimmed) - 1 {
					return root
				}
				input := inputStringsTrimmed[offset + k + j * 2]
				count += 1
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
		offset += count
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

func PrintTreeNode(node *TreeNode) string {
	nodeString := make([]string, 0)
	queue := make([]*TreeNode, 0)

	currNode := node

	for {
		if currNode != nil {
			nodeString = append(nodeString, strconv.Itoa(currNode.Val))
			queue = append(queue, currNode.Left)
			queue = append(queue, currNode.Right)
		} else {
			nodeString = append(nodeString, "null")
		}

		if len(queue) > 0 {
			currNode = queue[0]
			queue = queue[1:]
		} else {
			break
		}
	}

	index := len(nodeString) - 1
	for i := len(nodeString) - 1 ; i >= 0 ; i-- {
		if nodeString[i] != "null" {
			index = i + 1
			break
		}
	}

	return strings.Join(nodeString[:index], ",")
}