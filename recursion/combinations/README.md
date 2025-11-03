# Problem
https://leetcode.com/problems/combinations/description/

Given two integers `n` and `k`, return all possible combinations of `k` numbers chosen from the range `[1, n]`.

You may return the answer in any order.


### Example 1:

    Input: n = 4, k = 2
    Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
    Explanation: There are 4 choose 2 = 6 total combinations.
    Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.

### Example 2:

    Input: n = 1, k = 1
    Output: [[1]]
    Explanation: There is 1 choose 1 = 1 total combination.

### Constraints:

    1 <= n <= 20
    1 <= k <= n

# Solution
- `res`: holds the list of combinations we find during the algorithm. This is the return value of the function.
- `comb`: This is the current combination. We’ll use this as a temp variable to build the combinations.

The solution uses a `backtrack` function that creates the combinations. It’s most important parameter is `start`, that indicates the start point of the items that can be added to the current combination. We need this to avoid repeating numbers.

The **base case is reached** when the amount of items in the current combination is `k`, which is the constraint set by the problem. At that point we add the combination to `res`.

Next, we have a loop goes from `start` upto `n`. Inside the loop we start building the new combination by adding the current number `num` to the array. Then we recursively call `backtrack` passing `num + 1` as `start`, because we don’t want the next number in the combination to be the same as the one we just added. The process will repeat recursively until we reach the base case.

After the recursive call we remove the last added item from `comb` so we can form a different combination on the next iteration. There is no point in keeping this particular item because of several reasons. First, it is the last item of a valid combination. We are certain of this because the only way that a recursive call returns is when a valid combination is formed. Second, the problem doesn’t allow for repeated combinations. Third, we want to form all the possible combinations so we have to include all the numbers upto `n`.
