package utils

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