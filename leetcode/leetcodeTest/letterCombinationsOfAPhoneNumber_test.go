package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleLetterCombinations() {
	fmt.Println(leetcode.LetterCombinations(""))
	fmt.Println(leetcode.LetterCombinations("23"))
	fmt.Println(leetcode.LetterCombinations("234"))
	// output:
	// [ad ae af bd be bf cd ce cf]
}
