package main

import (
	"fmt"
	"math"
)

func main() {
	stack := Constructor()
	stack.Push(10)
	stack.Push(11)
	stack.Push(12)
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
}

type MinStack struct {
	data []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	data := make([]int, 0)
	return MinStack{data}
}


func (this *MinStack) Push(x int)  {
	this.data = append(this.data, x)
}


func (this *MinStack) Pop()  {
	this.data = this.data[:len(this.data)-1]
}


func (this *MinStack) Top() int {
	return this.data[len(this.data) - 1]
}

func (this *MinStack) GetMin() int {
	min := math.MaxInt64
	for _, data := range(this.data) {
		if data < min {
			min = data
		}
	}
	return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
