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
The array is rotated, so the most straightforward way to solve this problem is finding the rotation point, so we can then perform a regular binary search.

The rotation point is obviously the index that points to the smaller element, therefore, we use binary search to find it. At the end of the following loop, `left` will point to smallest element. For details on why this works, look → [Finding a minimum value `k` such that a condition for `k` is true or…](https://www.notion.so/Binary-Search-2e9fdddd7cce80898bddcf7c34a6da64?source=copy_link#2edfdddd7cce8039b609f534e6eee1b8)

```go
	for left < right {
		mid = left + (right-left)/2

		//This means the array is rotated so we look for the smallest element on this range
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
```

---

When then do a regular binary search but first, set our left and right pointers based on the relation that `target` has with our rotation point.

```go
	rot := left
	left, mid, right = 0, 0, len(nums)-1

	//This means that target is between rot and right
	if target >= nums[rot] && target <= nums[right] {
		left = rot
	} else {
		right = rot
	}
```

Finally, do binary search

```go
for left <= right {
		mid = left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
```