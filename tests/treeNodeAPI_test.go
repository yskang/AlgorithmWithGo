package tests

import (
	"fmt"
	"AlgorithmWithGo/myLibs"
)

func ExamplePrintTree()  {
	fmt.Println(myLibs.PrintTreeNode(myLibs.MakeTreeNode("1,2,3,4,null,null,null,5,5")))
	// output:
	// 1,2,3,4,null,null,null,5,5
}
