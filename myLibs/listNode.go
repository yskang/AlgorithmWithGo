package myLibs

// ListNode library for leetcode solutions

type ListNode struct {
	Val int
	Next *ListNode
}

func InitListNode(nums []int) *ListNode{
	if len(nums) == 0 {
		return nil
	}

	head := ListNode{nums[0], nil}
	tempPointer := &head

	for i := 1 ; i < len(nums) ; i++ {
		tempPointer.Next = &ListNode{nums[i], nil}
		tempPointer = tempPointer.Next
	}

	return &head
}
