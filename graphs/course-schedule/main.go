package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	visited := make([]bool, numCourses)
	path := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		adj[i] = make([]int, 0)
	}

	for i := 0; i < len(prerequisites); i++ {
		a := prerequisites[i][0]
		b := prerequisites[i][1]
		adj[b] = append(adj[b], a)
	}

	var dfs func(node int) bool
	dfs = func(node int) bool {
		visited[node], path[node] = true, true

		for _, next := range adj[node] {
			if !visited[next] && dfs(next) {
				return true
			} else if path[next] {
				return true
			}
		}

		path[node] = false
		return false
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] && dfs(i) {
			return false
		}
	}

	return true
}
