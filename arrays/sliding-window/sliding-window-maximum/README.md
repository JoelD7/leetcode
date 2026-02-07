# Problem
https://leetcode.com/problems/sliding-window-maximum/description/

You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.


### Example 1:

    Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
    Output: [3,3,5,5,6,7]
    Explanation:
    Window position                Max
    ---------------               -----
    [1  3  -1] -3  5  3  6  7      3
    1 [3  -1  -3] 5  3  6  7       3
    1  3 [-1  -3  5] 3  6  7       5
    1  3  -1 [-3  5  3] 6  7       5
    1  3  -1  -3 [5  3  6] 7       6
    1  3  -1  -3  5 [3  6  7]      7

### Example 2:

    Input: nums = [1], k = 1
    Output: [1]


### Constraints:

    1 <= nums.length <= 105
    -104 <= nums[i] <= 104
    1 <= k <= nums.length

# Solution
### Rationale

The solution uses a **Decreasing monotonic queue**, that maintains all the elements of the queue in decreasing order. It does this by either ignoring or removing elements from the back of the queue that are larger than the current element being added. A DMQ is often implemented as a doubly-ended queue(deque) because with this queue variation we often need to add and remove from the front and back of the queue.

**Why use this data structure?** Because of two reasons:

1. A DMQ is “*decreasing*”, which means it will always have the largest element at the front of the queue, which can be extracted in O(1) time. This is what we want in this problem: collect all the largest elements of each subsequent window as we slide through the input array.
2. A DMQ ignores/deletes smaller elements than the current one being added. This adapts to our use case as there is no need to consider smaller elements because after all we’re looking for the largest ones.

### Implementation

We’ll iterate through `nums` maintaining the queue in decreasing order and adding to our `results` array, all the largest elements of the windows.

1. Create a queue `q` that will store indeces of the elements. Storing indeces is easier because they’re unique and allow us to easily reference elements in `nums` without needing additional data structures.
2. Remove from the front all the elements that don’t belong to the window.
    1. For example, if `i = 5` and `k = 3` then the current window(`i-k+1`) consist only of indeces 3, 4 and 5, so we need to delete from the queue all the elements before index 3 to avoid polluting our results. Remember, we need to gather the largest elements of **each individual window,** so if we keep the queue with elements of other windows we’ll get invalid results.
3. Remove from the back, elements smaller than the current one being added(`nums[i]`). This is the key quality that keeps and makes the queue a decreasing monotonic queue, which we need to be able to solve this problem as explained in previous paragraphs.
4. After the previous loop is done, `i` will be placed where it should in the queue to maintain decreasing order.
5. As we already established, the front of the queue will be the largest element in the current window, so we add it to `results`.
    1. We do this only after having a **complete window of size `k`**, which is why the condition `if i >= k-1` is in place. Example, if `i = 1` and `k = 3`, then our window will have only two elements(index 0 and index 1) so we can’t assume we have already found the largest element because we haven’t analyzed the complete window yet.