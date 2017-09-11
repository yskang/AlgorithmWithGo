package tests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExamplePartitionEqualSubsetSum() {
	fmt.Println(leetcode.CanPartition([]int{1,5,11,5}))
	// output:
	// true
}
