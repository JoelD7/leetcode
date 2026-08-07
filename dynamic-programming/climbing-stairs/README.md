# Problem
https://leetcode.com/problems/climbing-stairs/description/

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?



### Example 1:

    Input: n = 2
    Output: 2
    Explanation: There are two ways to climb to the top.
    1. 1 step + 1 step
    2. 2 steps

### Example 2:

    Input: n = 3
    Output: 3
    Explanation: There are three ways to climb to the top.
    1. 1 step + 1 step + 1 step
    2. 1 step + 2 steps
    3. 2 steps + 1 step



### Constraints:

    1 <= n <= 45

# Solution
### Rationale

I see this as a dynamic programming problem because the total amount of ways we can climb a stair of `n` steps, is the sum of the total # of ways of climbing `n-1` steps + `n-2` steps + …. + `1` step. Or…

```markdown
climbStairs(n) = climbStairs(n-1) + climbStairs(n-2) + ... + climbStairs(1)
```

This clearly is a *top-down* approach where we progresively decompose the problem into smaller parts and the solution of those parts builds the solution of the problem.

**What counts as a “way” to climb `n` steps?** We subtract 1 or 2 steps from `n`, and every time `n` reaches 0, you count one way. This exhibits the characteristics of a backtracking problem where we must figure out all possible combinations of doing something, therefore, we use the same technique of loop+recursion to achieve that goal. In the loop we iterate over the available steps, subtract each step from `n` and initiate recursion. Doing this at every recursive call means we’ll get all possible ways to step over `n`.

During the recursive calls we’ll reach the same `n` values multiple times, so we can make this algorithm a lot more efficient by memoizing it. Every time we reach a previously calculated `n`, we return the ways we can step over it instead of calculating again. For `n = 44`, memoizing makes the program to run in 200ms instead of **21 seconds!**