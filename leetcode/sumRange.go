package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	arr := Constructor(nums)
	fmt.Println(arr.SumRange(0,2));
	fmt.Println(arr.SumRange(2,5));
	fmt.Println(arr.SumRange(0,5));
}

type NumArray struct {
	intArray []int
}


func Constructor(nums []int) NumArray {
	return NumArray{nums}
}


func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for _, elem := range(this.intArray[i:j+1]) {
		sum += elem
	}
	return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
