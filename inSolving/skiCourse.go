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
		inStrs := strings.Split(inStr, " ")
		N, _ := strconv.Atoi(inStrs[0])
		M, _ := strconv.Atoi(inStrs[1])
		S, _ := strconv.Atoi(inStrs[2])

		graph := InitGraph(N)
		cache := make([][101]int, N + 1)

		for m := 0 ; m < M ; m++ {
			var inStr2 string
			for scanner.Scan() {
				inStr2 = scanner.Text()
				break
			}
			inStrs2 := strings.Split(inStr2, " ")
			A, _ := strconv.Atoi(inStrs2[0])
			B, _ := strconv.Atoi(inStrs2[1])
			C, _ := strconv.Atoi(inStrs2[2])

			if C >= 0 {
				if A < B {
					graph.AddEdge(B, A, C)
				} else {
					graph.AddEdge(A, B, C)
				}
			}
		}

		maxExcitingValue := math.MinInt16
		for node := 1 ; node <= N ; node++ {
			var excitingValue int

			if cache[node][S] != 0 {
				excitingValue = cache[node][S]
			} else {
				excitingValue = getMaxExcitingValueAt(&cache, graph, node, S)
				cache[node][S] = excitingValue
			}


			if maxExcitingValue < excitingValue {
				maxExcitingValue = excitingValue
			}

		}
		fmt.Println("call count", count)
		fmt.Println("cache hit count",cacheCount)
		fmt.Println(maxExcitingValue)
	}

}

func getMaxExcitingValueAt(cache *[][101]int, graph Graph, node int, depth int) int {
	count += 1
	if len((*graph.GetEdges())[node]) == 0 {
		return 0
	}

	if depth == 1 {
		if (*cache)[node][depth] != 0 {
			count -= 1
			return (*cache)[node][depth]
		} else {
			maxExcitingValue := math.MinInt16
			for _, value := range (*graph.GetEdges())[node] {
				if maxExcitingValue < (*graph.GetDistances())[Edge{node, value}] {
					maxExcitingValue = (*graph.GetDistances())[Edge{node, value}]
				}
			}
			(*cache)[node][depth] = maxExcitingValue
			return maxExcitingValue
		}
	}

	maxExcitingValue := math.MinInt16
	maxNode := 0

	for _, prevNode := range (*graph.GetEdges())[node]{
		var excitingValue int
		candidates := make([]int, 0)
		if (*cache)[prevNode][depth-1] != 0 {
			//fmt.Println("hit!")
			cacheCount += 1
			excitingValue = (*cache)[prevNode][depth-1]
		} else {
			excitingValue = getMaxExcitingValueAt(cache, graph, prevNode, depth -1)
		}
		if excitingValue > maxExcitingValue {
			maxExcitingValue = excitingValue
			maxNode = prevNode
			candidates = candidates[:0]
		} else if excitingValue == maxExcitingValue {
			if (*graph.GetDistances())[Edge{node, prevNode}] > (*graph.GetDistances())[Edge{node, maxNode}] {
				maxNode = prevNode
			}
		}
	}

	excitingValue := maxExcitingValue + (*graph.GetDistances())[Edge{node, maxNode}]
	if excitingValue <= 0 {
		return 0
	}
	return excitingValue
}

type Edge struct {
	start int
	end int
}

type Graph struct {
	edges map[int][]int
	distances map[Edge]int
}

func InitGraph(n int) Graph {
	graph := Graph{
		make(map[int][]int),
		make(map[Edge]int),
	}
	for i := 1 ; i <= n ; i++ {
		graph.edges[i] = []int{}
	}
	return graph
}

func (g *Graph) AddEdge(start int, to int, distance int) {
	g.edges[start] = append(g.edges[start], to)
	g.distances[Edge{start, to}] = distance
}

func (g *Graph) GetEdges() *(map[int][]int) {
	return &g.edges
}

func (g *Graph) GetDistances() *(map[Edge]int) {
	return &g.distances
}