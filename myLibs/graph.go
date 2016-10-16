package myLibs

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
