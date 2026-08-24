# Problem
https://leetcode.com/problems/longest-increasing-subsequence/description/

Given an integer array nums, return the length of the longest strictly increasing .



### Example 1:

    Input: nums = [10,9,2,5,3,7,101,18]
    Output: 4
    Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

### Example 2:
    
    Input: nums = [0,1,0,3,2,3]
    Output: 4

### Example 3:

    Input: nums = [7,7,7,7,7,7,7]
    Output: 1



### Constraints:

    1 <= nums.length <= 2500
    -104 <= nums[i] <= 104

# Solution
### Rationale

The problem can be solved by decomposing it to smaller subproblems: this is a classic dynamic programming approach. Instead of trying to find the longest increasing subsequence(LIS from now on) for the entire `nums` array at once, we find the LIS for the subsequence that *ends* at a specific element.

So let `dp[i]` be the LIS for the subsequence that ends at `nums[i]`. Each value of `dp[i]` will be populated with a bottom-up approach, looking back at all the elements before `nums[i]` checking whether they’re smaller and updating accordingly.

The initialized value of every `dp[i]` element is 1, as the LIS for a single element is 1.

### Algorithm

To fill out `dp[i]` for any given number `nums[i]`, we look back at all the numbers that came *before* it (let's call each previous number `nums[j]`).

For every previous number `nums[j]` where `j < i`:

- **Is it strictly increasing?** We check if `nums[i] > nums[j]`.
    - **If yes:** It means we can attach `nums[i]` to the end of the longest subsequence that ended at `nums[j]`.
    - We then update `dp[i]` to be the maximum between its current value and `dp[j] + 1`. Remember that `dp[j]` represents the LIS of a previously calculated subsequence that ended at `nums[j]`. The reason we add 1 here is because we’re including `nums[i]` to this sequence. The reason we choose the max between the current value of `dp[i]` and `dp[j] + 1` is because there may be several LIS before `nums[i]`, and since we want the longest one, we always choose the longest.