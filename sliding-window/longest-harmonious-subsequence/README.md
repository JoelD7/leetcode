# Problem
https://leetcode.com/problems/longest-harmonious-subsequence/description

We define a harmonious array as an array where the difference between its maximum value and its minimum value is exactly 1.

Given an integer array nums, return the length of its longest harmonious subsequence among all its possible subsequences.


### Example 1:

    Input: nums = [1,3,2,2,5,2,3,7]
    
    Output: 5
    
    Explanation:
    
    The longest harmonious subsequence is [3,2,2,2,3].

### Example 2:

    Input: nums = [1,2,3,4]
    
    Output: 2
    
    Explanation:
    
    The longest harmonious subsequences are [1,2], [2,3], and [3,4], all of which have a length of 2.

### Example 3:
    
    Input: nums = [1,1,1,1]
    
    Output: 0
    
    Explanation:
    
    No harmonic subsequence exists.



Constraints:

    1 <= nums.length <= 2 * 104
    -109 <= nums[i] <= 109

# Solution
## Abstract
Use a sliding window in a sorted array to keep track of harmonious items within the window, and return the size of the largest window found.

## Premises
- We don't need to build a subarray or another array with the elements of the subsequence; we only need to count elements and return the largest `count`. 
- A sorted array is easier to process for this problem because it's easier to compare elements that are directly next to each other, and also saves us the need to perform a lookup of the whole array on each iteration. 
- It's safe to sort the array because: 
  1. we're being asked for a subsequence, not a subarray
  2. the order of the elements in a subsequence doesn't matter to know if it is harmonious, i.e. both `[1, 2]` and `[2,1]` are harmonious.

## Main actors
- `i`: Demarks the **beggining** of the window. 
- `j`: Demarks the **end** of the window. We move this value on each iteration to enlarge the window, only when we know that there is a possibility that the next item can be part of the sequence.
- `dif`: Difference between `nums[j] - nums[i]`.
- `newI`: This is the value `i` will take when we decide to change the start of the window. We'll change `newI` every time we find a value that is larger than `i`. Why? Because we want to find out wether that new value is be the start of a new, larger subsequence or not.
- `count`: A counter that allows us to track the size of the subsequence.
- `largest`: The size of the largest subsequence found so far. We update it every time we find a bigger size. This is the return value of the function. 
- `updateLargest`: Utility function to update `largest`. 
  - We update this value only when `count > 0` because if it's 0, it means that all items are equal(this is the only way`count` remains 0).
  - We add `1` to `count` because by subtracting `i` from `j`, the `jth` element is not included. For example: for `[1,2,2,2], i = 0, j = 3`, `j - i = 3` but the subsequence size should be `4`.


## Algorithm
We move `j` until reaching certain conditions, using `dif` as the main decider.

On this case we just increase `j` because the next item can be part of the window(equal to the current item or larger by 1). We don't increase `count` here because all the items of `nums` might be equal:
```go
if dif == 0 {
    j++
    continue
}
```
----
`(dif == 1 || dif == -1)` means that the current item(`j`) is part of the subsequence we're counting from `i`. We change the value of `count` to be `j - i` because we know for a fact that the values between `i` and `j` are harmonious. How do we know that? The array is sorted.

Since we want `i` to always be at the start of a new subsequence, we update `newI` only when the current `j` is different than the previous one. For example, take an array `[1,2,2,2]`, if `i = 0` and `j = 1`, we want `newI` to be `1`. If we update `newI` on each iteration this value would endup being the last time the number `2` appeared in the array(index 3).

We then increase `j` to keep looking.
```go
if dif == 1 || dif == -1 {
    if nums[j-1] != nums[j] {
        newI = j
    }

    count = j - i
    j++
    continue
}

```
> **Why `count = j - i` instead of `count++`?** Because the first element might be repeated several times before we find a difference of one. Example: `[1,1,1,1,2]`
---
We'll change the start of the window every time the difference is larger than 1. This means that `j` and `i` cannot possibly be harmonious so there is no point in keeping enlarging the window.

Due to the matter that we only change `newI` when the difference is `1`, the fact that `newI == 0` means that `i` was the first item of `nums` and `j` the second, so to reset the window we move both pointers one position. We update `newI` to later look for a new window starting at `j`. 
```go
if (dif > 1 || dif < -1) && newI == 0 {
    newI = j
    i++
    j++
    continue
}
```
---
In this case we know that we are not comparing the 1st and 2nd element of `nums` because we already checked that in a previous conditional, so this means that we have a subsequence with more than one item. For that reason we update `largest`, set `i` to `newI`, `j` directly next to this new `i`and reset `count`.
```go
if dif > 1 || dif < -1 {
    largest = updateLargest(largest, count)
    count = 0
    i = newI
    newI = j
    j = i + 1
continue
}
```
---
Inside the loop, `largest` is updated only when `dif` exceeds `1`. This doesn't always happen, i.e., when the whole array is the LHS. Example: `nums = [1,2,1,2]`. For that reason we update `largest` outside the loop as well, just in case.
```go
largest = updateLargest(largest, count)
```



















