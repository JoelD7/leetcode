# Problem
https://leetcode.com/problems/last-stone-weight/

You are given an array of integers stones where stones[i] is the weight of the ith stone.

We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash them together. Suppose the heaviest two stones have weights x and y with x <= y. The result of this smash is:

    If x == y, both stones are destroyed, and
    If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.

At the end of the game, there is at most one stone left.

Return the weight of the last remaining stone. If there are no stones left, return 0.



### Example 1:

    Input: stones = [2,7,4,1,8,1]
    Output: 1
    Explanation:
    We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
    we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
    we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
    we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.

### Example 2:

    Input: stones = [1]
    Output: 1



### Constraints:

    1 <= stones.length <= 30
    1 <= stones[i] <= 1000

# Solution
Use a max-heap and always take the root and its largest child as the two stones: `x` and `y`. You should always *remove* the two largest stones out of the max-heap. Regardless of the stone smashing outcome, those two values will no longer be used. We will either add the subtraction to the max-heap(if one stone was larger than the other) or not add anything. In any case `x` and `y` have to be removed.

1. Iterate over the max-heap until it has a single element
    1. Remove the largest element(root, which we’ll call `x`) and its largest child(`y`).
    2. If `x == y`, continue. The stones were already removed in the previous step.
    3. if `x != y`, add `math.Abs(x - y)` to the max-heap.
2. Return the single element of the heap.