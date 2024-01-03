# Problem
https://leetcode.com/problems/daily-temperatures/

Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.



Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Example 2:

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Example 3:

Input: temperatures = [30,60,90]
Output: [1,1,0]



Constraints:

    1 <= temperatures.length <= 105
    30 <= temperatures[i] <= 100

# Solution
On each iteration ask whether the next element(`next`) is bigger than the current one(`cur`). If it is, then the answer is just 1. That is the simplest case. When this doesn't happen, we will use a stack to store the indices of those elements in the hopes of finding another bigger element up a head. 

**Why storing the indices and not the elements?**
1. There may be repeated elements
2. What we want is to solve the question "how far away is X from Y?". This question can only be answered by keeping track of positions, i.e. indexes. 
3. It's much, much easier to get an element from it's position than a position from it's element. 


When `next > cur` we will also inspect the stack to check whether if `next` is also bigger than any other elements that are in the stack. We will only pop the elements that are smaller than `next`. In a further iteration there may be another "`next`" that is greater than the remaining elements in the stack.

After we are done iterating over `temperatures` we should set the answer for all remaining elements in the stack to 0, because we never found any elements bigger than them. 