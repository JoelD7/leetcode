# Problem
https://leetcode.com/problems/maximum-product-subarray/description/

Given an integer array `nums`, find a subarray that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.


### Example 1:

    Input: nums = [2,3,-2,4]
    Output: 6
    Explanation: [2,3] has the largest product 6.

### Example 2:

    Input: nums = [-2,0,-1]
    Output: 0
    Explanation: The result cannot be 2, because [-2,-1] is not a subarray.



### Constraints:

    1 <= nums.length <= 2 * 104
    -10 <= nums[i] <= 10
    The product of any subarray of nums is guaranteed to fit in a 32-bit integer.

# Solution
Find the maximum subarray product by making two passes over the array. In each pass, calculate the prefix product and track the maximum value seen so far—first from **left-to-right**, and then from **right-to-left**.

## Implementation

### Variables

- `max`: The maximum product found after traversing the array in both directions. This is the result of the function.
- `curDot`: Holds the dot product of all the previous elements in the current iteration.

To find the result, follow these steps:

1. Calculate the prefix product from left to right, keeping track of the maximum product found.
2. Reset the prefix product and repeat the process from right to left, updating the maximum if a larger product is found. 
3. Return the overall maximum product recorded from both passes.

### Caveats

- We should not consider 0s it in the multiplication as it will nullify the whole prefix product array. Instead, we first ask whether 0 > `max`(it might be a negative value), and update the `max` accordingly. Then reset `curDot` to 1 in order to start building the prefix product array again from the next element. 
  - **Why should we ignore all the previous elements?** The subarray ends in zero. There is no point in considering that subarray because multiplying all it's elements will result in zero. If `curDot` was negative, then 0 should be `max`; if `curDot` was positive, we update `max` if needed and continue searching for another subarray that might produce a larger value. 
- Kadane's algorithm doesn't apply here because it relies on the fact that if the current result plus the current element is less than the current element, then current element is the larger subarray so far and there is no point in considering the previous elements. This is not true for the "max product subarray" because of negative numbers. Let's say that `curDot` is a negative number and `nums[i]` isn't. The product of these values will result in a negative number that is obviously smaller than `nums[i]`. However, there is the chance that after `nums[i]` there is another negative value that when multiplied with `curDot` and `nums[i]` produces a larger number than `nums[i]` alone. It's precisely because of this reason that we can't restart the array building process when `curDot` * `nums[i]` < `nums[i]`. Try to apply Kadane's algorithm to the following array so you can see what I mean: `[2, -5, -2, -4, 3]`.