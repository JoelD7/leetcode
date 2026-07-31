# Problem
https://leetcode.com/problems/walls-and-gates/

You are given a `m × n` 2D `grid` initialized with these three possible values:

    -1 - A water cell that can not be traversed.
    0 - A treasure chest.
    INF - A land cell that can be traversed. We use the integer 2^31 - 1 = 2147483647 to represent INF.

Fill each land cell with the distance to its nearest treasure chest. If a land cell cannot reach a treasure chest then the value should remain INF.

Assume the grid can only be traversed up, down, left, or right.

Modify the grid in-place.

### Example 1:

    Input: [
    [2147483647, -1,            0,          2147483647],
    [2147483647, 2147483647,    2147483647, -1],
    [2147483647, -1,            2147483647, -1],
    [0,          -1,            2147483647, 2147483647]
    ]
    
    Output: [
    [3,-1,0,1],
    [2,2,1,-1],
    [1,-1,2,-1],
    [0,-1,3,4]
    ]

### Example 2:

    Input: [
    [0,-1],
    [2147483647,2147483647]
    ]
    
    Output: [
    [0,-1],
    [1,2]
    ]

### Constraints:

    m == grid.length
    n == grid[i].length
    1 <= m, n <= 100
    grid[i][j] is one of {-1, 0, 2147483647}

# Solution
### Rationale

The problem can be modeled as a BFS where we perform BFS from all treasure cells simultaneously, updating each land cell with the distance they are from the treasure cell. Instead of starting the search from every land cell upto treasure cells, we do the opposite as it is more efficient(there are less treasure cells than land cells).

Also, we do BFS from all treasure cells at the same time to make the program even faster. This technique is called “multi-source BFS”.

It has to be noted that when we find a wall(-1 cell) we can’t continue going down that path, i.e., we can’t enqueue that cell because that would imply we move through the wall and that is not valid.

### Algorithm

1. Enqueue all treasure cells.
    1. This is not the usual way of doing BFS but remember that we aren’t doing conventional BFS. We are implementing a multi-source BFS where the “search” part starts from multiple nodes simultaneously.
2. While the queue has elements we do the following:
    1. Dequeue a `node`. Note that all the treasure cells will be processed first as we enqueued them on the previous step
    2. Move through every 4-directionally adjacent cell of `node` to verify if there is a land cell that should be updated.
        1. Check if the coordinates are inside the bounds of the grid AND more importantly, if the adjacent cell is *untouched* land, i.e., if it has a value of 2147483647. If it is, use the level of it’s parent to set the level(distance) of this cell and enqueue it so that you’re able to do the same with it’s children.

           Here we got to stop and explain a little further. Note the word “untouched”, this means a land cell that has never been updated. Shouldn’t we also update cells that are already touched as well? No because with this multi-source BFS approach we do a **single** pass over the entire grid, so every cell is updated exactly once. The way we assign a distance of cells that are more than one level apart from a treasure cell is by checking the value of it’s parent cell. Here is the key part:

            ```go
            //r,c = dequeued cell
            //x, y = 4-directionally adjacent cell to r,c
            grid[x][y] = grid[r][c] + 1
            queue = append(queue, [2]int{x, y})
            ```

           On the first pass `r,c` is a treasure cell so the adjacent cell will have a value of 1 because is adjacent to the treasure. All the adjacent cells to a treasure chest will have a value of 1, so we can say they’re 1 level away from the start. Then, the adjacent cells to those 1’s should be 2 levels apart from the start or 1 level more away from their parents. The adjacents to those 2’s should be 3 levels away, and so on. We can see here that to set the level of a cell, all we need to do is adding 1 to the level of its parent. Therefore, we don’t need to revisit any cells. The set value of a land cell is final because we started counting from the treasure cell, so we already have the closest distance to a treasure, which is precisely what the problem is asking. This is another reason to start the BFS from the treasure cells: it makes the algorithm a lot simpler.