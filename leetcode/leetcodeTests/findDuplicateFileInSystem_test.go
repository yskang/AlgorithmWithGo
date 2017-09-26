package leetcodeTests

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleFindDuplicateFileInSystem() {
	fmt.Println(leetcode.FindDuplicateFileInSystem([]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}))
	fmt.Println(leetcode.FindDuplicateFileInSystem([]string{"root/a 1.txt(abcd) 2.txt(efsfgh)","root/c 3.txt(abdfcd)","root/c/d 4.txt(efggdfh)"}))
	// Output:
	// [[root/a/2.txt root/c/d/4.txt root/4.txt] [root/a/1.txt root/c/3.txt]]
	// []
}
