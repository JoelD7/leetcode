# Problem
https://leetcode.com/problems/house-robber/description/

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.



### Example 1:

    Input: nums = [1,2,3,1]
    Output: 4
    Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
    Total amount you can rob = 1 + 3 = 4.

### Example 2:

    Input: nums = [2,7,9,3,1]
    Output: 12
    Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
    Total amount you can rob = 2 + 9 + 1 = 12.


### Constraints:

    1 <= nums.length <= 100
    0 <= nums[i] <= 400

# Solution
### Rationale

The reucrrence relation goes like this:

```go
rob(i) = Math.max( rob(i - 2) + currentHouseValue, rob(i - 1) )
```

In other words, the maximum amount of money you can get by robbing `i` houses, can either be the loot gotten from the current house plus the ones before the previous house, or the loot from the previous house and any loot capture before that.

### Implementation 1 - recursive + top down

This recurrence relation can be translated in a very straightforward manner to code, using memoization:

```go
int[] memo;
public int rob(int[] nums) {
    memo = new int[nums.length + 1];
    Arrays.fill(memo, -1);
    return rob(nums, nums.length - 1);
}

private int rob(int[] nums, int i) {
    if (i < 0) {
        return 0;
    }
    if (memo[i] >= 0) {
        return memo[i];
    }
    int result = Math.max(rob(nums, i - 2) + nums[i], rob(nums, i - 1));
    memo[i] = result;
    return result;
}
```

### Implementation 2 - iterative + bottom up

If we build the memo array from the “bottom”(`i = 0`) we can note that at every decision step we only need to choose between two values, so memoizing the all possible inputs(`i`) is not necessary:

```go
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var prev1, prev2 int
	for _, num := range nums {
		tmp := prev1
		prev1 = max(prev2+num, prev1)
		prev2 = tmp
	}

	return prev1
}
```