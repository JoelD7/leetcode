# Problem
https://leetcode.com/problems/word-search-ii/description/

Given an `m x n` board of characters and a list of strings words, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.



### Example 1:

    Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
    Output: ["eat","oath"]

### Example 2:

    Input: board = [["a","b"],["c","d"]], words = ["abcb"]
    Output: []



### Constraints:

    m == board.length
    n == board[i].length
    1 <= m, n <= 12
    board[i][j] is a lowercase English letter.
    1 <= words.length <= 3 * 104
    1 <= words[i].length <= 10
    words[i] consists of lowercase English letters.
    All the strings of words are unique.

# Solution
### Rationale

The solution uses a combination of backtracking and a trie to efficiently find words in the board. It picks a cell and goes as deep as possible until finding the word or backtracking. The approach is perfect because thanks to the trie, we can realize inmediately and efficiently that a particular path through the board doesn’t have a valid word, allowing us to prune it right away.

### Algorithm

1. The first thing we do is bulding the `trie` with all the `words`. We save each fully formed word in a leaf node of the trie. We’ll use that to identify whether we’ve found a `word` or not.
    1. Each node of the `trie` is of type `TrieNode`, which will only hold it’s children and a string `word`, denoting if it’s a word node or prefix node.
    2. The `trie` is built using a utility function: `buildTrie()`
2. Iterate over all the cells in the board and call the `dfs()` function for each character `c`
    1. If the current `TrieNode` doesn’t have a child with `c`, return. This indicates that either we are on the wrong path or that this character is not on the `board`/`trie`.
    2. If `c == #`, return. The “#” is an indicator we use to denote that a cell has been visited in the current path being explored. We also return in this case to avoid repeating characters in the same path.
    3. If `c` exists in the TrieNode, we move deeper to continue our search. If the node is a word, we update our `res` array and set that node to nil to prevent adding it multiple times.
    4. Update `board[i][j` to be “#”. Again, since we plan to go deeper down this path in multiple directions, we mark the current cell as “visited” to prevent duplicates. Instead of using extra space with a “visited” array, we use the same board for this.
    5. Call `dfs()` on all four directions from the current cell to further expand the search
    6. After all the 4 calls return, re-store `board[i][j]` to it’s original value of `c`. At this point we’ll have explored all paths starting at this cell so we need to free it so that other paths may use it.