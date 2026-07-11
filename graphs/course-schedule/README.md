# Problem
https://leetcode.com/problems/course-schedule/description/

There are a total of numCourses courses you have to take, labeled from 0 to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [ai, bi]` indicates that you must take course bi first if you want to take course ai.

    For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

Return true if you can finish all courses. Otherwise, return false.



### Example 1:

    Input: numCourses = 2, prerequisites = [[1,0]]
    Output: true
    Explanation: There are a total of 2 courses to take.
    To take course 1 you should have finished course 0. So it is possible.

### Example 2:

    Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
    Output: false
    Explanation: There are a total of 2 courses to take.
    To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.



### Constraints:

    1 <= numCourses <= 2000
    0 <= prerequisites.length <= 5000
    prerequisites[i].length == 2
    0 <= ai, bi < numCourses
    All the pairs prerequisites[i] are unique.

# Solution
### Rationale

This can be modeled like a graph because it’s basically a dependency graph. When we detect a cycle we should return `false` inmediately because it’s impossible to finish courses with circular dependencies. Otherwise, return `true`.

We’ll do DFS to explore the paths of every node, one by one, until we find a circular dependency. The main idea is that if while traversing a graph you find a node X in the *same DFS path* more than once, it’s because there is a cycle. That is the main principle under which we’ll build our solution.

### Variables

- `visited`: array that has the visited nodes. A node is visited if we have fully explored its entire path.
- `path`: array that has the visited nodes of the current path we’re exploring.
- `adj`: adjacency list used to model the pre-requisite courses as a dependency graph. It is modeled such that the list `adj[b]` includes all the courses that have `b` as pre-requisite.
- `dfs`: function to perform DFS in the graph. Returns `true` when it finds a cycle

### Algorithm

1. Build the adjacency list and initialize the state arrays
2. Iterate over all the courses using `numCourses` and `adj`
    1. If the course/node hasn’t been visited, do DFS on it and return `false` when DFS returns `true`(we have found a cycle)
    2. Inside `dfs`, mark the node as visited and as in the current path
    3. we’ll move through a node’s adjacency list and if a neighbor hasn’t been visited and `dfs` returns `true`, it means we have found a cycle, so return `true`.
        1. Remember, we use two state arrays `visited` and `path` to distinguish between a node that has been visited in another DFS path vs one that has been visited in the *current* DFS path. Logically, when you visit the a node X in the same DFS recursion path twice, it’s because you have found a cycle. We add the `!visited[neighbor]` condition because that neighbor might have been visited, but in another path! Therefore it shouldn’t be considered a cycle. It’s a way to avoid false positives.
    4. Mark the node as “not visited” in the `path` array, because we just completed all paths that included this node. This is also another way to prevent false positives.