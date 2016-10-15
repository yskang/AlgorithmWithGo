package myLibs

import (
	"math"
)

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
		minNode := 9999
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
