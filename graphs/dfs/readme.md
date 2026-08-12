# About
This is for practicing DFS implementation of graph traversal.

The code should follow these directives: 
```go
// Graph is an adjacency list mapping a node ID to a slice of its neighbors.
type Graph map[int][]int

// DFS performs a depth-first search and returns a slice of node IDs in the 
// exact order they were first visited.
func DFS(graph Graph, start int) []int
```

## Constraints

- Nodes are unique. For example, there could only be a single node `1`. 