# Problem
https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/

Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.



### Example 1:
![img.png](img.png)

    Input: root = [3,1,4,null,2], k = 1
    Output: 1

### Example 2:
![img_1.png](img_1.png)
    
    Input: root = [5,3,6,2,4,null,null,1], k = 3
    Output: 3



### Constraints:

    The number of nodes in the tree is n.
    1 <= k <= n <= 10^4
    0 <= Node.val <= 10^4

# Solution
**The $k^{th}$ smallest element in a BST is the $k^{th}$ value in the in-order traversal array of the tree**. So we perform in-order traversal, adding each element to the array and returning the value at the moment the array reaches a length of `k`. There is no need to keep going deeper.

### Note

- We use `-1` as sentinel value because according to the problem’s constraints, all nodes in the tree are positive numbers.