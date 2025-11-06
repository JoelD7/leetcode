# Problem
https://leetcode.com/problems/subsets/description/

Given an integer array `nums` of unique elements, return all possible
subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.


### Example 1:

    Input: nums = [1,2,3]
    Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

### Example 2:

    Input: nums = [0]
    Output: [[],[0]]


### Constraints:

    1 <= nums.length <= 10
    -10 <= nums[i] <= 10
    All the numbers of nums are unique.

# Solution
### Variables

- `comb`: the combination we’re forming
- `res`: the result of the function, which is, the list of subsets
- `backtrack`: the recursive backtracking function we’ll use to form the combinations. It’s parameters are:
    - `start`: the index at which we’ll begin iterating over `nums`
    - `nums`: the input array
    - `comb`: A pointer to the original `comb`. Needs to be a pointer to be able to modify it across recursive calls
    - `res`: A pointer to the original `res`. Needs to be a pointer to be able to modify it across recursive calls

### Algorithm

1. **Add `comb` to `res`**.
    1. The first thing we do at the start of the `backtrack` function is adding `comb` to the list of combinations. We do this here because it allows us to add the empty set to `res` without additional code, and separates the logic that creates combinations from the logic that saves them. Updating `res` inside the loop makes the code behave unexpectedly.
    2. There is no condition for including `comb` in `res` because the problem is asking for any set of elements(albeit not repeated) to be counted as a combination.
2. **Iterate over `nums`, from `start`**.
    1. Next, we start creating combinations by moving through `nums` and beginning at `start`index. We first add the current element to the current combination and then make the recursive call. Every time we make a recursive call, the `start` index will be moved to the right to avoid repeating the same numbers in the combinations.
    2. On this algorithm, combinations are formed every time we make a recursive call because the only  restriction on the combination size is the size of `nums`.
    3. After we `start` reaches the end of `nums` we start to backtrack. Backtracking is done by eliminating the last added element to the current combination to make space for a new element which will be added on the next iteration.
