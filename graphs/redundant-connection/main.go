package main

func findRedundantConnection(edges [][]int) []int {
	N := len(edges)
	adj := make([][]int, N)
	parent := make([]int, N)

	for i := 0; i < N; i++ {
		adj[i] = make([]int, 0)
		parent[i] = -1
	}

	visited := make([]bool, N)
	cycleStart := -1

	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true

		for _, next := range adj[node] {
			if !visited[next] {
				parent[next] = node
				dfs(next)
			} else if next != parent[node] && cycleStart == -1 {
				parent[next] = node
				cycleStart = next
			}
		}
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	dfs(0)

	node := parent[cycleStart]
	cycleNodes := map[int]struct{}{
		cycleStart: {},
	}

	for node != cycleStart {
		cycleNodes[node] = struct{}{}
		node = parent[node]
	}

	for i := N - 1; i >= 0; i-- {
		u, v := edges[i][0]-1, edges[i][1]-1
		_, hasU := cycleNodes[u]
		_, hasV := cycleNodes[v]

		if hasU && hasV {
			return edges[i]
		}
	}

	return []int{}
}
