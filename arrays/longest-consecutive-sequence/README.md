# Problem
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.



Example 1:

    Input: nums = [100,4,200,1,3,2]
    Output: 4
    Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9



Constraints:

    0 <= nums.length <= 105
    -109 <= nums[i] <= 109

# Solution
## Variables
- `count`: counts the times an item is consecutive to another
- `cur`: current item in the array being processed
- `prev`: previous item in the array being processed
- `max`: length of the current consecutive array
---
1. Sort the array
2. Use a counter(`count`) to count the times an item in the array(`cur`) is consecutive to the 
previous item(`prev`). When that condition fails: 
   1. reset `count` to 1 no to 0. Why? `count` increments only when comparing **two** items.
   If the comparison succeeds, then we have 2 consecutive items. Resetting`count`to zero would
   mean that after that comparison, `count` would have a value of 1.
   2. update `max` each time `count` is greater
3. We have to check `count` > `max` again after the loop to cover the scenario where the 
longest consecutive sequence reaches up until the last element in the input array.