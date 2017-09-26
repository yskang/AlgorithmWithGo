package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleMagicDictionary() {
	magicDic := leetcode.ConstructorMagicDic()
	magicDic.BuildDict([]string{"hello", "hallo", "leetcode"})
	fmt.Println(magicDic.Search("hello"))
	fmt.Println(magicDic.Search("hhllo"))
	fmt.Println(magicDic.Search("hell"))
	fmt.Println(magicDic.Search("leetcoded"))
	// output:
	// true
	// true
	// false
	// false
}