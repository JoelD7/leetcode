package main

func countComponents(n int, edges [][]int) int {
	adjList := make([][]int, n)
	visited := make(map[int]bool)

	for i := 0; i < n; i++ {
		visited[i] = false
	}

	var a, b int
	for _, edge := range edges {
		a, b = edge[0], edge[1]
		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a)
	}

	var dfs func(node, parent int)
	dfs = func(node, parent int) {
		visited[node] = true

		for _, next := range adjList[node] {
			if next == parent {
				continue
			}

			if !visited[next] {
				dfs(next, node)
			}
		}
	}

	count := 0

	for node, vis := range visited {
		//If this node hasn't been visited on previous DFS traversals, it can only because is another graph, so we
		//increase the count.
		if !vis {
			count++
			dfs(node, -1)
		}
	}

	return count
}
