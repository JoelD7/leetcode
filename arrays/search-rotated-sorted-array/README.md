# Problem
https://leetcode.com/problems/search-in-rotated-sorted-array/description/

There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index `k (1 <= k < nums.length)` such that the resulting array is `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` (0-indexed). For example, `[0,1,2,4,5,6,7]` might be rotated at pivot index 3 and become `[4,5,6,7,0,1,2]`.

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with `O(log n)` runtime complexity.


### Example 1:

    Input: nums = [4,5,6,7,0,1,2], target = 0
    Output: 4

### Example 2:

    Input: nums = [4,5,6,7,0,1,2], target = 3
    Output: -1

### Example 3:

    Input: nums = [1], target = 0
    Output: -1


### Constraints:

    1 <= nums.length <= 5000
    -104 <= nums[i] <= 104
    All values of nums are unique.
    nums is an ascending array that is possibly rotated.
    -104 <= target <= 104

# Solution
We use a variation of binary search with a twist. First let's remind ourselves how binary search works. We have three indices: `low`, `mid`, and `high`. `low` points to the lowest value of the current subarray,`high` to the highest and `mid` to the middle. 
Depending on weather `mid` is higher or lower than our `target`, we change the values of `low` or `high`, modifying the searchable subarray accordingly. 

Even though the array(`nums`) is sorted, since it's also rotated there may be some situations in which the `low` index points to a value **larger** than `mid`. This specifically occurs when there is a pivot. We use this caveat to know if we really should move the `low` index or the `high` index. 

## Implementation
If both `low` and `mid` are larger than `target` **and also** `low` is larger than `mid`...

**Example**: `nums` = `[..,6,7,0,1,2,3,...]` and `target` = `0`. Assuming there are other elements in the array and in one of the iterations we end up with `mid` = `2` and `low` = `6`

Here we know for a fact that we are **partially inside** a pivot range, which means that after `low` there are numbers lower than `mid`, and since we just checked that `low` and `mid` are larger than `target`, we can be certain that target is in that range. For this reason we keep looking in this range by reducing the `high` pointer to a place before `mid`..

```go

    } else if nums[mid] > target && nums[low] > target && nums[low] > nums[mid] {
            high = mid - 1
```


---
What if we have the same condition but `low` isn't larger than `mid`? 

**Example**: `nums` = `[4,5,6,7,0,...]` and `target` = `0`. Assuming there are other elements in the array and in one of the iterations we end up with `mid` = `6` and `low` = `4`

On this circumstance we are **inside** the pivot range completely. So even though this portion of the array is sorted, since we are inside the pivot then after the `mid` point there are elements that are lower than both `low` and `mid`. Hence, we look for `target` in that range, ignoring all the elements before `mid`...

```go
     else if nums[mid] > target && nums[low] > target {
                low = mid + 1
            }
```

---
If we aren't inside or partially inside a pivot range, then we use the regular logic for binary search:
```go
 else if nums[mid] > target {
			high = mid - 1
		}
```
---
The same applies for the rest of the conditionals