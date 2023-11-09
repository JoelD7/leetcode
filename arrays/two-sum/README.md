# Problem
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]



Constraints:

    2 <= nums.length <= 104
    -109 <= nums[i] <= 109
    -109 <= target <= 109
    Only one valid answer exists.

# Solution
Iterate over `nums` subtracting `target` from each value(`val`) of `nums`, and keeping the difference
along with its index in a map, in the form `map[dif] = index`. This difference may be one values that added
with another may result in `target`. On each iteration we ask whether `val` was already obtained from a 
subtraction in a previous iteration by querying the `map`, if it was, this means that there is another value somewhere in `nums`
that added with `val` results in `target`. Because we stored that other value's index in the map, we just need
to return that index and the index of `val`.
