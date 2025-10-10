# Problem
https://leetcode.com/problems/minimum-size-subarray-sum/

Given an array of positive integers `nums` and a positive integer `target`, return the minimal length of a subarray whose `sum` is greater than or equal to `target`. If there is no such subarray, return `0` instead.


### Example 1:

    Input: target = 7, nums = [2,3,1,2,4,3]
    Output: 2
    Explanation: The subarray [4,3] has the minimal length under the problem constraint.

### Example 2:

    Input: target = 4, nums = [1,4,4]
    Output: 1

### Example 3:

    Input: target = 11, nums = [1,1,1,1,1,1,1,1]
    Output: 0


### Constraints:

    1 <= target <= 10^9
    1 <= nums.length <= 10^5
    1 <= nums[i] <= 10^4

# Solution
## Intuition
A naive approach would involve finding the sum of all the subarrays that can be formed with each element of `nums`, but this would result in `O(n^2)` time complexity and a `time limit exceeded` error from Leetcode.

We're being asked to find the smallest subarray that satisfies a condition. This screams **sliding window**. 

## Main actors 
- `i`: left pointer of the window
- `j`: right pointer of the window
- `sum`: contains the sum of the current subarray window
- `minLength`: the length of the smallest window(subarray) found so far. This is the value the function will return. It must be initialized with a value larger than the maximum length the `nums` array can have according to problem description to produce accurate results. 

## Algorithm
We create a loop that on each iteration will add the `j`th element to `sum` and ask whether `sum >= target`. If and when the condition is satisfied we must do two things: 
1. Update `minLength` with the length of the current window, if possible. The current window might have a length larger than `minLength`, so we won't update this variable every time. 
2. Reduce the window by increasing the left pointer `i` and subtract the `i`th element from `sum`, due to the possibility of forming an even smaller subarray that satisfies the condition. Note that there is a chance that the smaller subarray can only be formed after increasing `i` several times. Because of this reason we need to create an inner loop that does this continously while `sum >= target`. 

Before returning we check if `minLength` still has the same value it was initialized with. If that's the case it means we couldn't find any subarray that satisfies the condition so we must return `0`.  





















