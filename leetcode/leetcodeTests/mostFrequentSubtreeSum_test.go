package leetcodeTests

import (
	"testing"
	"AlgorithmWithGo/leetcode"
	"AlgorithmWithGo/myLibs"
	"fmt"
	"sort"
)

func TestNullInput(t *testing.T) {
	root := myLibs.MakeTreeNode("")

	if len(leetcode.MostFrequentSubtreeSum(root)) != 0 {
		fmt.Println(leetcode.MostFrequentSubtreeSum(root))
		t.Error("null error")
	}
}

func TestSimpleInput(t *testing.T) {
	root := myLibs.MakeTreeNode("5,2,-3")
	ans := []int{-3,2,4}
	result := leetcode.MostFrequentSubtreeSum(root)
	if len(ans) != len(result) {
		t.Error("Length error!")
	}

	sort.Ints(result)
	for i := 0 ; i < len(result) ; i++ {
		if ans[i] != result[i] {
			t.Error("Wrong answer")
			fmt.Println("Expected:", ans)
			fmt.Println("result:  ", result)
		}
	}
}

func Test2ndInput(t *testing.T) {
	root := myLibs.MakeTreeNode("5,2,-5")
	ans := []int{2}
	result := leetcode.MostFrequentSubtreeSum(root)
	if len(ans) != len(result) {
		t.Error("Length error!")
	}

	sort.Ints(result)
	for i := 0 ; i < len(result) ; i++ {
		if ans[i] != result[i] {
			t.Error("Wrong answer")
			fmt.Println("Expected:", ans)
			fmt.Println("result:  ", result)
		}
	}
}