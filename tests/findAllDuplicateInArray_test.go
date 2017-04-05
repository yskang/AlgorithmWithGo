package tests

import (
	"testing"
)

func TestList(t *testing.T) {
	myList := make([]int, 0)

	addToList(&myList, 99)
	addToList(&myList, 100)

	if len(myList) == 0 {
		t.Error("list is empty")
	} else {
		t.Log(myList)
	}
}

func addToList(list *[]int, num int) {
	*list = append(*list, num)
}