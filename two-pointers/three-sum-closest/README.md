# Problem
https://leetcode.com/problems/3sum-closest/description

Given an integer array `nums` of length `n` and an integer `target`, find three integers in `nums` such that the sum is closest to `target`.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.


### Example 1:

    Input: nums = [-1,2,1,-4], target = 1
    Output: 2
    Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

### Example 2:

    Input: nums = [0,0,0], target = 1
    Output: 0
    Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).


### Constraints:

    3 <= nums.length <= 500
    -1000 <= nums[i] <= 1000
    -104 <= target <= 104


# Solution
Use three pointers on a sorted array and find the lowest difference between the sum of the three pointer's elements and the target. 

## Main variables
- `i`: The "base" pointer. This is the one that it'll move slower. This is our starting point for adding numbers. 
- `j`: The second pointer that will start one position after `i`. 
- `k`: The third and final pointer. It will always start at the end of `nums`
- `threeSum`: Sum of the threepointers.
- `closestTotal`: The closest sum to `target` we have found so far. This is the return value of the function.
- `minDiff`: The minimum difference we have found between `threeSum` and `target`.
- `curDiff`: The difference between `threeSum` and `target` of the **current iteration**.

## Algorithm
The closest value will be the one that has the lower difference(subtraction) from `target`, meaning, a value such that when we substract `target` from it we get a difference as close to 1 as possible. Hence, on each iteration we subtract `target` from `threeSum` and keep track of the lowest difference found so far in `minDiff`. Every time we find a lower difference it means a closest sum ahs been found, so we update the value of `closestTotal` with `threeSum` accordingly. 

By sorting the array, the task of finding a sum closest to `target` gets considerably easier because we know _where_ to look for a closest sum. How? if `threeSum` is larger than `target` we reduce the pointer(`k`) that has the largest values of `nums` in order to get a smaller, closest total to `target`; if the `threeSum` is smaller, we increase the pointer has the smallest values of `nums` in order to get a larger, closest total to `target`.

