package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make([][]int, numCourses)
	for i := 0; i < len(adj); i++ {
		adj[i] = make([]int, 0)
	}
	var a, b int

	for i := 0; i < len(prerequisites); i++ {
		a, b = prerequisites[i][0], prerequisites[i][1]
		adj[b] = append(adj[b], a)
	}

	stack := make([]int, 0)
	visited := make([]bool, len(adj))
	path := make([]bool, len(adj))

	var dfs func(node int) bool //returns true when finding a cycle
	dfs = func(node int) bool {
		visited[node], path[node] = true, true

		for _, neighbor := range adj[node] {
			if !visited[neighbor] && dfs(neighbor) {
				return true
			} else if path[neighbor] {
				return true
			}
		}

		path[node] = false
		stack = append(stack, node)

		return false
	}

	for i := 0; i < len(adj); i++ {
		if !visited[i] && dfs(i) { //return empty array when finding a cycle
			return []int{}
		}
	}

	result := make([]int, 0)
	for i := len(stack) - 1; i >= 0; i-- {
		result = append(result, stack[i])
	}

	return result
}
