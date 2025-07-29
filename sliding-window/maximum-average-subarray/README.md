# Problem
https://leetcode.com/problems/maximum-average-subarray-i/description

You are given an integer array `nums` consisting of `n` elements, and an integer `k`.

Find a contiguous subarray whose length is equal to `k` that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.


## Example 1:

    Input: nums = [1,12,-5,-6,50,3], k = 4
    Output: 12.75000
    Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

## Example 2:

    Input: nums = [5], k = 1
    Output: 5.00000



## Constraints:

    n == nums.length
    1 <= k <= n <= 105
    -104 <= nums[i] <= 104

# Solution
Maintain a window of size `k`. On the first iteration add up the first `k` elements, save the total and get the average. Subsequent iterations only require to subtract the first element of the previous window from the `total`, and adding up the last element of the current one. 

For illustration, take the first example -> `nums = [1,12,-5,-6,50,3], k = 4`

**Iteration 1**

    1 + 12 - 5 - 6 = 2(curTotal)

**Iteration 2**

Instead of looping from index `1` up to index `1 + k - 1`, do:
    
    curAvg - prevFirst + lastElement
    2 - 1 + 50 = 51(newTotal)


## Main variables

- `i`: The start of the window
- `i+k-1`: The end of the window 

## Intuition
The key to solve this problem in `O(n)` time, is that you only need to traverse the entire window during the first iteration; subsequent iterations require only arithmetic.


## Implementation
The code is a pretty straightforward implementation of the solution I explained before. Of course, updating the `maxAvg` when we find a larger one.

The loop goes upto `len(nums)-k` because that is the maximum value the start of the window can have.  

----
We calculate the average by iterating over all the elements of the window **only the first time**:
```go
...
for i == 0 && (j-i) < k {
    total += float64(nums[j])
    j++
}

if i == 0 {
    maxAvg = max(maxAvg, total/float64(k))
    continue
}
...
```

## Edge cases

- `len(nums) == k`: In this case we iterate over the whole array, add up all it's elements and return the average.  
