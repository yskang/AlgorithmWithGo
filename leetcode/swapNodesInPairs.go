package leetcode

import (
	"../myLibs"
)

func swapPairs(head *myLibs.ListNode) *myLibs.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := new(myLibs.ListNode)
	first := new(myLibs.ListNode)
	second := new(myLibs.ListNode)
	temp := new(myLibs.ListNode)

	dummyHead.Next = head

	first.Next = head
	second.Next = head.Next
	temp.Next = dummyHead

	for {
		first.Next.Next = second.Next.Next
		second.Next.Next = first.Next
		temp.Next.Next = second.Next
		temp.Next = first.Next

		if first.Next.Next == nil || first.Next.Next.Next == nil {
			break
		}

		first.Next = first.Next.Next
		second.Next = first.Next.Next
	}

	return dummyHead.Next
}