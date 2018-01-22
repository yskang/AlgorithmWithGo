package leetcode

import (
	"github.com/yskang/leetcodeUtil/leetData"
)

func AddTwoNumbers(l1 *leetData.ListNode, l2 *leetData.ListNode) *leetData.ListNode {
	return addTwoNumbers(l1, l2)
}

func addTwoNumbers(l1 *leetData.ListNode, l2 *leetData.ListNode) *leetData.ListNode {
	res, a, b := &leetData.ListNode{0, nil}, &leetData.ListNode{0, l1}, &leetData.ListNode{0, l2}
	sum, carry := 0, 0
	resCur := res
	for a.Next != nil || b.Next != nil {
		aVal, bVal := 0, 0
		if a.Next != nil {
			aVal = a.Next.Val
			a.Next = a.Next.Next
		}
		if b.Next != nil {
			bVal = b.Next.Val
			b.Next = b.Next.Next
		}
		sum = aVal + bVal + carry
		if sum >= 10 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}
		resCur.Next = &leetData.ListNode{Val: sum, Next: nil}
		resCur = resCur.Next
	}

	if carry != 0 {
		resCur.Next = &leetData.ListNode{Val: carry, Next: nil}
	}

	return res.Next
}
