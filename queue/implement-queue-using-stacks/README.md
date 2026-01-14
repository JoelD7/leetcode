# Problem
https://leetcode.com/problems/implement-queue-using-stacks/description/

Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

Implement the `MyQueue` class:

    void push(int x) Pushes element x to the back of the queue.
    int pop() Removes the element from the front of the queue and returns it.
    int peek() Returns the element at the front of the queue.
    boolean empty() Returns true if the queue is empty, false otherwise.

### Notes:

    You must use only standard operations of a stack, which means only push to top, peek/pop from top, size, and is empty operations are valid.
    Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.


### Example 1:

    Input
    ["MyQueue", "push", "push", "peek", "pop", "empty"]
    [[], [1], [2], [], [], []]
    Output
    [null, null, null, 1, 1, false]
    
    Explanation
    MyQueue myQueue = new MyQueue();
    myQueue.push(1); // queue is: [1]
    myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
    myQueue.peek(); // return 1
    myQueue.pop(); // return 1, queue is [2]
    myQueue.empty(); // return false


### Constraints:

    1 <= x <= 9
    At most 100 calls will be made to push, pop, peek, and empty.
    All the calls to pop and peek are valid.



**Follow-up**: Can you implement the queue such that each operation is amortized O(1) time complexity? In other words, performing n operations will take overall O(n) time even if one of those operations may take longer.

# Solution
### Intuition

Use two stacks. The solution of this problem stems from the fact that when you pop elements out of a stack X and inmediately push them to a stack Y, stack Y will contain the elements of X in reverse order. The reverse of stack is basically a queue because now the top of stack is the first pushed element.

### Variables/attributes

- `s1`: stack one. This is the primary stack and it will always have the elements in the order of the queue because we’ll keep it that way.
- `s2`: stack two. This the auxiliary stack we’ll use to reverse the primary

### Algorithm

**For push**

The idea here is using the auxiliary stack `s2` to maintain the primary stack `s1` in the same order of the queue, this is, the top of stack `s1` will always be the head of the queue.

1. On every push operation, pop all the elements out of `s1` and push them into `s2`.
    1. So if `s1` had `1 → 2 → 3 → 4` ,then `s2` will have `4 → 3 → 2 → 1`
    2. Push the new element(lets say 5) into `s2` so now it is `5 → 4 → 3 → 2 → 1`.
    3. Push all the elements of `s2` into `s1` so now it will be `1 → 2 → 3 → 4 → 5`

**For pop**

Since `s1` already has the elements in the right order, we need just to call the `pop` function of the stack.
