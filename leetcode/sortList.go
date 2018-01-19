package leetcode

import (
	"github.com/yskang/leetcodeUtil/leetData"
)

func SortList(head *leetData.ListNode) *leetData.ListNode {
	return sortList(head)
}

func sortList(head *leetData.ListNode) *leetData.ListNode {
	r, _ := quickSortList(head)
	return r
}

func quickSortList(list *leetData.ListNode) (*leetData.ListNode, *leetData.ListNode) {
	if list == nil {
		return nil, nil
	} else if list.Next == nil {
		return list, list
	}

	pivot := list
	var current, less, lessTail, more, moreTail, equal, equalTail *leetData.ListNode
	less, more, equal = &leetData.ListNode{}, &leetData.ListNode{}, &leetData.ListNode{}
	current, lessTail, equalTail, moreTail = list, less, equal, more

	for current.Next != nil {
		if current.Val < pivot.Val {
			lessTail.Next = current
			lessTail = current
			current = current.Next
			lessTail.Next = nil
		} else if current.Val > pivot.Val {
			moreTail.Next = current
			moreTail = current
			current = current.Next
			moreTail.Next = nil
		} else {
			equalTail.Next = current
			equalTail = current
			current = current.Next
			equalTail.Next = nil
		}
	}

	if current.Val < pivot.Val {
		lessTail.Next = current
		lessTail = current
	} else if current.Val > pivot.Val {
		moreTail.Next = current
		moreTail = current
	} else {
		equalTail.Next = current
		equalTail = current
	}

	lessH, lessT := quickSortList(less.Next)
	moreH, moreT := quickSortList(more.Next)

	hts := make([]headTail, 0)
	if lessH != nil {
		hts = append(hts, headTail{head: lessH, tail: lessT})
	}
	if equal != nil {
		hts = append(hts, headTail{head: equal.Next, tail: equalTail})
	}
	if moreH != nil {
		hts = append(hts, headTail{head: moreH, tail: moreT})
	}

	retHead := hts[0].head
	retTail := hts[0].tail

	for _, h := range hts[1:] {
		retTail.Next = h.head
		retTail = h.tail
	}

	return retHead, retTail
}

type headTail struct {
	head *leetData.ListNode
	tail *leetData.ListNode
}
