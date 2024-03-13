# Problem
https://leetcode.com/problems/binary-tree-inorder-traversal/

Given the root of a binary tree, return the inorder traversal of its nodes' values.



Example 1:

![img.png](img.png)

    Input: root = [1,null,2,3]
    Output: [1,3,2]

Example 2:

    Input: root = []
    Output: []

Example 3:

    Input: root = [1]
    Output: [1]



Constraints:

    The number of nodes in the tree is in the range [0, 100].
    -100 <= Node.val <= 100


**Follow up**: Recursive solution is trivial, could you do it iteratively?

# Solution
Using recursion, first navigate to all the left children, then add root, then all the right children.