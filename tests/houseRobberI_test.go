package tests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleHouseRobI() {
	fmt.Println(leetcode.Rob1([]int{114,117,207,117,235,82,90,67,143,146,53,108,200,91,80,223,58,170,110,236,81,90,222,160,165,195,187,199,114,235,197,187,69,129,64,214,228,78,188,67,205,94,205,169,241,202,144,240}))
	// output:
	// 4173
}

func ExampleHouseRobOldSolution() {
	leetcode.HouseRobber()
	// output:
	//0
	//1
	//2
	//4
	//6
	//9
	//49
	//4173
}