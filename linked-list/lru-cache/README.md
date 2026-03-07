# Problem
https://leetcode.com/problems/lru-cache/description/


Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the `LRUCache` class:

- `LRUCache(int capacity)` Initialize the LRU cache with positive size capacity.
- `int get(int key)` Return the value of the key if the key exists, otherwise return -1.
- `void put(int key, int value)` Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.

The functions get and put must each run in O(1) average time complexity.


### Example 1:

    Input
    ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
    [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
    Output
    [null, null, null, 1, null, -1, null, -1, 3, 4]

    Explanation
    LRUCache lRUCache = new LRUCache(2);
    lRUCache.put(1, 1); // cache is {1=1}
    lRUCache.put(2, 2); // cache is {1=1, 2=2}
    lRUCache.get(1);    // return 1
    lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
    lRUCache.get(2);    // returns -1 (not found)
    lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
    lRUCache.get(1);    // return -1 (not found)
    lRUCache.get(3);    // return 3
    lRUCache.get(4);    // return 4


### Constraints:

    1 <= capacity <= 3000
    0 <= key <= 104
    0 <= value <= 105
    At most 2 * 105 calls will be made to get and put.

# Solution
## Variables/attributes

- `type *Node`: node of the linked list. Holds the key and value of a cache entry along with pointers to next and previous nodes.
- `head  *Node`: head of the linked list. Every “use”(get or put operation) will be represented as a node placed on the head.
- `tail  *Node`: tail of the linked list. Since the head will always have the latests uses, consequently the tail will have the LRU key. This is where we’ll delete keys from.
- `table map[int]*Node`: the actual cache. Maps keys to nodes.

## TL;DR

We use a doubly linked list to maintain a sequence of cache entry usages. A cache entry is used whenever we add or get an entry from the cache. Every time we use a cache entry, we *move* it to the head of the list, which means that the LRU entry/node will always be at the tail of the list. When it comes time to evict a cache entry, all we have to do is removing the tail.

**Why a doubly instead of singly linked list?** Easy and constant removal. Removing and moving a node of a doubly linked list is faster than doing it for a singly linked list, because we’re able to quickly reference the node that is ahead and before the current one, so we can then manipulate the pointers of those two nodes to “ignore” the node we’re deleting.

## Implementation

### Note on dummy nodes

First of all is essential to note that this implementation uses **dummy nodes** both in the head and tail nodes which considerably simplifies our code. Is important to keep this in mind at all times because to avoid confusion:

```go
**this.head.next** = the actual head of the list
**this.tail.prev** = the actual tail of the list
```

---

### Node manipulation functions

Remember we’re using dummy nodes, so `this.head` is a dummy. The actual head is `this.head.next`.

This function is straigthforward. We manipulate the prev and next pointers to “place” `node` in between the dummy head and the previous head.

```go
func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node
}
```

---

Manipulate the pointers of the previous and next nodes of `node` to ignore it.

```go
func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
```

---

A combination of the two previous functions. We use this function to denote a new use of a cache entry.

 > *Cache entry usage = moving a node to head*

Note that we “**move**” instead of plainly adding it. Moving it’s better because it saves us the pain of maintaining duplicate nodes, and since we’re using a doubly linked list the moving part is trivial and done quickly.

```go
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}
```

The `removeTail` function is straigthforward.

### `Get` function

Get the node from the `table` and return it’s value. We move that node to head to denote it has been used.

### `Put` function

If the key exists we update the node’s value and move it to head to denote a use. Otherwise, we create a new node, add it to head(no need to move it because it doesn’t exists) and add it to the table(cache).

If the size exceeded the capacity, remove the tail and remove it from the table. Remember, because of the “usage-adding-to-head-logic”, the node in tail will always be the LRU.