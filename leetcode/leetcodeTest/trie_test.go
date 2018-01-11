package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleTrie() {
	trie := leetcode.ConstructTire()
	trie.Insert("hello")
	trie.Insert("world")
	trie.Insert("help")
	trie.Insert("wow")

	fmt.Println(trie.Search("hello"))
	fmt.Println(trie.Search("wow"))
	fmt.Println(trie.Search("wo"))
	fmt.Println(trie.Search("worldcup"))
	fmt.Println(trie.Search("hell"))
	fmt.Println(trie.StartsWith("hel"))
	fmt.Println(trie.StartsWith("he"))
	fmt.Println(trie.StartsWith("h"))
	fmt.Println(trie.StartsWith("wa"))
	// output:
	// true
	// true
	// false
	// false
	// false
	// true
	// true
	// true
	// false
}
