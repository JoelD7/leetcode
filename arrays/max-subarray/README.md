# Problem
https://leetcode.com/problems/maximum-subarray/description/

Given an integer array nums, find the subarray with the largest sum, and return its sum.



**Example 1:**

    Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
    Output: 6
    Explanation: The subarray [4,-1,2,1] has the largest sum 6.

**Example 2:**

    Input: nums = [1]
    Output: 1
    Explanation: The subarray [1] has the largest sum 1.

**Example 3:**

    Input: nums = [5,4,-1,7,8]
    Output: 23
    Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.



**Constraints:**

    1 <= nums.length <= 105
    -104 <= nums[i] <= 104



Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

# Solution
This is solved Kadane's algorithm that is precisely designed for finding a subarray that produces the maximum sum within an array.

## Implementation
First, we initialize our starting values:

    func maxSubArray(nums []int) int {
        maxSum := -10000 
        curMaxSum := 0
        max := func(a, b int) int {
            if a > b {
                return a
            }
            return b
        }
    ...
    }

- `maxSum`: This will hold the maximum sum found so far. We'll return this value at the end of the function.
- `curMaxSum`: Holds the maximum sum of all the subarrays that end in `i`, or said in aonther way, of the subarray we're currently iterating over.
- `max`: Utility function that returns the maximum value between two integers.

Then, we iterate over the array and compare each element `nums[i]` with `curMaxSum` plus `nums[i]`. If adding `nums[i]` to `curMaxSum` produces a larger value than `nums[i]`, we add that value to `curMaxSum`, update `maxSum` and continue with the next iteration to see if adding the next element produces a larger value(we're looking for the maximum after all):

    for i := 0; i < len(nums); i++ {
		if nums[i]+curMaxSum >= nums[i] {
			maxSum = max(maxSum, nums[i]+curMaxSum)
			curMaxSum += nums[i]
			continue
		}

If it doesn't, then it means that the subarray upto `nums[i]` included negative values, and that `nums[i]` is a subarray larger than all the subarrays ending in it, so it doesn't make sense to keep track of the previous values. Hence, we reset `curMaxSum` to `nums[i]` and continue iterating. 

    for i := 0; i < len(nums); i++ {
        ...
        curMaxSum = nums[i]
        maxSum = max(maxSum, nums[i])
    }

When we are done iterating, we return `maxSum`.