# Problem
https://leetcode.com/problems/combination-sum


Given an array of distinct integers `candidates` and a target integer `target`, return a list of _all unique combinations_ of `candidates` where the chosen numbers sum to `target`. You may return the combinations in **any order**.

The **same** number may be chosen from `candidates` an **unlimited number of times**. Two combinations are unique if the frequency(The frequency of an element x is the number of times it occurs in the array.) of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to `target` is less than `150` combinations for the given input.


### Example 1:

    Input: candidates = [2,3,6,7], target = 7
    Output: [[2,2,3],[7]]
    Explanation:
    2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
    7 is a candidate, and 7 = 7.
    These are the only two combinations.

### Example 2:

    Input: candidates = [2,3,5], target = 8
    Output: [[2,2,2,2],[2,3,3],[3,5]]

### Example 3:

    Input: candidates = [2], target = 1
    Output: []


### Constraints:

    1 <= candidates.length <= 30
    2 <= candidates[i] <= 40
    All elements of candidates are distinct.
    1 <= target <= 40

# Solution
### Variables

- `res`. The list of all combinations.
- `comb`. The current combination of candidates
- `sum`. The sum of all the elements of `comb`
- `candidate`. The current candidate we’re evaluating, i.e., `candidates[i]`.

### Algorithm

- Use a backtrack function that takes `comb`and `sum`.
- Iterate though `candidates` and on each new iteration
    1. Add `candidate` to `comb` and update`sum`
    2. Call `backtrack`.
        1. Check if the base case has been reached. The base is reached in two ways:
            1. `sum == target`. In that case we add `comb` to `res` because we have a valid combination.
            2. `sum > target`. We return without adding `comb` to `res` because this combination is not valid.
    3. After the recursive call to `backtrack` returns, remove `candidate` from `comb` and `sum` in order to try with a new candidate.

The loop begins at `start` because we want to ensure that past candidates don’t repeat themselves in the final solution. For example, for `candidates = [1,2,3]`, having combinations `[2,3],[3,2]`.

The value of `start` will always be the index of the candidate we’re currently evaluating in the iteration. We do this because according to the problem’s constraints is possible for a same candidate to appear several times as a valid combination. For example, for `candidates = [2,3,6,7], target = 7` a valid combination is `[2,2,3]`.
 