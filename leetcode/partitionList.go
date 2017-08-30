package leetcode

import (
	"AlgorithmWithGo/myLibs"
)

func Partition(head *myLibs.ListNode, x int) *myLibs.ListNode {
	return partition(head, x)
}

func partition(head *myLibs.ListNode, x int) *myLibs.ListNode {
	currentOrigin := head
	headOfSmall := myLibs.ListNode{0, nil}
	headOfLarge := myLibs.ListNode{1, nil}
	currentSmall := &headOfSmall
	currentLarge := &headOfLarge

	for currentOrigin != nil {
		if currentOrigin.Val < x {
			currentSmall.Next = &myLibs.ListNode{currentOrigin.Val, nil}
			currentSmall = currentSmall.Next
		} else {
			currentLarge.Next = &myLibs.ListNode{currentOrigin.Val, nil}
			currentLarge = currentLarge.Next
		}
		currentOrigin = currentOrigin.Next
	}

	currentSmall.Next = headOfLarge.Next

	return headOfSmall.Next
}