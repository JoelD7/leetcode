# Problem
https://leetcode.com/problems/contains-duplicate-ii/description

Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.


### Example 1:

    Input: nums = [1,2,3,1], k = 3
    Output: true

### Example 2:

    Input: nums = [1,0,1,1], k = 1
    Output: true

### Example 3:

    Input: nums = [1,2,3,1,2,3], k = 2
    Output: false


# Constraints:

    1 <= nums.length <= 105
    -109 <= nums[i] <= 109
    0 <= k <= 105

# Solution
This is basically a sliding window problem for which the "window" is of size `k`.


While iterating through `nums`, use a set(`seen`) to keep track of an item's last position. If we have already `seen` `nums[i]`, we check if it's last position is within the window of size `k`.

## Caveats
- An item may appear several times in the input but we only care if it appears within the `k` window. Why? The `k` window moves as we move through `nums`, so if we are currently looking at the 6th item and `k` = `2`, we don't care about the items before the 4th index. It's because of this reason we keep track of the last time(last position) we saw a specific item in the iteration. 

