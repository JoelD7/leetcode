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
The first interesting thing to note here is that for this problem there is no need to implemented a linked list stack, which involves considerably more code. An array is enough to do what we want here.

## Variables

- `answer`: output of the program. We initialize it with zeroes on all positions to handle the default values indicated by the problem.
- `stack`: the stack that holds the indices of temperatures. We deal with indices because temperatures can be repeated, and also what we care about in the context of the problem are *when* temperatures happen, in other words, positions in the array.
- `temp`: current temperature being evaluated in the iteration.

## Algorithm

In a nutshell… Iterate over all the temperatures and continously add each index to a stack. On each temperature iteration, if the top of the stack is smaller than the current temperature, then we have found an answer for the top of the stack. Therefore, update the answers array of the index at the top of stack with the amount of days that passed between the top of stack and the current temperature. Pop the answered element out of stack and do the process again just in case the new top is also smaller than the current temperature.

1. Iterate over all the temperatures and push every index to the stack.
2. If the top of the stack references a temperature lower than the current one…
    1. we calculate the difference between the index of the current temperature and the index of that stack element and set it in `answer`. The difference is how many days we “waited” before finding a larger temperature than the top of the stack. Why? Remember…
        1. `index` = index at the top of the stack
        2. `i` = index of the current temperature
    2. we pop the top of the stack because we have found the answer for that element so we don’t care about it anymore.
    3. we repeat this process **while** the stack has elements. We do this to process “pending” elements. In some iterations the top of stack won’t always be smaller than `temp`, so an accumulation of pending elements will happen. We do this while loop because there is the possibility that `temp` is larger than more than one pending element in the stack.
3. We add the index of the current temperature to the stack so that it can be processed by the while loop in the following iterations. 

### **Why storing the indices and not the elements?**
1. There may be repeated elements
2. What we want is to solve the question "how far away is X from Y?". This question can only be answered by keeping track of positions, i.e. indexes. 
3. It's much, much easier to get an element from it's position than a position from it's element. 
