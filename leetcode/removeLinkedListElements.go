package leetcode

import (
	"../myLibs"
	"fmt"
)

func RemoveLinkedListElements() {
	head := myLibs.InitListNode([]int{1,1})
	head = removeElements(head, 1)

	for {
		if head.Next == nil {
			fmt.Println(head.Val)
			break
		} else {
			fmt.Println(head.Val)
			head = head.Next
		}
	}
}

func removeElements(head *myLibs.ListNode, val int) *myLibs.ListNode {
	if head == nil {
		return nil
	}

	tempPointer := head

	for {
		if head.Val != val {
			tempPointer = head
			break
		} else {
			if head.Next == nil {
				return nil
			}
			head = head.Next
		}
	}

	for {
		if head.Next != nil {
			if head.Next.Val == val {
				head.Next = head.Next.Next
			} else {
				head = head.Next
			}
		} else {
			break
		}
	}

	return tempPointer
}