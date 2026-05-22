# Problem
https://leetcode.com/problems/kth-largest-element-in-a-stream/

You are part of a university admissions office and need to keep track of the kth highest test score from applicants in real-time. This helps to determine cut-off marks for interviews and admissions dynamically as new applicants submit their scores.

You are tasked to implement a class which, for a given integer k, maintains a stream of test scores and continuously returns the kth highest test score after a new score has been submitted. More specifically, we are looking for the kth highest score in the sorted list of all scores.

Implement the KthLargest class:

    KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of test scores nums.
    int add(int val) Adds a new test score val to the stream and returns the element representing the kth largest element in the pool of test scores so far.


### Example 1:

    Input:
    ["KthLargest", "add", "add", "add", "add", "add"]
    [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
    
    Output: [null, 4, 5, 5, 8, 8]
    
    Explanation:
    
    KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
    kthLargest.add(3); // return 4
    kthLargest.add(5); // return 5
    kthLargest.add(10); // return 5
    kthLargest.add(9); // return 8
    kthLargest.add(4); // return 8

### Example 2:

    Input:
    ["KthLargest", "add", "add", "add", "add"]
    [[4, [7, 7, 7, 7, 8, 3]], [2], [10], [9], [9]]
    
    Output: [null, 7, 7, 7, 8]
    
    Explanation:
    KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);
    kthLargest.add(2); // return 7
    kthLargest.add(10); // return 7
    kthLargest.add(9); // return 7
    kthLargest.add(9); // return 8



### Constraints:

    0 <= nums.length <= 104
    1 <= k <= nums.length + 1
    -104 <= nums[i] <= 104
    -104 <= val <= 104
    At most 104 calls will be made to add.

# Solution
### TL;DR

Using a min-heap of size `k`, means the kth largest element will always be at the root. If we maintain at most `k` elements in that heap, the minimum number will always be the kth largest in the entire sequence.

### Implementation

1. Add elements to the min heap as long as it hasn’t reached a size of `k` OR the top of the heap(the smallest element) is less than the new element.
2. If the new element is larger than the top of the heap that means we now have to remove some element out of the min-heap for two reasons. First, we need to keep the min-heap at a size `k` so that our root is always the kth largest item. Second, since the root is the smallest item and we’re looking for the `kth` *largest*, we obviously need to remove the smaller item which is the root. By removing the root we automatically promote the next largest item after root to be now the kth largest.
    1. We can picture this behaviour like a queue: `5 -> 3 -> 1`. For k = 3, the kth largest item is 1. If we introduce a number larger than 1 to the heap(let’s say, 2) 1 can no longer be the kth largest item because “2” has just displaced him `5 -> 3 -> 2`. So that is why we need to remove the min-heap’s root.
3. Return the min heap’s root. This will be the new root after removing the previous one.