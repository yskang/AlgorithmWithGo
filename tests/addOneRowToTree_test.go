package tests

import (
	"fmt"
	"AlgorithmWithGo/myLibs"
)

func ExampleAddOneRowToTree() {
	fmt.Println(myLibs.PrintTreeNode(myLibs.MakeTreeNode("4,1,1,2,null,null,6,3,1,5")))
	// output:
	// false
}
