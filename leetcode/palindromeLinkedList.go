package main

import (
	"fmt"
	"AlgorithmWithGo/myLibs"
)

func main() {
	head := myLibs.InitListNode([]int{})
	fmt.Println(isPalindrome(head))
}

func isPalindrome(head *myLibs.ListNode) bool {
	if head == nil {
		return true
	}

	mem := []int{}

	for {
		mem = append(mem, head.Val)
		if head.Next != nil {
			head = head.Next
		} else {
			break
		}
	}

	length := len(mem)

	for i := 0 ; i < length ; i++ {
		if mem[i] != mem[length - i - 1] {
			return false
		}
	}

	return true
}