# Problem
https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/

Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

    [4,5,6,7,0,1,2] if it was rotated 4 times.
    [0,1,2,4,5,6,7] if it was rotated 7 times.

Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.


### Example 1:

    Input: nums = [3,4,5,1,2]
    Output: 1
    Explanation: The original array was [1,2,3,4,5] rotated 3 times.

### Example 2:

    Input: nums = [4,5,6,7,0,1,2]
    Output: 0
    Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

### Example 3:

    Input: nums = [11,13,15,17]
    Output: 11
    Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

### Constraints:

    n == nums.length
    1 <= n <= 5000
    -5000 <= nums[i] <= 5000
    All the integers of nums are unique.
    nums is sorted and rotated between 1 and n times.

# Solution
1. Use binary search to find the minimum element that satifies a condition
2. The condition is: “is `mid` higher than `right`?”
    1. If it is → `left = mid + 1` because we now know that the array is rotated after `mid`, so the smaller elements are all after this point, so we continue searching after `mid`.
    2. If it isn’t → `right = mid` because this means all the smaller elements are before `mid`, so we must continue searching in that portion of the array
3. We do this until `left` and `right` meet.
4. Return the value that `left` points to → `nums[left]`

---

By continuously asking whether `mid` is higher than `right` and moving one pointer or the other accordingly we end up with `left` pointing to the smallest element.