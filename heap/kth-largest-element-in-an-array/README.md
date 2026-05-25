# Problem
https://leetcode.com/problems/kth-largest-element-in-an-array/description/

Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?


### Example 1:
    
    Input: nums = [3,2,1,5,6,4], k = 2
    Output: 5

### Example 2:

    Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
    Output: 4


### Constraints:

    1 <= k <= nums.length <= 105
    -104 <= nums[i] <= 104

# Solution
Keep a min heap of a maximum size `k`. Only add elements to it when either the size is lower than `k` or if the new item to be added is larger than the heap’s root.

The root of a min heap is always the smallest element of the entire heap. By exclusively adding elements larger than the current root and not letting the heap exceed a size of `k`, the heap will always have the `k` largest items, with the $k^{th}$ largest item being smallest of the heap(root).

After processing the entire input sequence return the min heap’s root.