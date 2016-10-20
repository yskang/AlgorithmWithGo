// https://algospot.com/judge/problem/read/WHITECOLLAR

package algospot

import (
	"fmt"
	"strconv"
	"strings"
	"math"
	"bufio"
	"os"
)

func WhiteCollar() {
	var numTest int
	scanner := bufio.NewScanner(os.Stdin)
	var in string
	for scanner.Scan() {
		in = scanner.Text()
		break
	}
	numTest, _ = strconv.Atoi(in)

	for t := 0 ; t < numTest ; t++ {
		var inLine string
		var N, M int

		for scanner.Scan() {
			inLine = scanner.Text()
			break
		}
		ins := strings.Split(inLine, " ")
		N, _ = strconv.Atoi(ins[0])
		M, _ = strconv.Atoi(ins[1])

		graph := InitGraph(N)

		for m := 0 ; m < M ; m++ {
			var cA, cB int
			for scanner.Scan() {
				inLine = scanner.Text()
				break
			}
			ins := strings.Split(inLine, " ")
			cA, _ = strconv.Atoi(ins[0])
			cB, _ = strconv.Atoi(ins[1])

			graph.AddEdge(cA, cB, 1)
		}
		path, _ := dijsktra(graph, 1)
		cities := make([]bool, graph.GetNumOfNodes()+1)
		DFS(&path, N, &cities)
		fmt.Println(makeString(cities))
	}
}
func makeString(cities []bool) string {
	str := make([]string, 0)
	for city, visit := range cities {
		if visit {
			str = append(str, strconv.Itoa(city))
		}
	}
	return strings.Join(str, " ")
}

func dijsktra(graph Graph, start int) (map[int][]int, map[int]int) {
	visited := make([]bool, graph.GetNumOfNodes()+1)
	distance := make(map[int]int)
	path := make(map[int][]int)
	restNode := make([]int, graph.GetNumOfNodes()+1)

	for i, v := range *graph.GetNodes() {
		if v {
			distance[i] = math.MaxInt16
			restNode[i] = 1
			visited[i] = false
			path[i] = []int{}
		}
	}

	distance[start] = 0

	numOfRestNode := graph.GetNumOfNodes()

	for numOfRestNode != 0 {
		minNode := math.MinInt16
		minDistance := math.MaxInt16

		for node := range distance {
			if distance[node] < minDistance && !visited[node] {
				minDistance = distance[node]
				minNode = node
			}
		}

		if minNode < 0 {
			break
		}
		if restNode[minNode] != 0 {
			restNode[minNode] = 0
			numOfRestNode -= 1
		}

		for _, node := range ((*graph.GetEdges())[minNode]) {
			newDistance := distance[minNode] + 1
			if newDistance < distance[node] {
				path[node] = []int{}
				distance[node] = newDistance
				path[node] = append(path[node], minNode)
			} else if newDistance == distance[node] {
				path[node] = append(path[node], minNode)
			}
		}

		visited[minNode] = true
	}

	return path, distance
}

type Edge struct {
	start int
	end int
}

type Graph struct {
	nodes []bool
	edges map[int][]int
}

func InitGraph(n int) Graph {
	graph := Graph{
		make([]bool, n + 1),
		make(map[int][]int),
	}

	return graph
}

func (g *Graph) AddEdge(start int, to int, distance int) {
	g.nodes[start] = true
	g.nodes[to] = true
	g.edges[start] = append(g.edges[start], to)
}

func (g *Graph) GetNumOfNodes() int {
	return len(g.nodes) - 1
}

func (g *Graph) GetNodes() *([]bool) {
	return &g.nodes
}

func (g *Graph) GetEdges() *(map[int][]int) {
	return &g.edges
}

func DFS(path *map[int][]int, v int, discovered *[]bool) {
	(*discovered)[v] = true

	for _, node := range (*path)[v] {
		if (*discovered)[node] == false {
			DFS(path, node, discovered)
		}
	}
}
