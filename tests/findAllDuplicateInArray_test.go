package tests

import (
	"testing"
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func TestList(t *testing.T) {
	myList := make([]int, 0)

	addToList(&myList, 99)
	addToList(&myList, 100)

	if len(myList) == 0 {
		t.Error("The list is not updated!")
	} else {
		fmt.Println(myList)
	}
}

func addToList(list *[]int, num int) {
	*list = append(*list, num)
}

func TestDuplicateInArray(t *testing.T) {
	result := leetcode.FindAllDuplicatInArray([]int{4,3,2,7,8,2,3,1})
	answer := []int{2,3}

	if len(result) != len(answer) {
		t.Error("Wrong! Size is different,", result)
		return
	}

	for i, r := range result {
		if r != answer[i] {
			t.Error("Wrong!", result)
			return
		}
	}
}