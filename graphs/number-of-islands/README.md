# Problem
https://leetcode.com/problems/number-of-islands/description/

Given an `m x n` 2D binary grid `grid` which represents a map of `1`s (land) and `0`s (water), return the number of islands.

An **island** is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.


### Example 1:

    Input: grid = [
    ["1","1","1","1","0"],
    ["1","1","0","1","0"],
    ["1","1","0","0","0"],
    ["0","0","0","0","0"]
    ]
    Output: 1

### Example 2:

    Input: grid = [
    ["1","1","0","0","0"],
    ["1","1","0","0","0"],
    ["0","0","1","0","0"],
    ["0","0","0","1","1"]
    ]
    Output: 3


### Constraints:

    m == grid.length
    n == grid[i].length
    1 <= m, n <= 300
    grid[i][j] is '0' or '1'.

# Solution
### Concepts

- `land`: a grid cell with the value of 1
- `water`: a grid cell with the value of 0. Note that everything outside the grid’s bound is also water, as per the problem’s description

### Rationale

The key to this solution is two-fold:

1. Grouping as many pieces of adjacent land as possible to form *a single island*.
2. Use a piece of land as an island *once*. After a piece of land is “used” in a island group, we can’t use it again.

To achieve the first objective we can see that DFS presents a behaviour similar to what we need: go as deep as possible into a chain of adjacent nodes until a condition is met, then we backtrack. So we’ll use DFS to expand our islands to their maximum size.

Beyond that we need a way to avoid using the same piece of land in separate islands. A valid approach could be using a `visited` array or map to hold the already used pieces. A better approach is just converting the visited cells to “0”, signalling they’ve been used and avoiding the need of using extra space for the algorithm.

When an island can no longer be expanded, we’ll look for other non-water cells in our grid and repeat the process again, increasing our island count every time we find a new piece of land.

### Variables

- `islands`: total island count. Return value of the function
- `dfs`: recursive function used to explore adjacent cells in order to know whether the current island can be expanded or not.

### Algorithm

1. Iterate over every cell
    1. If the cell is 0, continue
    2. When you find a 1:
        1. increase `islands` by 1
        2. mark the cell as “visited” by turning it into a 0. Doing this will prevent us from processing the same cells several times, without the need of using extra space in the form of a `visited` array.
        3. Do DFS on every direction of that cell(up, down, left, right). The purpose of this recursive search is to know whether the current island being formed can be expanded or not.
            1. If the current cell is outside bounds or **it’s 0(water)**, return. We can’t keep on expanding when we find a 0 because an island is “expandable” as long as there are adjacent 1’s.
            2. Else, mark the current cell as visited by turning it into a 0 and invoke `dfs` for the adjacent cells in every direction