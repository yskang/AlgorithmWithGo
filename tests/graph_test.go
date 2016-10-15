package tests

import (
	"testing"
	"AlgorithmWithGo/myLibs"
	"fmt"
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

	fmt.Println(path)

	past := make([]int, 0)
	multipath := make([][]int, 0)

	myLibs.GetMultiPath(path, 1, 6, past, &multipath)
	fmt.Println(multipath)

}

