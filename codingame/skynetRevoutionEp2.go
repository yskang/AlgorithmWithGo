package main

import (
	"fmt"
	"math"
	"os"
)

//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type nodeNumber int

func main() {
	// N: the total number of nodes in the level, including the gateways
	// L: the number of links
	// E: the number of exit gateways
	var N, L, E int

	fmt.Scan(&N, &L, &E)

	fmt.Fprintln(os.Stderr, "N, L, E :", N, L, E)

	graph := InitGraph(N)
	exits := make([]int, E)

	for i := 0; i < L; i++ {
		// N1: N1 and N2 defines a link between these nodes
		var N1, N2 int
		fmt.Scan(&N1, &N2)
		graph.AddEdge(N1, N2, 1)
		graph.AddEdge(N2, N1, 1)
	}
	for i := 0; i < E; i++ {
		// EI: the index of a gateway node
		var EI int
		fmt.Scan(&EI)
		exits[i] = EI
	}
	for {
		// SI: The index of the node on which the Skynet agent is positioned this turn
		var SI int
		fmt.Scan(&SI)

		nodeA, nodeB := 0, 0

		mostDangerousNodes := make([]int, 0)

		for i := 0 ; i < L ; i++ {
			exitCount := 0
			for _, node := range (*graph.GetEdges())[i] {
				if isOneOf(node, exits) {
					exitCount += 1
				}
			}
			if exitCount == 2 {
				mostDangerousNodes = append(mostDangerousNodes, i)
			}
		}

		fmt.Fprintln(os.Stderr, "most dangerous path: ", mostDangerousNodes)

		if emergency, node := isEmergency(SI, exits, graph); emergency {
			fmt.Fprintln(os.Stderr, "Emergency!")
			nodeA = SI
			nodeB = node
		} else {
			fmt.Fprintln(os.Stderr, "Not Emergency!")
			path, distances := dijsktra(graph, SI, exits)
			fmt.Fprintln(os.Stderr, "path: ", path)
			fmt.Fprintln(os.Stderr, "distance: ", distances)

			if len(mostDangerousNodes) == 0 {
				for _, exit := range exits {
					for _, node := range (*graph.GetEdges())[exit] {
						nodeA = node
						nodeB = exit
					}
				}
			}


			distDanger := make([]int, 0)
			minDist := math.MaxInt64

			for _, dNode := range mostDangerousNodes {
				distDanger = append(distDanger, distances[dNode])
				if distances[dNode] < minDist {
					minDist = distances[dNode]
					nodeA = dNode
				}
			}

			for _, node := range (*graph.GetEdges())[nodeA] {
				if isOneOf(node, exits) {
					nodeB = node
				}
			}

		}


		result := fmt.Sprintf("%d %d",nodeA, nodeB)
		// Example: 3 4 are the indices of the nodes you wish to sever the link between
		fmt.Println(result)

		graph.RemoveEdge(nodeA, nodeB)
		graph.RemoveEdge(nodeB, nodeA)

	}
}

func isEmergency(agent int, exits []int, graph Graph) (bool, int) {
	for _, node := range (*graph.GetEdges())[agent] {
		if isOneOf(node, exits) {
			return true, node
		}
	}
	return false, -1
}

func isOneOf(one int, ones []int) bool {
	for _, o := range ones {
		if o == one {
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

	if past[len(past)-1] == start {
		*multiPath = append(*multiPath, past)
	}
	past = past[:len(past)-1]
}

func dijsktra(graph Graph, start int, exits []int) (map[int][]int, map[int]int) {
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

		for _, node := range (*graph.GetEdges())[minNode] {
			newDistance := distance[minNode] + graph.GetDistance(Edge{minNode, node}, exits)
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
	end   int
}

type Graph struct {
	nodes []bool
	edges map[int][]int
	distances map[Edge]int
}

func InitGraph(n int) Graph {
	graph := Graph{
		make([]bool, n+1),
		make(map[int][]int),
		make(map[Edge]int),
	}

	return graph
}

func (g *Graph) AddEdge(start int, to int, distance int) {
	g.nodes[start] = true
	g.nodes[to] = true

	isExist := false
	for _, oldTo := range g.edges[start] {
		if to == oldTo {
			isExist = true
		}
	}
	if !isExist {
		g.edges[start] = append(g.edges[start], to)
	}

	g.distances[Edge{start, to}] = distance
}

func (g *Graph) RemoveEdge(start int, to int) {
	for i, v := range g.edges[start] {
		if to == v {
			g.edges[start][i] = g.edges[start][len(g.edges[start])-1]
			g.edges[start] = g.edges[start][:len(g.edges[start])-1]
		}
	}
	delete(g.distances, Edge{start, to})
}

func (g *Graph) GetDistance(edge Edge, exits []int) int {
	if isOneOf(edge.start, exits) || isOneOf(edge.end, exits) {
		return 1000
	}

	if isNeighborOf((*g.GetEdges())[edge.start], exits) && !isNeighborOf((*g.GetEdges())[edge.end], exits) {
		return 1
	}

	if isNeighborOf((*g.GetEdges())[edge.end], exits) && !isNeighborOf((*g.GetEdges())[edge.start], exits) {
		return 1
	}

	for _, neighbor := range (*g.GetEdges())[edge.end] {
		if isOneOf(neighbor, exits) {
			return 0
		}
	}

	return g.distances[edge]
}

func isNeighborOf(nodes []int, exits []int) bool {
	for _, ndoe := range nodes {
		if isOneOf(ndoe, exits) {
			return true
		}
	}
	return false
}

func (g *Graph) SetDistance(edge Edge, distance int) {
	g.distances[edge] = distance
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
