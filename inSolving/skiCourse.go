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
		memo := MyCache{ make(map[Ctag]int) }

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


			if A < B {
				graph.AddEdge(B, A, C)
			} else {
				graph.AddEdge(A, B, C)
			}
		}

		maxExcitingValue := math.MinInt16
		for node := 1 ; node <= N ; node++ {
			for i := 1 ; i <= S ; i++ {
				var excitingValue int

				if val, ok := memo.cache[Ctag{node,i}]; ok {
					excitingValue = val
				} else {
					excitingValue = getMaxExcitingValueAt(memo, graph, node, i)
					memo.cache[Ctag{node, i}] = excitingValue
				}


				if maxExcitingValue < excitingValue {
					maxExcitingValue = excitingValue
				}
			}
		}

		fmt.Println(maxExcitingValue)
	}

}

func getMaxExcitingValueAt(memo MyCache, graph Graph, node int, depth int) int {
	if depth == 1 {
		maxExcitingValue := math.MinInt16
		for _, value := range (*graph.GetEdges())[node] {
			if maxExcitingValue < (*graph.GetDistances())[Edge{node, value}] {
				maxExcitingValue = (*graph.GetDistances())[Edge{node, value}]
			}
		}
		return maxExcitingValue
	}

	maxExcitingValue := math.MinInt16
	maxNode := 0

	for _, prevNode := range (*graph.GetEdges())[node]{
		var excitingValue int
		if val, ok := memo.cache[Ctag{prevNode, depth-1}] ; ok {
			excitingValue = val
		} else {
			excitingValue = getMaxExcitingValueAt(memo, graph, prevNode, depth -1)
		}
		if excitingValue > maxExcitingValue {
			maxExcitingValue = excitingValue
			maxNode = prevNode
		}
	}

	excitingValue := maxExcitingValue + (*graph.GetDistances())[Edge{node, maxNode}]
	if excitingValue <= 0 {
		return 0
	}
	return excitingValue
}

type Ctag struct {
	node int
	depth int
}

type MyCache struct {
	cache map[Ctag]int
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