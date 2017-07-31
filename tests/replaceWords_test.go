package tests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleReplaceWords() {
	fmt.Println(leetcode.ReplaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))

	// output:
	// the cat was rat by the bat
}
