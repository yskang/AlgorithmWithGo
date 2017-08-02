package leetcode

import (
	"AlgorithmWithGo/myLibs"
	"math/rand"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	head *myLibs.ListNode
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func ConstructorListNode(head *myLibs.ListNode) Solution {
	return Solution{head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	res := this.head.Val
	node := this.head.Next
	i := 2

	for node != nil {
		j := rand.Intn(i)
		if j == 0 {
			res = node.Val
		}
		i += 1
		node = node.Next
	}

	return res
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */