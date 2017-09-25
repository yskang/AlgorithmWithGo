package utils

import (
	"math"
)

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
