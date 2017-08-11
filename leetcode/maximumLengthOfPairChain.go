package leetcode

import (
	"sort"
)

func FindLongestChain(pairs [][]int) int {
	return findLongestChain(pairs)
}

func findLongestChain(pairs [][]int) int {
	memo := make(map[tuple]int)
	max := 0

	sort.Slice(pairs, func(i int, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	for index, pair := range pairs {
		currentMax := getMax(pair, pairs[index+1:], memo)
		memo[tuple{pair[0], pair[1]}] = currentMax

		if max < currentMax {
			max = currentMax
		}
	}

	return max
}

type tuple struct {
	start int
	end int
}

func getMax(start []int, pairs [][]int, memo map[tuple]int) int {
	max := 0

	for index, pair := range pairs {
		if start[1] < pair[0] {
			currentMax := 0
			if maxValue, isExist := memo[tuple{pair[0], pair[1]}]; isExist {
				currentMax = maxValue
			} else {
				currentMax = getMax(pair, pairs[index+1:], memo)
				memo[tuple{pair[0], pair[1]}] = currentMax
			}

			if max < currentMax {
				max = currentMax
			}
		}
	}

	return max + 1
}
