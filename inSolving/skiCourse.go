// https://algospot.com/judge/problem/read/SKICOURSE

package main

import (
	"fmt"
	"math"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var count, cacheCount int

func main() {
	var T int
	fmt.Scanf("%d\n", &T)

	scanner := bufio.NewScanner(os.Stdin)

	for t := 0 ; t < T ; t++ {
		var inStr string
		for scanner.Scan() {
			inStr = scanner.Text()
			break
		}
		inLine := strings.Split(inStr, " ")
		N, _ := strconv.Atoi(inLine[0])
		M, _ := strconv.Atoi(inLine[1])
		S, _ := strconv.Atoi(inLine[2])

		graph := make([][]int, 0)
		for i := 0 ; i <= N ; i++ {
			graph = append(graph, make([]int, N+1))
		}

		cache := make([][]int, 0)
		for i := 0 ; i <= N ; i++ {
			cache = append(cache, make([]int, S+1))
		}

		for m := 0 ; m < M ; m++ {
			var inStr2 string
			for scanner.Scan() {
				inStr2 = scanner.Text()
				break
			}
			inLine2 := strings.Split(inStr2, " ")
			A, _ := strconv.Atoi(inLine2[0])
			B, _ := strconv.Atoi(inLine2[1])
			C, _ := strconv.Atoi(inLine2[2])

			graph[A-1][B-1] = C
		}

		maxExcitingValue := math.MinInt16
		for node := 1 ; node <= N ; node++ {
			var excitingValue int

			excitingValue = getMaxExcitingValueAt(&cache, &graph, node, S)
			cache[node][S] = excitingValue


			if maxExcitingValue < excitingValue {
				maxExcitingValue = excitingValue
			}

		}
		fmt.Println("call count", count)

		fmt.Println("cache hit count",cacheCount)
		fmt.Println(cache)
		count = 0
		cacheCount = 0
		fmt.Println(maxExcitingValue)
	}

}

func getMaxExcitingValueAt(cache *[][]int, graph *[][]int, node int, depth int) int {
	if (*cache)[node][depth] != 0 {
		cacheCount += 1
		return (*cache)[node][depth]
	}

	count += 1

	if depth == 1 {
		maxExcitingValue := math.MinInt16

		for prevNode := 0 ; prevNode < len(*graph) ; prevNode++ {
			if maxExcitingValue < (*graph)[prevNode][node] {
				maxExcitingValue = (*graph)[prevNode][node]
			}
		}

		(*cache)[node][depth] = maxExcitingValue
		return maxExcitingValue
	}

	maxExcitingValue := math.MinInt16
	maxNode := 0

	for prevNode := 0 ; prevNode < len(*graph) ; prevNode++ {
		if (*graph)[prevNode][node] != 0 {
			var excitingValue int
			candidates := make([]int, 0)
			excitingValue = getMaxExcitingValueAt(cache, graph, prevNode, depth -1)
			if excitingValue > maxExcitingValue {
				maxExcitingValue = excitingValue
				maxNode = prevNode
				candidates = candidates[:0]
			} else if excitingValue == maxExcitingValue {
				if (*graph)[prevNode][node] > (*graph)[maxNode][node] {
					maxNode = prevNode
				}
			}
		}
	}

	if maxExcitingValue == math.MinInt16 && maxNode == 0 {
		return 0
	}

	excitingValue := maxExcitingValue + (*graph)[maxNode][node]
	if excitingValue <= 0 {
		return 0
	}
	return excitingValue
}