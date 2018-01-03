package leetcode

// CanFinish is a solution of a problem "207. Course Schedule" of leetcode.
func CanFinish(numCourses int, prerequisites [][]int) bool {
	return canFinish(numCourses, prerequisites)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
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
