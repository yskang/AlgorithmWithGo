package leetcode

import (
	"math/rand"
	"time"
)

type ShSolution struct {
	origin []int
	shuffled []int
}


func ConstructorShuffle(nums []int) ShSolution {
	solution := ShSolution{
		nums,
		make([]int, len(nums)),
	}
	return solution
}

/** Resets the array to its original configuration and return it. */
func (this *ShSolution) Reset() []int {
	this.shuffled = append([]int{}, this.origin...)
	return this.origin
}


/** Returns a random shuffling of the array. */
func (this *ShSolution) Shuffle() []int {
	rand.Seed(time.Now().UTC().UnixNano())
	randIndex := rand.Perm(len(this.origin))
	for i, randI := range randIndex {
		this.shuffled[i] = this.origin[randI]
	}
	return this.shuffled
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
