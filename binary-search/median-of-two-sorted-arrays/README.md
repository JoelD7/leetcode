# Problem
https://leetcode.com/problems/median-of-two-sorted-arrays/description/

Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be $O(\log (m+n))$.



### Example 1:

    Input: nums1 = [1,3], nums2 = [2]
    Output: 2.00000
    Explanation: merged array = [1,2,3] and median is 2.

### Example 2:
    
    Input: nums1 = [1,2], nums2 = [3,4]
    Output: 2.50000
    Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.



### Constraints:

    nums1.length == m
    nums2.length == n
    0 <= m <= 1000
    0 <= n <= 1000
    1 <= m + n <= 2000
    -106 <= nums1[i], nums2[i] <= 106
# Solution
### Rationale

A brute force approach would merge the two arrays and the find the median using regular array indexing, but this leads to linear time complexity $O(m+n)$ and the problem is asking for logarithmic time $O(\log (m+n))$. This constraint gives us a hint into how to solve the problem: binary search.

We have to find a middle element(or cut) that perfectly splits both arrays in two, so that every element before the cut is smaller and every element after is larger. Imagine combining both arrays into one sorted array. The median naturally splits this combined array into a **Left Half** and a **Right Half**.

- **Total Even:** Left and Right halves are exactly equal in size.
- **Total Odd:** Left half gets one extra element (the median itself).

The goal of this code is to find a cut in `nums1` and a cut in `nums2` so that:

1. **Size Rule:** The number of elements on the left side of *both* cuts equals our target Left Half size. Remember, the *median* splits both merged arrays in two sorted portions, so maintaining this rule allows us to find a median that is actually in the middle.
2. **Value Rule:** Every number on the left side is ≤ every number on the right side.

   Visually, it looks like this:

    ```python
    nums1:  ... l1  |  r1 ...
    nums2:  ... l2  |  r2 ...
    ```

   where:

    - `l1` and `l2` are the elements just before the middle of `nums1` and `nums2`, respectively
    - `r1` and `r2` are the elements right after the middle of `nums1` and `nums2`, respectively

   Since the arrays are sorted, we already now that `l1 < r1` and `l2 < r2`, so we need to find a value or cut, for which  `l1 <= r2` and `l2 <= r1`. In other words, the left elements of both arrays are smaller than the right elements of both arrays.


### Algorithm

1. We ensure to always do binary search on the smaller array to prevent out-of-bound errors. So we do:

    ```cpp
    if(n1> n2)return findMedianSortedArrays(nums2, nums1);
    ```

2. Setting the target size `left`
    1. This calculates exactly how many elements need to be in the Left Half. Adding `+ 1` gracefully handles both even and odd total lengths. (e.g., if total length is 5, `left` becomes 3. If total is 6, `left` becomes 3).
3. Set limits of binary search `low` and `high`.
    1. We’ll do binary search on `nums1`, so the limit pointers are set on that array
4. Inside the binary search loop
    1. Set `mid1` and `mid2`, which are the middle pointers of both arrays. `mid2 = left - mid1` because we have to maintain the **size rule**: “The number of elements on the left side of *both* cuts equals our target Left Half size”. Since we’re doing binary search on `nums1`, this has priority so we first set it’s middle index. We set `nums2` middle index with “what remains” after setting `mid1`. If our target Left Half needs 5 elements in total, and we made a cut at `mid1 = 2` (taking 2 elements from `nums1`), we are mathematically forced to make a cut at `mid2 = 3` in the second array (taking 3 elements from `nums2`).
    2. Determine the boundary values of `l1`,`l2`,`r1` and `r2`.
        1. These values are always right before or after the cuts. They always are because we need them to figure out the value rule explained above.
    3. When `l1 <= r2 && l2 <= r1` we found a perfect cut! So:
        1. If total length is odd, the median is just the biggest number on the left.
        2. If total length is even, it's the average of the biggest number on the left and the smallest number on the right.
    4. Adjusting the cut
        1. If `l1 > r2`, it means we included too many large numbers from `nums1` in the Left Half. We need to shrink `nums1`'s left side, so we move our search space to the left (`high = mid1 - 1`). Otherwise, we move it to the right.