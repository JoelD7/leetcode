# Problem
https://leetcode.com/problems/combination-sum-ii/description/

Given a collection of candidate numbers (`candidates`) and a target number (`target`), find all unique combinations in `candidates` where the candidate numbers sum to `target`.

Each number in `candidates` may only be used once in the combination.

**Note**: The solution set must not contain duplicate combinations.


### Example 1:

    Input: candidates = [10,1,2,7,6,1,5], target = 8
    Output:
    [
    [1,1,6],
    [1,2,5],
    [1,7],
    [2,6]
    ]

### Example 2:

    Input: candidates = [2,5,2,1,2], target = 5
    Output:
    [
    [1,2,2],
    [5]
    ]


### Constraints:

    1 <= candidates.length <= 100
    1 <= candidates[i] <= 50
    1 <= target <= 30

# Solution
### Variables

- `res`. The list of all combinations.
- `comb`. The current combination of candidates
- `sum`. The sum of all the elements of `comb`
- `candidate`. The current candidate we’re evaluating, i.e., `candidates[i]`.

### Algorithm

- Sort the `candidates` array.
- Use a backtrack function that takes `comb`, `sum`, `candidates`, `res`, `target` and `start`, the index at which the next iteration of the recursive call will begin.
- Iterate though `candidates` and on each new iteration
    1. Continue with the next iteration if the current `candidate` is equal to the previous candidate **AND** if `i > start`.
        1. The problem states that “e*ach number in `candidates` may only be used **once** in the combination*”. Here we’re talking about “uniqueness” in terms of the domain of `candidates` array, not of the whole natural numbers domain. That means that if the number `1` appears two times in `candidates`, a valid combination may have two number `1`'s, but not more. For that reason we only skip equal candidates if `i > start`, because remember, each candidate is visited several times by either the loop iteration or the recursive call. By setting the `i > start` condition we’re ensuring that only the loop is in charge of adding valid repeated numbers to the combination.
       2. **Why i > start instead of i > 0?** Each recursive call is in charge of processing a portion of the candidates array, specifically, `candidates[start:]`. In other words, the domain of a recursive call is on the range `[start:]`. We want _some_ duplication to allow for combinations such as `[1,1,6]` but not too much duplication to prevent repeated combinations such as a double `[1,7]`. If I do `i > 0` I won’t allow duplicates at all because `i` will never be `< 0`, and it will be `== 0` **only once**, so that is out of the question. Doing `i > start` strikes the perfect balance  because duplicates will be forbidden in the context of the current recursive call, but they will be allowed between recursive calls, handled by the higher-up loop that made this recursive call.
    2. Add `candidate` to `comb` and update`sum`
    3. Call `backtrack` setting `start` as `i + 1`
        1. Check if the base case has been reached. The base is reached in two ways:
            1. `sum == target`. In that case we add `comb` to `res` because we have a valid combination.
            2. `sum > target`. We return without adding `comb` to `res` because this combination is not valid.
    4. After the recursive call to `backtrack` returns, remove `candidate` from `comb` and `sum` in order to try with a new candidate.

The value of `start` will always be the **next** index of the candidate we’re currently evaluating in the iteration. We do this because one of the problem’s restrictions is that each number in `candidates` may only be used **once** in the combination.
 
