package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/myLibs"
)

func ExamplePrintTree()  {
	fmt.Println(myLibs.MakeTreeNode("1,2,3,null,4,null,5,6,null,7,8"))
	// output:
	// 1,2,3,null,4,null,5,6,null,7,8
}
