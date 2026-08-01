package main

func validTree(n int, edges [][]int) bool {
	adjList := make([][]int, n)

	for i := 0; i < len(adjList); i++ {
		adjList[i] = make([]int, 0)
	}

	var a, b int
	for _, edge := range edges {
		a, b = edge[0], edge[1]
		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a)
	}

	visited := make([]bool, n)
	visitedCount := 0

	//returns true when detecting a cycle
	var dfs func(node, parent int) bool
	dfs = func(node, parent int) bool {
		visited[node] = true
		visitedCount++

		for _, next := range adjList[node] {
			if next == parent {
				continue
			}

			if visited[next] || dfs(next, node) {
				return true
			}
		}

		return false
	}

	//cycle detected
	if dfs(0, -1) {
		return false
	}

	//We didn't visit every node so the graph must be disconnected
	if visitedCount != n {
		return false
	}

	return true
}
