// https://algospot.com/judge/problem/read/SKICOURSE

package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	fmt.Scanf("%d\n", &T)

	for t := 0 ; t < T ; t++ {
		var N, M, S int
		fmt.Scanf("%d %d %d\n", &N, &M, &S)
		graph := InitGraph()

		for m := 0 ; m < M ; m++ {
			var A, B, C int
			fmt.Scanf("%d %d %d\n", &A, &B, &C)


			graph.AddNode(A)
			graph.AddNode(B)
			if A > B {
				graph.AddEdge(B, A, -1 * C)
			} else {
				graph.AddEdge(A, B, -1 * C)
			}

		}
		path, excite := Dijsktra(graph, 1)

		fmt.Println(path)
		fmt.Println(excite)
	}

}

type Edge struct {
	start int
	end int
}

type Graph struct {
	nodes []int
	edges map[int][]int
	distances map[Edge]int
}

func InitGraph() Graph {
	graph := Graph{
		make([]int, 0),
		make(map[int][]int),
		make(map[Edge]int),
	}
	return graph
}

func (g *Graph) AddNode(node int) {
	for _, n := range g.nodes {
		if n == node {
			return
		}
	}
	g.nodes = append(g.nodes, node)
	g.edges[node] = []int{}
}

func (g *Graph) AddEdge(start int, to int, distance int) {
	g.edges[start] = append(g.edges[start], to)
	g.distances[Edge{start, to}] = distance
}

func (g *Graph) GetNodes() *([]int) {
	return &g.nodes
}

func (g *Graph) GetEdges() *(map[int][]int) {
	return &g.edges
}

func (g *Graph) GetDistances() *(map[Edge]int) {
	return &g.distances
}


func Dijsktra(graph Graph, start int) (map[int][]int, map[int]int) {
	visited := make([]int, 0)
	distance := make(map[int]int)
	restNode := make([]int, 0)
	path := make(map[int][]int)

	for _, node := range *(graph.GetNodes()) {
		distance[node] = math.MaxInt16
		restNode = append(restNode, node)
		path[node] = []int{}
	}

	distance[start] = 0

	for len(restNode) != 0 {
		minNode := -9999
		minDistance := math.MaxInt16

		for node := range distance {
			if distance[node] < minDistance && !isContain(node, &visited) {
				minDistance = distance[node]
				minNode = node
			}
		}

		if minNode < 0 {
			break
		}

		restNode = *Remove(minNode, &restNode)

		for _, node := range ((*graph.GetEdges())[minNode]) {
			newDistance := distance[minNode] + (*graph.GetDistances())[Edge{minNode, node}]
			if newDistance < distance[node] {
				path[node] = []int{}
				distance[node] = newDistance
				path[node] = append(path[node], minNode)
			} else if newDistance == distance[node] {
				path[node] = append(path[node], minNode)
			}
		}

		visited = append(visited, minNode)
	}

	return path, distance
}
func Remove(element int, slice *[]int) *[]int{
	temp := make([]int, len(*slice))
	copy(temp, *slice)
	for i, e := range *slice {
		if e == element {
			result := append(temp[:i], temp[i+1:]...)
			return &result
		}
	}
	return slice
}

func isContain(element int, slice *[]int) bool {
	for _, e := range *slice {
		if e == element {
			return true
		}
	}
	return false
}

func GetMultiPath(path map[int][]int, start int, end int, past []int, multiPath *[][]int) {
	if start != end {
		past = append(past, end)
	} else {
		past = append(past, start)
	}

	for _, node := range path[end] {
		GetMultiPath(path, start, node, past, multiPath)
	}

	if past[len(past) - 1] == start {
		*multiPath = append(*multiPath, past)
	}
	past = past[:len(past)-1]
}