package utils

import (
	"fmt"
	"sort"
)

func ExampleGCD() {
	fmt.Println(gcd(20, 15))
	fmt.Println(gcd(20, 20))
	fmt.Println(gcd(20, 40))
	fmt.Println(gcd(100, 70))
	//output:
	//5
	//20
	//20
	//10
}

func ExampleLCM() {
	fmt.Println(lcm(10, 20))
	//output:
	//20
}

func ExampleCombination() {
	comb(5, 3, func(c []int) {
		fmt.Println(c)
	})
	//output:
	//[0 1 2]
	//[0 1 3]
	//[0 1 4]
	//[0 2 3]
	//[0 2 4]
	//[0 3 4]
	//[1 2 3]
	//[1 2 4]
	//[1 3 4]
	//[2 3 4]
}

func ExamplePermutation() {
	permute([]int{1,2,3,4}, func(per []int) {
		fmt.Println(per)
	})
	//output:
	//[1 2 3 4]
	//[1 2 4 3]
	//[1 4 2 3]
	//[4 1 2 3]
	//[1 3 2 4]
	//[1 3 4 2]
	//[1 4 3 2]
	//[4 1 3 2]
	//[3 1 2 4]
	//[3 1 4 2]
	//[3 4 1 2]
	//[4 3 1 2]
	//[2 1 3 4]
	//[2 1 4 3]
	//[2 4 1 3]
	//[4 2 1 3]
	//[2 3 1 4]
	//[2 3 4 1]
	//[2 4 3 1]
	//[4 2 3 1]
	//[3 2 1 4]
	//[3 2 4 1]
	//[3 4 2 1]
	//[4 3 2 1]
}

func ExamplePermutationFirst() {
	d := []int{1,2,4,3}
	PermutationFirst(sort.IntSlice(d))
	fmt.Println(d)
	//output:
	//[1 2 3 4]
}

func ExamplePermutationNext() {
	d := []int{1,2,4,3}
	fmt.Println(PermutationNext(sort.IntSlice(d)))
	fmt.Println(d)
	fmt.Println(PermutationNext(sort.IntSlice(d)))
	fmt.Println(d)
	d = []int{4,3,2,1}
	fmt.Println(PermutationNext(sort.IntSlice(d)))
	fmt.Println(d)

	//output:
	//true
	//[1 3 2 4]
	//true
	//[1 3 4 2]
	//false
	//[4 3 2 1]
}