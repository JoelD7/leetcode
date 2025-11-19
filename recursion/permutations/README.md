# Problem
Given an array `nums` of distinct integers, return all the possible _permutations_(rearrangement of all the elements of an array). You can return the answer in any order.


### Example 1:

    Input: nums = [1,2,3]
    Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

### Example 2:

    Input: nums = [0,1]
    Output: [[0,1],[1,0]]

### Example 3:

    Input: nums = [1]
    Output: [[1]]


### Constraints:

    1 <= nums.length <= 6
    -10 <= nums[i] <= 10
    All the integers of nums are unique.

# Solution
## Intuition
The problem is asking for permutations, or in other words, is asking us to iterate over `nums` and _backtrack_ every time we find a valid permutation. This is self-evident when we look at the examples. From them we can tell that in each of the permutations the order of the elements is different, denoting that after reaching a particular element we need to "go back" to previous elements to form other permutations. 

Hence, the problem can be solved using **backtracking**.

## Variables
- `res`: the list of combinations to be returned.
- `comb`: the current combination of elements being formed
- `visited`: map of the elements already added to `comb`. We use this to avoid adding the same index element in `comb` more than once.
  - The **index** distinction here is very important because `nums` might have a number repeated several times but on different indeces. In that case is OK that, let's say `1` is added to `comb` more than once, as long as it is added from different indeces. For that reason we use `i` and not `nums[i]` as the key of `visited`.

## Algorithm
Create a `backtrack` function that iterates over all the elements of `nums` and only considers those who haven't been included in `comb` using the `visited` map.

1. **Base case**: When the length of `comb` is equal to the length of `nums`, we have a valid permutation. Add a copy of `comb` to `res` and return.
2. Iterate over `nums`:
   1. If the current element is already in `comb`, skip it.
   2. Otherwise, add the current element to `comb` and mark it as visited.
   3. Recursively call `backtrack` with the updated `comb` and `visited`. Inside this recursive call, another element different from the current one will be added to `comb`. This is thanks to the `visited` map.
   4. Remove the last element from `comb` and mark it as not visited. We do this so that is possible to form other combinations with that element.