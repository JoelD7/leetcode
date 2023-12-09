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
- Use a regular stack(called `stack` from now on) that holds another stack(`minStack`)
as one of it's fields. 
  - `minStack` is an auxiliary stack that will hold the minimum values
  of `stack` in ascending order, where `minStack.top()` will be the `stack`'s current minimum
  value.
- On each `push` to `stack`, check if the new value is smaller than
`minStack`'s top value. If it is, push to `minStack`, otherwise, do nothing. 
- On each `pop`, check if `minStack`'s top value is the same as the `stack`'s top value. If
it is, `pop` from `minStack` as well. 
  - It doesn't matter if the min value is repeated because: 
    1. we would have already added it to`minStack` in a previous `push`
    2. the values of `minStack` are sorted, so `minStack` will always have the minimum
    value

## Why use an auxiliary stack?
First of all, the problem description states that all operations should have an `O(1)`
complexity. A naive approach might involve using an array to hold all the added values. However, this is a problem because after poping an item from `stack` we would have to update that array to eliminate the
popped value, and also update the `min` value in case the poped value was the `min`. This is a `O(n)` operation.

If we substitute that array with an embedded stack that contains all the minimum values `stack` has had in ascending order, the current minimum value will always be at the top. It's removal will be `O(1)` and there won't be any need
of "updating" the `stack`'s minimum value.