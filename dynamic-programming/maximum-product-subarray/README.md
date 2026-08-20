# Problem
https://leetcode.com/problems/maximum-product-subarray/description/

Given an integer array nums, find a that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.

Note that the product of an array with a single element is the value of that element.


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
### Rationale

Do a prefix and suffix product of `nums` and return the max.

Every index `i` of a prefix product array will have the product of all the elements of `nums` from `0` up to `i`. Due to negative numbers, a higher index in the prefixProduct array may have a smaller value, for this reason we update max on every iteration of `nums`.

**What happens if the max product subarray doesn’t include the 0th index?** This is why we also calculate the suffix product array, starting at the last index and calculating accumulating products backwards.

The max product subarray will *always* include either the first and/or last indeces. There is no scenario where the max product subarray is only composed of middle elements without including the extremes. Mathematically this makes sense because we are *maximizing*, so in general, the more numbers we add to our product, the larger the product will be which is what we want.

Then, we return the max of either arrays.