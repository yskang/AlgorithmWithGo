package tests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleFriendCircles() {
	fmt.Println(leetcode.FriendgCircles([][]int{
		{1,1,0},
		{1,1,0},
		{0,0,1},
	}))

	fmt.Println(leetcode.FriendgCircles([][]int{
		{1,1,0},
		{1,1,1},
		{0,1,1},
	}))


	// output:
	// 2
	// 1
}
