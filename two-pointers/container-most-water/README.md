# Problem
https://leetcode.com/problems/container-with-most-water/

You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.



Example 1:
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.



Example 1:

![img.png](img.png)

    Input: height = [1,8,6,2,5,4,8,3,7]
    Output: 49
    Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:

    Input: height = [1,1]
    Output: 1



Constraints:

    n == height.length
    2 <= n <= 105
    0 <= height[i] <= 104



Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:

Input: height = [1,1]
Output: 1



Constraints:

    n == height.length
    2 <= n <= 105
    0 <= height[i] <= 104

# Solution
## Clarifications
- The problem is asking to find the largest rectangle area that can be formed using the height
of the lines: `w x h`
- `w = the numbers between the points + 1`
- `h = height of the smaller of the two points`

---
Use two pointers: `l` and `r`. Find the max area by calculating it according to the aforementioned
formula. That calculation requires `w` and `h`. The width is the distance between the pointers,
the height is the smallest value of the two pointers. This is why the height should always point to
the smallest value; we should move the correspoding pointer to match that condition. 