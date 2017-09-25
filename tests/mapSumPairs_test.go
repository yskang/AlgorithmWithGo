package tests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleMapSumPairs() {
	mapSum := leetcode.ConstructorMap()
	mapSum.Insert("apple", 3)
	fmt.Println(mapSum.Sum("ap"))
	mapSum.Insert("app", 2)
	fmt.Println(mapSum.Sum("ap"))
	// output:
	// 3
	// 5
}