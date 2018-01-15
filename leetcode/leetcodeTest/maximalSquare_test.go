package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleMaximalSquare() {
	fmt.Println(leetcode.MaximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))

	// output:
	// 4
}

func ExampleMaximalSquare2() {
	fmt.Println(leetcode.MaximalSquare([][]byte{
		{'0', '0', '0'},
		{'0', '0', '0'},
		{'1', '1', '1'},
	}))
	// output:
	// 1
}

func ExampleMaximalSquare3() {
	fmt.Println(leetcode.MaximalSquare([][]byte{
		{'1', '1'},
		{'1', '1'},
	}))
	// output:
	// 4
}

func ExampleMaximalSquare4() {
	fmt.Println(leetcode.MaximalSquare([][]byte{
		{'0', '0', '0', '1'},
		{'1', '1', '0', '1'},
		{'1', '1', '1', '1'},
		{'0', '1', '1', '1'},
		{'0', '1', '1', '1'},
	}))
	// output:
	// 9
}
