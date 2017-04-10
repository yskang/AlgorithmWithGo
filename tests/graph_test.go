package tests

import (
	"testing"
	"fmt"
	"AlgorithmWithGo/myLibs"
)

func TestGraph(t *testing.T) {
	graph := myLibs.InitGraph()

	t.Run("add node", func(t *testing.T) {
		graph.AddNode(1)
		graph.AddNode(2)
		graph.AddNode(3)
	})
	t.Run("add edge", func(t *testing.T) {
		graph.AddEdge(1, 2, 10)
		graph.AddEdge(2, 1, 10)
		graph.AddEdge(1, 3, 5)
	})

	fmt.Println(graph.GetNodes())
	fmt.Println((*graph.GetNodes())[0])

	nodes := *graph.GetNodes()
	fmt.Println(nodes[1])
}

func TestName(t *testing.T) {
	a := make([]int, 0)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	a = append(a, 4)
	a = append(a, 5)

	a = *myLibs.Remove(2, &a)

	fmt.Println(&a)
	fmt.Println(&a)
	fmt.Println(a)
}

func TestPath(t *testing.T) {
	graph := myLibs.InitGraph()

	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(3)
	graph.AddNode(4)
	graph.AddNode(5)
	graph.AddNode(6)

	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 2)
	graph.AddEdge(1, 4, 2)
	graph.AddEdge(1, 5, 2)
	graph.AddEdge(2, 3, 1)
	graph.AddEdge(3, 6, 1)
	graph.AddEdge(4, 6, 1)
	graph.AddEdge(5, 6, 1)

	path, _ := myLibs.Dijsktra(graph, 1)

	past := make([]int, 0)
	multiPath := make([][]int, 0)

	myLibs.GetMultiPath(path, 1, 6, past, &multiPath)
	fmt.Println(multiPath)
}

func TestMultiPath(t *testing.T) {
	paths := make(map[int][]int)

	paths[1] = []int{}
	paths[2] = []int{1}
	paths[3] = []int{1, 2}
	paths[4] = []int{1}
	paths[5] = []int{1}
	paths[6] = []int{3, 4, 5}

	visit := make([]int, 0)
	DFS(&paths, 6, &visit)
	fmt.Println(visit)

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