# Problem
https://leetcode.com/problems/house-robber-ii/description/

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.



### Example 1:

    Input: nums = [2,3,2]
    Output: 3
    Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.

### Example 2:

    Input: nums = [1,2,3,1]
    Output: 4
    Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
    Total amount you can rob = 1 + 3 = 4.

### Example 3:

    Input: nums = [1,2,3]
    Output: 3



### Constraints:

    1 <= nums.length <= 100
    0 <= nums[i] <= 1000

# Solution
This problem is an iteration of https://leetcode.com/problems/house-robber/description/, with the distinction that the houses at the “start” and the “end” are adjacent. That is all there is to it. House Robber II can be solved just like House Robber I but bearing that distinction in mind.

We can see that this is an optimization problem because we are trying to maximize something. Therefore we can use dynamic programming to solve it. The problem can be broken down in the following recursive manner:

```go
rob(i) = Math.max( rob(i - 2) + currentHouseValue, rob(i - 1) )
```

In other words, at every house we get to we must decide whether to rob it and take all the loot gotten up to the house before the previous one, or ignore the current house and take all the loot gotten up to the previous house.

Up to this point the solution is just like House Robber I, what makes this problem different is that since the first and last houses are adjacent we need to separate the maxes in two:

For `n = length of nums`...

1. First we solve the problem including the first house and not the last one, from `0...n-2`
2. Second we solve the problem including the last house and not the first one, from `1...n-1`
3. The output of will be the max between the two values