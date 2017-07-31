package leetcode

import (
	"sort"
)

func TopKFrequentElements(nums []int, k int) []int {
	return topKFrequent(nums, k)
}

func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num] += 1
	}
	pairs := make([]pair, 0)
	for num, freq := range numMap {
		pairs = append(pairs, pair{num, freq})
	}

	sort.Slice(pairs, func(i int, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	output := make([]int, k)

	for i := 0 ; i < k ; i++ {
		output[i] = pairs[i].num
	}

	return output
}

type pair struct {
	num int
	freq int
}