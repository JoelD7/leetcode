# Problem
https://leetcode.com/problems/design-circular-queue/description/


Design your implementation of the circular queue. The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle, and the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".

One of the benefits of the circular queue is that we can make use of the spaces in front of the queue. In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space in front of the queue. But using the circular queue, we can use the space to store new values.

Implement the `MyCircularQueue` class:

    MyCircularQueue(k) Initializes the object with the size of the queue to be k.
    int Front() Gets the front item from the queue. If the queue is empty, return -1.
    int Rear() Gets the last item from the queue. If the queue is empty, return -1.
    boolean enQueue(int value) Inserts an element into the circular queue. Return true if the operation is successful.
    boolean deQueue() Deletes an element from the circular queue. Return true if the operation is successful.
    boolean isEmpty() Checks whether the circular queue is empty or not.
    boolean isFull() Checks whether the circular queue is full or not.

You must solve the problem without using the built-in queue data structure in your programming language.



### Example 1:

    Input
    ["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]
    [[3], [1], [2], [3], [4], [], [], [], [4], []]
    Output
    [null, true, true, true, false, 3, true, true, true, 4]
    
    Explanation
    MyCircularQueue myCircularQueue = new MyCircularQueue(3);
    myCircularQueue.enQueue(1); // return True
    myCircularQueue.enQueue(2); // return True
    myCircularQueue.enQueue(3); // return True
    myCircularQueue.enQueue(4); // return False
    myCircularQueue.Rear();     // return 3
    myCircularQueue.isFull();   // return True
    myCircularQueue.deQueue();  // return True
    myCircularQueue.enQueue(4); // return True
    myCircularQueue.Rear();     // return 4



### Constraints:

    1 <= k <= 1000
    0 <= value <= 1000
    At most 3000 calls will be made to enQueue, deQueue, Front, Rear, isEmpty, and isFull.


# Solution
## Intuition
Use an array as the underlying data structure that holds the elements and use indeces that indicate the beginning and end of the queue. 

## Attributes
- `cap`: capacity of the queue. Indicates the max amount of elements it can hold
- `len`: length of the queue. Indicates the amount of elements the queue currently has
- `head`: index of the underlying array that points to the beginning of the queue
- `tail`: index of the underlying array that points to the end of the queue
- `arr`: the underlying array holding the elements of the queue

## Algorithm
We first initialize the key attributes of the circular queue. The `tail` should be initialized as -1 to make the code of the `Enqueue` function simpler. If we didn't initialized as -1, we would have to include a few lines of code just to handle the first enqueued item.
```go
func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        cap:  k,
        arr:  make([]int, k),
        tail: -1,
    }
}

```
---
The key thing to note here are the `tail` increment logic. First, we increment `tail` **before** enqueuing the `value` because if we did it after, then the `tail` index would potentially point to an index that doesn't have a item. Secondly, since this is **circular** queue it means the `tail` should go back to the beginning once it has reached the final index of the allocated capacity. We can safely set `tail` to 0 here because we already checked that the queue wasn't full so it means that index 0 is vacant.   

```go
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.tail == this.cap-1 {
		this.tail = 0
	} else {
		this.tail++
	}

	this.arr[this.tail] = value
	this.len++

	return true
}
```
---
This function is more straightforward. We apply the same "circling back" logic to the `head` pointer
```go
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	//This means that the "enqueueing" has "circled back"
	if this.head == this.cap-1 {
		this.head = 0
	} else {
		this.head++
	}

	this.len--

	return true
}
```

The `Enqueue` and `Dequeue` functions contain the logic that makes a queue *circular*. The other functions are self-explanatory. 