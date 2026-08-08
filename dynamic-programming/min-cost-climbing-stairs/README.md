# Problem
https://leetcode.com/problems/min-cost-climbing-stairs/description/

You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor.



### Example 1:

    Input: cost = [10,15,20]
    Output: 15
    Explanation: You will start at index 1.
    - Pay 15 and climb two steps to reach the top.
      The total cost is 15.

### Example 2:

    Input: cost = [1,100,1,1,1,100,1,1,100,1]
    Output: 6
    Explanation: You will start at index 0.
    - Pay 1 and climb two steps to reach index 2.
    - Pay 1 and climb two steps to reach index 4.
    - Pay 1 and climb two steps to reach index 6.
    - Pay 1 and climb one step to reach index 7.
    - Pay 1 and climb two steps to reach index 9.
    - Pay 1 and climb one step to reach the top.
      The total cost is 6.



### Constraints:

    2 <= cost.length <= 1000
    0 <= cost[i] <= 999

# Solution
Notes taken from: https://leetcode.com/problems/min-cost-climbing-stairs/solutions/476388/4-ways-step-by-step-from-recursion-top-d-cbtt

### Rationale

The first thing to note is the recursive nature of this problem:

```go
mincost(i) = cost[i]+min(mincost(i-1), mincost(i-2))
```

The total minimum cost to reach step $i$ is the cost of step $i$ itself, plus the cheaper of the minimum costs to reach either the step right before it ($i-1$) or two steps before it ($i-2$). In short: **Current Step Cost + Cheaper of the Last Two Step Totals**.

This makes sense as we always want to minimize the cost at *each step* to reach the top.

Another way to put the recurrence relation is saying that: the minimum cost of reaching the top step is the minimum cost of reaching every possible step along the way to the top. Why? Because logically you cannot get to the top unless you go through some steps in the way.

### Approach

Instead of getting to the most optimal solution right away, we’ll walk through several ever more optimized solutoins. In this manner we can better understand why the most optimal solution makes sense.

### Algorithm

We have our recursive relation so let’s establish our base case:

```go
minCost[0] = cost[0]
minCost[1] = cost[1]
```

The min cost of the first two steps is the cost of the steps themselves. This is clear as we don’t have to take into consideration any other steps to reach them. **The base case applies to all of the following solutions**.

### Solution 1: Convert recurrence relation to recursion

This one is a direct translation of the recurrence relation/formula described above.

```go
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	return min(minCost(cost, n-1), minCost(cost, n-2))
}

func minCost(cost []int, n int) int {
	if n < 0 {
		return 0
	}

	if n == 1 || n == 0 {
		return cost[n]
	}

	return cost[n] + min(minCost(cost, n-1), minCost(cost, n-2))
}
```

### Solution 2: Top down with DP memoization

We use a `dp` array where every index holds the min cost of each step.

```go
// Top Down Memoization - O(n) 1ms
int[] dp;
public int minCostClimbingStairs(int[] cost) {
	int n = cost.length;
	dp = new int[n];
	return Math.min(minCost(cost, n-1), minCost(cost, n-2));
}
private int minCost(int[] cost, int n) {
	if (n < 0) return 0;
	if (n==0 || n==1) return cost[n];
	if (dp[n] != 0) return dp[n];
	dp[n] = cost[n] + Math.min(minCost(cost, n-1), minCost(cost, n-2));
	return dp[n];
}
```

### Solution 3: Bottom-up DP recursive approach

Since we have our bases clearly defined and they are the minimum values `n` can have, we can ditch the recursive approach altogether by building the solution in a bottom-up manner using iteration.

```go
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i < 2 {
			dp[i] = cost[i]
			continue
		}

		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}

	return min(dp[n-1], dp[n-2])
}
```

### Solution 4: Bottom-up O(1) space

We don’t really need to keep stored the minimum cost of *every* step permanently. Let’s recall the recurrence relation:

```go
mincost(i) = cost[i]+min(mincost(i-1), mincost(i-2))
```

We only need the min cost of the two inmediately previous steps, so after we figure out those two costs we can discard all the rest(from `i - 3` and bellow). So what we do here is using two variables to keep track of the two previous min costs.

```go
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	first, second := cost[0], cost[1]
	var cur int
	for i := 2; i < n; i++ {
		cur = cost[i] + min(first, second)
		first = second
		second = cur
	}

	return min(first, second)
}
```