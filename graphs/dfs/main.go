package main

type Graph map[int][]int

func DFS(graph Graph, start int) []int {
	res := make([]int, 0)
	visited := make(map[int]bool)

	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true

		res = append(res, node)
		for _, neighbour := range graph[node] {
			if visited[neighbour] {
				continue
			}

			dfs(neighbour)
		}
	}

	dfs(start)

	return res
}
