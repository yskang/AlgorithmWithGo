package main

import (
	"fmt"
	"strconv"
	"strings"
	"sort"
	"math"
	"bufio"
	"os"
)

func main() {
	var numTest int
	fmt.Scanf("%d\n", &numTest)

	scanner := bufio.NewScanner(os.Stdin)

	for t := 0 ; t < numTest ; t++ {
		graph := InitGraph()
		var inLine string
		var N, M int
		//fmt.Scanf("%d %d\n", &N, &M)

		for scanner.Scan() {
			inLine = scanner.Text()
			break
		}
		ins := strings.Split(inLine, " ")
		N, _ = strconv.Atoi(ins[0])
		M, _ = strconv.Atoi(ins[1])

		for m := 0 ; m < M ; m++ {
			var cA, cB int
			//fmt.Scanf("%d %d\n", &cA, &cB)
			for scanner.Scan() {
				inLine = scanner.Text()
				break
			}
			ins := strings.Split(inLine, " ")
			cA, _ = strconv.Atoi(ins[0])
			cB, _ = strconv.Atoi(ins[1])

			graph.AddNode(cA)
			graph.AddNode(cB)
			graph.AddEdge(cA, cB, 10)
		}
		path, _ := Dijsktra(graph, 1)
		//past := make([]int, 0)
		//multiPath := make([][]int, 0)
		//GetMultiPath(path, 1, N, past, &multiPath)
		cities := make([]int, 0)
		DFS(&path, N, &cities)
		fmt.Println(makeString(cities))
	}
}
func makeString(cities []int) string {
	str := make([]string, 0)
	sort.Ints(cities)
	for _, city := range cities {
		str = append(str, strconv.Itoa(city))
	}
	return strings.Join(str, " ")
}
func getCities(paths [][]int) string {
	cities := make([]int, 0)
	outStr := make([]string, 0)
	for _, path := range paths {
		for _, city := range path {
			if !isInList(cities, city) {
				cities = append(cities, city)
			}
		}
	}

	sort.Ints(cities)

	for _, city := range cities {
		outStr = append(outStr, strconv.Itoa(city))
	}

	return strings.Join(outStr, " ")
}

func isInList(list []int, elem int) bool {
	for _, v := range list {
		if v == elem {
			return true
		}
	}
	return false
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

func DFS(path *map[int][]int, v int, discovered *[]int) {
	*discovered = append(*discovered, v)

	for _, node := range (*path)[v] {
		if isInSlice(discovered, node) == false {
			DFS(path, node, discovered)
		}
	}
}

func isInSlice(slice *[]int, elem int) bool {
	for _, v := range *slice {
		if elem == v {
			return true
		}
	}
	return false
}
