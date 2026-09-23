package main

type Graph map[int][]int

func topologicalSort(graph Graph, start int) []int {
	res := make([]int, 0)
	stack := make([]int, 0)
	visited := make(map[int]bool)

	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true

		for _, next := range graph[node] {
			if !visited[next] {
				dfs(next)
			}
		}

		stack = append(stack, node)
	}
	dfs(start)
	for i := len(stack) - 1; i >= 0; i-- {
		res = append(res, stack[i])
	}
	return res
}
