package main

type Graph map[int][]int

func DFS(graph Graph, start int) []int {
	visited := make([]bool, len(graph))
	result := make([]int, 0)

	var dfs func(node int)
	dfs = func(node int) {
		if visited[node] {
			return
		}

		visited[node] = true
		for _, neighbors := range graph[node] {

		}
	}
}
