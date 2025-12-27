# Problem
https://leetcode.com/problems/min-stack/

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

    MinStack() initializes the stack object.
    void push(int val) pushes the element val onto the stack.
    void pop() removes the element on the top of the stack.
    int top() gets the top element of the stack.
    int getMin() retrieves the minimum element in the stack.

You must implement a solution with O(1) time complexity for each function.



Example 1:

    Input
    ["MinStack","push","push","push","getMin","pop","top","getMin"]
    [[],[-2],[0],[-3],[],[],[],[]]

    Output
    [null,null,null,null,-3,null,0,-2]

    Explanation
    MinStack minStack = new MinStack();
    minStack.push(-2);
    minStack.push(0);
    minStack.push(-3);
    minStack.getMin(); // return -3
    minStack.pop();
    minStack.top();    // return 0
    minStack.getMin(); // return -2

# Solution
From analyzing the problem we reach the conclusion that some sort of history of minimum elements must be created so that when minimum elements are popped out of the stack we can easily reference the “previous” minimum in `O(1)` time.

The solution is **encoding the minimum element history on each node**. Each node will keep track of the minimum element on the list *at the time* said node was added. Every time we add a new node to the stack, that node will save the minimum element *so far*.

By making each node hold the minimum element at the time the node was added to the list, when an element is popped one out of two things can happen:

1. If the top element was the minimum, both the minimum and the head are removed from the stack, and the new minimum is now the one that was before the minimum we just eliminated, which is already saved on the node after the removed head
2. If the top element was not the minimum, the head is popped and the minimum remains the same because in that case the popped head and the node before would both have the same minimum, which is what we want.