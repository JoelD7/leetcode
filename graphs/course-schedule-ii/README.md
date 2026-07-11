# Problem
https://leetcode.com/problems/course-schedule-ii/

There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

    For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.



### Example 1:

    Input: numCourses = 2, prerequisites = [[1,0]]
    Output: [0,1]
    Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].

### Example 2:

    Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
    Output: [0,2,1,3]
    Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
    So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

### Example 3:

    Input: numCourses = 1, prerequisites = []
    Output: [0]



### Constraints:

    1 <= numCourses <= 2000
    0 <= prerequisites.length <= numCourses * (numCourses - 1)
    prerequisites[i].length == 2
    0 <= ai, bi < numCourses
    ai != bi
    All the pairs [ai, bi] are distinct.

# Solution
### Observations

- The description says “*If it is impossible to finish all courses, return **an empty array**.*” In this context I believe “impossible” means cycles in the dependency graph

### Rationale

The problem can be modeled as a dependency graph because this is what pre-requisites and courses form between each other. A dependency graph that requires sorting, is precisely what topological sort was built for. Therefore, the topological sort of this dependency graph is the output of this problem.

While building the topological sort, we’ll also keep track of `visited` nodes and nodes that have been visited in the current DFS path(or recursive call stack) we’re moving through in the array `path`. When detecting a node in `path`, we’ll know we have detected a cycle, which means the user cannot possibly complete all courses because there is a circular dependency in the courses. In that case we return the empty array inmediately.

### Variables

- `visited`: array that has the visited nodes. A node is visited if we have fully explored its entire path.
- `path`: array that has the visited nodes of the current path we’re exploring.
- `adj`: adjacency list used to model the pre-requisite courses as a dependency graph. It is modeled such that the list `adj[a]` includes all the courses that have `a` as pre-requisite.
- `dfs`: function to perform DFS traversal in the graph to both build the topological sort and detect cycles. Returns `true` when it finds a cycle
- `stack`: holds the reversed order of vertices. After visiting all vertices that point to a vertex X, we’ll add X to the stack
- `result`: return value of the function.

### Algorithm

1. Initialize the `adj` matrix. The length of the matrix is naturally the number of vertices(courses) in the graph.
    1. Iterate over pre-requisites to build it
2. Iterate over all the vertices of the graph
    1. If the vertex hasn’t been visited, do DFS over it
        1. If we detect a cycle, return empty array because a cycle means there is a circular dependencies between the courses’ pre-requisites, so is impossible to finish all courses.
        2. Inside `dfs`, mark the node as visited and as in the current path
        3. we’ll move through a node’s adjacency list and if a neighbor hasn’t been visited and `dfs` returns `true`, it means we have found a cycle, so return `true`.
            1. Remember, we use two state arrays `visited` and `path` to distinguish between a node that has been visited in another DFS path vs one that has been visited in the *current* DFS path. Logically, when you visit the a node X in the same DFS recursion path twice, it’s because you have found a cycle. We add the `!visited[neighbor]` condition because that neighbor might have been visited, but in another path! Therefore it shouldn’t be considered a cycle. It’s a way to avoid false positives.
        4. Mark the node as “not visited” in the `path` array, because we just completed all paths that included this node. This is also another way to prevent false positives.
        5. Add the node to `stack`. At this point we’ll have visited all vertices that lead up to node, in other words, all the pre-requisite courses, so we can add it to the stack to denote its ordering in relation to the other nodes.
3. Return the reversed order of `stack` as the output of the function. Naturally, vertices are ordered in the stack such that the dependant nodes appear before their dependencies. Since the problem is asking us “*the ordering of courses you should take to finish all courses”* we must invert the stack’s order.