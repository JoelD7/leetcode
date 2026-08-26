# Problem 
https://leetcode.com/problems/partition-equal-subset-sum/description/

Given an integer array nums, return true if you can partition the array into two subsets such that the sum of the elements in both subsets is equal or false otherwise.


### Example 1:

    Input: nums = [1,5,11,5]
    Output: true
    Explanation: The array can be partitioned as [1, 5, 5] and [11].

### Example 2:

    Input: nums = [1,2,3,5]
    Output: false
    Explanation: The array cannot be partitioned into equal sum subsets.


### Constraints:

    1 <= nums.length <= 200
    1 <= nums[i] <= 100

# Solution
### Rationale

In simpler terms, we’re being asked to see if there are two subsets in `nums` whose sum is equal. The key piece of information here is the total sum(`total` from now on). The two subsets must produce the same total, which is `total/2`. First of all, if `total` is an odd number we must return false inmediately as this implies no two subsets will produce the same result. On the other hand, if we find a subset that adds up to `total/2`(`target`) we know that there is another subset that adds up to the same thing because the existence of the first half implies the existence the second. Therefore, we don’t need to find the second half.

So, **the main job here is finding the numbers that when adding them up we get `total/2`(`target`)**. To do that we’ll use dynamic programming: finding subsets that add up to a number `i`, progressively building up our solution until `i` reaches `target`.

### Algorithm

The algorithm mainly relies on the following idea: `*dp[i]` will be true if we can make a subset that sums up to `i` hence, `dp[0] = true` because can always make a sum of 0 using an empty subset*.

We’ll build an array that keeps track of every possible sum between `target` and `num` we can create using the numbers we have seen so far.

1. Calculate `total`. Return `false` is `total` is an odd number as it is impossible to find two subsets that add up to the same `total` half
2. Calculate our `target` and initialize the `dp` array
    1. We initialize the array to `target + 1` because arrays are zero-indexes and we want to be able to reference `dp[target]` without getting into an out-of-bounds error.
3. Set `dp[0] = true`. This line it’s the base truth that allows us to set every other sum to have an accurate value
4. Iterate through `nums` and for each `num`
    1. Iterate backwards from `target` to `num` to set the value of `dp[i]`
        1. The logic is the following: we can form a sum `i` if we could previously OR if we can form a sum `i - num`. The last part of this condition is the most important one, so let’s use an example to explain it.

           Imagine `i = 12` and `num = 5`. Suppose that on past iterations we figured out that we could form a valid sum subset with 7. The key piece of reasoning here is that if we could form a valid subset sum with 7, we could also do it with 12 **BECAUSE** **there is a number 5 on the `nums` array**, and 5 + 7 = 12.


### How does the “build-up” of `dp` occurs

You may have noticed that during the first iterations `dp` is mostly false which might have produced some confusion. Why do this if it’s “always” false? Well, it’s not “always” false. Let’s say we have this setup:

- **nums Array:** `[5, 5, 6, 7, 11, 13, 23]`
- **Total Sum:** 70
- **Target (`S/2`):** 35

The inner loop goes backwards from the `target` down to `num` (so, from 35 down to 5). For most of that loop, nothing interesting happens. `dp[i]` stays `false`.

But watch what happens at the very end of that inner loop, when **`i = 5`**:

The code evaluates this formula:
`dp[i] = dp[i] || dp[i - num]`

Let's plug in the numbers (`i = 5`, `num = 5`):

1. `dp[5] = dp[5] || dp[5 - 5]`
2. `dp[5] = dp[5] || dp[0]`
3. `dp[5] = false || true` *(Because we manually set dp[0] to true!)*
4. **`dp[5] = true`**

Here we’re saying that it is possible to form valid subset sum that adds up to 5. From this point on, further iterations will populate all values from `target` to the current `num`. When we get to the second 5 and `i = 10`, `dp[10]` will be `true` because `dp[10 - 5] = true` as we established on the first iteration, which means that if we could form a valid subset sum with 5, we can also do it with 10 because there is another number 5 on the nums array and 5 + 5 = 10.