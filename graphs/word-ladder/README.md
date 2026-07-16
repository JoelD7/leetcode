# Problem
https://leetcode.com/problems/word-ladder/description/

A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:

    Every adjacent pair of words differs by a single letter.
    Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
    sk == endWord

Given two words, beginWord and endWord, and a dictionary wordList, return the number of words in the shortest transformation sequence from beginWord to endWord, or 0 if no such sequence exists.


### Example 1:

    Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
    Output: 5
    Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.

### Example 2:

    Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
    Output: 0
    Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.


### Constraints:

    1 <= beginWord.length <= 10
    endWord.length == beginWord.length
    1 <= wordList.length <= 5000
    wordList[i].length == beginWord.length
    beginWord, endWord, and wordList[i] consist of lowercase English letters.
    beginWord != endWord
    All the words in wordList are unique.

# Solution
### Rationale

At it’s most basic level we’re being asked to find the shortest path between two nodes. This is a problem that can be modeled using BFS. Transform the “transformation sequence” to a graph where an adjacent node is a node that only differs in 1-letter. Likewise, a node that differs in 2-letters, is 2 levels apart. After the graph is built with these rules in mind, do BFS from `beginWord` node to `endWord`. Since BFS produces one of shortests paths between two nodes, you’ll get the minimum steps required to transform `beginWord` in `endWord`.

### Algorithm

To build/identify adjacent nodes/words we use a “pattern”, in which we swap a single letter of every word into a flag like “*”, so we can easily group words that share the same pattern. For example, the word `"hot"` can be transformed into `"*ot"`, `"h*t"`, and `"ho*"`. We can see that by applying this pattern logic, any other word that shares the same pattern will be adjacent to “hot”(dot, hat, hog, etc.). The idea here is that when using BFS we can quickly identify adjacent nodes by matching their patterns.

An important thing to note is that the graph **will not be built using only the elements of `wordList` as nodes**, but also the aforementioned patterns.

The next thing is doing BFS over the graph, starting obviously with `beginWord`. Remember that doing BFS on a graph means using a queue to enqueue the *adjacent* elements on the current node to later process them in order. Only after those particular elements have been processed, it’s considered that a single level has been processed. This is why we record the size of the queue in a variable `size` *before* the processing starts: to be able to separate the nodes of the previous level, from the nodes that are going to be enqueued to form the next level.

As previously mentioned, the *adjacent elements are identified by the pattern*. So we’ll enqueue those that have the same pattern as adjacents. When we find a word that matches `endWord`, we will have reached our shortest path, so we return the level.

### Variables

- `curWord`: the current word being processed during our BFS traversal
- `lvl`: current level in the BFS traversal. Indicates how far a part are we from the starting node(`beginWord`).

### Implementation

1. Check if `endWord` is on `wordList` using a map. Note that `endWord` doesn’t have to be the last word in `wordList`, despite what the problem description might imply. This is why we use a map for checking this.
2. Add `beginWord` to `wordList` to make sure this word is in our graph and we’re able to effectively do BFS. The problem states that it might not necessarily be on `wordList`.
3. Iterate over all the words on `wordList` to initialize our graph with the patterns. Every word will be adjacent to a node representing any of the patterns that can be formed with that word. As such, a word like “hot” will be adjacent to nodes “\*ot”, “h\*t”, “ho*”. This will allow us to group other words that match those same patterns and make them adjacent to each other. For example, a word like “hot” will be adjacent to “hog” because they both share the same “ho*” pattern.
4. Start BFS
    1. `lvl` begins at 1 because he problem counts the number of words in the sequence, so beginWord counts as 1.
    2. When the dequeued word matches `endWord`, we have reached the end and must return the amount of levels traveled so far
    3. Enqueue all the adjacent nodes of `curWord`. As mentioned previously, all the adjacent nodes of a node are the ones that share the same pattern, because we established that a pattern only differs by one letter
        1. We only do this with unvisited nodes to prevent infinite loops
        2. Before enqueue we mark the node as visited