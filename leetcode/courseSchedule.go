package leetcode

// CanFinish is a solution of a problem "207. Course Schedule" of leetcode.
func CanFinish(numCourses int, prerequisites [][]int) bool {
	return canFinish(numCourses, prerequisites)
}

// Kahn's algorithm
//
// L ← Empty list that will contain the sorted elements
// S ← Set of all nodes with no incoming edge
// while S is non-empty do
//     remove a node n from S
//     add n to tail of L
//     for each node m with an edge e from n to m do
//         remove edge e from the graph
//         if m has no other incoming edges then
//             insert m into S
// if graph has edges then
//     return error (graph has at least one cycle)
// else
//     return L (a topologically sorted order)
func canFinish(numCourses int, prerequisites [][]int) bool {
	courseGraph := graph{nodes: []int{}, edges: []edge{}}
	courseGraph.addNodes(numCourses)
	courseGraph.addEdges(prerequisites)

	L := make([]int, 0)
	S := nodeQueue{nodes: make([]int, 0)}

	for _, n := range courseGraph.nodes {
		if courseGraph.isStartNode(n) {
			S.pushNode(n)
		}
	}

	for !S.isEmpty() {
		n := S.popNode()
		L = append(L, n)
		for _, m := range courseGraph.getNodesFrom(n) {

			courseGraph.removeEdge(n, m)
			if courseGraph.isStartNode(m) {
				S.pushNode(m)
			}
		}
	}

	if len(courseGraph.getEdges()) != 0 {
		return false
	}
	return true
}

type edge struct {
	from int
	to   int
}

type graph struct {
	nodes []int
	edges []edge
}

func (g *graph) addNodes(n int) {
	for i := 0; i < n; i++ {
		g.nodes = append(g.nodes, i)
	}
}

func (g *graph) addEdges(es [][]int) {
	for _, e := range es {
		g.edges = append(g.edges, edge{from: e[1], to: e[0]})
	}
}

func (g *graph) getNodesFrom(n int) []int {
	ans := make([]int, 0)
	for _, e := range g.edges {
		if e.from == n {
			ans = append(ans, e.to)
		}
	}
	return ans
}

func (g *graph) getNodesTo(n int) []int {
	ans := make([]int, 0)
	for _, e := range g.edges {
		if e.to == n {
			ans = append(ans, e.from)
		}
	}
	return ans
}

func (g *graph) removeEdge(from int, to int) {
	for i, e := range g.edges {
		if e.from == from && e.to == to {
			g.edges = append(g.edges[:i], g.edges[i+1:]...)
		}
	}
}

func (g *graph) getEdges() []edge {
	return g.edges
}

func (g *graph) isStartNode(n int) bool {
	for _, e := range g.edges {
		if e.to == n {
			return false
		}
	}
	return true
}

type nodeQueue struct {
	nodes []int
}

func (s *nodeQueue) pushNode(node int) {
	s.nodes = append(s.nodes, node)
}

func (s *nodeQueue) popNode() int {
	p := s.nodes[0]
	s.nodes = s.nodes[1:]
	return p
}

func (s *nodeQueue) isEmpty() bool {
	if len(s.nodes) == 0 {
		return true
	}
	return false
}

func canFinishUsingDFS(numCourses int, prerequisites [][]int) bool {
	type node int

	type graph struct {
		nodes []node
		edges map[int][]int
	}

	courseGraph := graph{nodes: make([]node, numCourses), edges: make(map[int][]int)}

	for _, prerequisite := range prerequisites {
		courseGraph.edges[prerequisite[1]] = append(courseGraph.edges[prerequisite[1]], prerequisite[0])
	}

	visited := make([]bool, numCourses)
	finished := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if visited[i] == false {
			if !checkLoop(i, &courseGraph.edges, &visited, &finished) {
				return false
			}
		}
	}

	return true
}

func checkLoop(hereNode int, edges *map[int][]int, visited *[]bool, finished *[]bool) bool {
	(*visited)[hereNode] = true
	childrenNodes := (*edges)[hereNode]

	for _, childNode := range childrenNodes {
		if !(*visited)[childNode] {
			if checkLoop(childNode, edges, visited, finished) == false {
				return false
			}
		} else if !(*finished)[childNode] {
			return false
		}
	}
	(*finished)[hereNode] = true
	return true
}
