package leetcode

import (
	"AlgorithmWithGo/myLibs"
)

func MostFrequentSubtreeSum(root *myLibs.TreeNode) []int {
	return findFrequentTreeSum(root)
}

func findFrequentTreeSum(root *myLibs.TreeNode) []int {
	sumOfSubtree := make(map[int]int)
	freqMap := make(map[int][]int)

	if root == nil {
		return []int{}
	}

	getSumOfSubtree(root, sumOfSubtree)

	bestFreq := 0
	for sum, freq := range sumOfSubtree {
		if freq > bestFreq {
			bestFreq = freq
		}
		if _, isExist := freqMap[freq]; isExist {
			freqMap[freq] = append(freqMap[freq], sum)
		} else {
			freqMap[freq] = []int{sum}
		}
	}

	return freqMap[bestFreq]
}

func getSumOfSubtree(node *myLibs.TreeNode, sumOfSubtree map[int]int) int {
	if node == nil {
		return 0
	}

	sum := 0

	sum = node.Val + getSumOfSubtree(node.Left, sumOfSubtree) + getSumOfSubtree(node.Right, sumOfSubtree)

	sumOfSubtree[sum] += 1

	return sum
}