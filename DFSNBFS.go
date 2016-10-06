// https://www.acmicpc.net/problem/1260

package main

import (
	"fmt"
	"sort"
	"strings"
)

type Graph struct {
	nodes []int
	edges map[int][]int
}

func (g *Graph) addNode(node int) {
	for _, v := range g.nodes {
		if v == node {
			return
		}
	}
	g.nodes = append(g.nodes, node)
}

func (g *Graph) addEdge(nA int, nB int) {
	for _, v := range g.edges[nA] {
		if v == nB {
			return
		}
	}
	g.edges[nA] = append(g.edges[nA], nB)
	sort.Ints(g.edges[nA])
}

type MyQueue struct{
	myQueue []int
}

func (q *MyQueue) push(data int) {
	if isInSlice(&q.myQueue, data) == true {
		return
	}
	q.myQueue = append(q.myQueue, data)
}

func (q *MyQueue) pop() int {
	if len(q.myQueue) == 0 {
		return -1
	}

	result := q.myQueue[0]
	q.myQueue = q.myQueue[1:]
	return result
}

func (q *MyQueue) size() int {
	return len(q.myQueue)
}

func (q *MyQueue) empty() int {
	if len(q.myQueue) == 0 {
		return 1
	}
	return 0
}

func (q *MyQueue) front() int {
	if len(q.myQueue) == 0 {
		return -1
	}
	return q.myQueue[0]
}

func (q *MyQueue) back() int {
	if len(q.myQueue) == 0 {
		return -1
	}
	return q.myQueue[len(q.myQueue) - 1]
}



func DFS(g *Graph, v int, discovered *[]int) {
	*discovered = append(*discovered, v)

	for _, node := range g.edges[v] {
		if isInSlice(discovered, node) == false {
			DFS(g, node, discovered)
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

func BFS(g *Graph, v int) []int {
	visit := make([]int, 0)
	queue := MyQueue{make([]int, 0)}
	queue.push(v)
	for {
		if queue.empty() == 1 {
			return visit
		}
		for _, node := range g.edges[queue.front()] {
			if isInSlice(&visit, node) == false {
				queue.push(node)
			}
		}
		visit = append(visit, queue.pop())
	}
}

func main() {
	var n, m, v int
	fmt.Scanf("%d %d %d\n", &n, &m, &v)

	graph := Graph {
		make([]int, 0),
		make(map[int][]int),
	}

	for i := 0 ; i < m; i++ {
		var a, b int
		fmt.Scanf("%d %d\n", &a, &b)
		graph.addNode(a)
		graph.addNode(b)
		graph.addEdge(a, b)
		graph.addEdge(b, a)
	}

	visit := make([]int, 0)
	DFS(&graph, v, &visit)

	fmt.Println(strings.Trim(fmt.Sprint(visit), "[]"))
	fmt.Println(strings.Trim(fmt.Sprint(BFS(&graph, v)), "[]"))
}
