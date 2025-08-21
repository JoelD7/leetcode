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
Associate a global, incremental value `incr` with the "use" of a key. A key is used when you add or obtain it from the cache via `put` or `get`, respectively. Every time a key is used, it's associated `incr` value will increase. In this manner, the key with the highest `incr` will be the one most recently used, and the one with the lowest `incr` will be the LRU. 

## Main actors
### `cap` 
Capacity of the cache or the maximum amount of elements the cache can hold.

### `length`
Quantity of items currently on the cache.

### `incr`
A global, incremental value that is used to track the recency of use for a specific key. The highest `incr` indicates the most recent key; conversely, the lowest `incr` indicates the least recently used key(LRU).

### `lowestIncr`
Tracks the lowest `incr` found so far. This value will increase by 1 once we have to evict a cache entry. Since every time **any** key is used `incr` increases by one, we can be certain that there is a key that has an `incr` associated whose value is `lowestIncr + 1`.

### `cache`
The cache were values are stored.

### `keyByIncr`
Associates a particular key to an `incr` value -> `incr:key`. The purpose of this map is solely so we can know what's the key with the lowest `incr` so we can evict it. We'll use the `lowestIncr` field to do that.

### `incrByKey`
Associates an `incr` with a key -> `key:incr`. 

The purpose of this map is to help us keep `keyByIncr` updated with valid entries, which in turn will lead to less iterations when we are looking for the LRU key(more on that later). Every time an existing key is used again -added to the cache with a different value or gotten from the cache multiple times- its `incr` changes. This means that any other `incr`(s) that might have been associated with that key in the "past" are no longer useful. Remember, `incr` measures recency, so if a key "4" was the first one used in a cache with 500 keys, never used again for 100 operations and then used on the 101th operation, it means it's the most recent one. In other words, past `incr`s don't matter. 


## Implementation
### `Get(key int)`
Simple, get the value from the cache, return it if it exists or return -1 if it doesn't:
```go
func (this *LRUCache) Get(key int) int {
	val, ok := this.cache[key]
	if !ok {
		return -1
	}
	//...
	return val
}
```
---
If there is an `incr` associated with this key(`ok = true`), it means this is not the first time the key is being accessed, so we check if that `incr` is lower than `lowestIncr` to update it, and delete the found `incr` from `keyByIncr` because `key` will have a new `incr` so there is no point in keeping the old one. 

`lowestIncr` is updated to the `incr` that is right next to `keyIncr` because `keyIncr` will be deleted, and since the list of `incrs` is incremental, it means it's sorted, so the next lowest incr should obviously be `keyIncr + 1`. 
```go
keyIncr, ok := this.incrByKey[key]
if ok {
    if keyIncr < this.lowestIncr {
        this.lowestIncr = keyIncr + 1
    }
    delete(this.keyByIncr, keyIncr)
}
```
---
Next, we update the global `incr` and each "incr-tracker" map. 
```go
this.incr++
this.keyByIncr[this.incr] = key
this.incrByKey[key] = this.incr
```
### `Put(key int, value int)`
If there is a cache entry associated with the key, we update it(overwriting is allowed), increase the `incr` and update the "incr-tracker" maps accordingly. But first, we have to delete the old `incr` associated with this key from `keyByIncr` for the reasons we have stated before.
```go
_, ok := this.cache[key]
 if ok {
    this.cache[key] = value
    this.incr++

    cur, _ := this.incrByKey[key]
    delete(this.keyByIncr, cur)

    this.keyByIncr[this.incr] = key
    this.incrByKey[key] = this.incr
    return
}
```
---
If the cache is already full, we have to evict the LRU item from it. We use `lowestIncr` to find the key with the lowest incr from the `keyByIncr` map, and then be able to delete the entry for that key from the cache.

We need to use a loop here because `lowestIncr` might be "pointing" to an `incr` that is no longer valid, i.e., the key associated with that `incr` has been updated with new `incr`s. There is no way(that I have found) to **always** keep `lowestIncr` with the real lowest incr without doing some sort of iteration every time we either add or get some element from the cache. And that is definitely worst from a performance standpoint than doing it here.  
```go
if this.length == this.cap {
    ok = false
    var k int
    for !ok {
        k, ok = this.keyByIncr[this.lowestIncr]
        this.lowestIncr++
    }
    delete(this.cache, k)
}
```
---
Next, insert the new key in the cache, update `incr` and update the "incr-tracker" maps accordingly. 

Here we set `lowestIncr` to the new `incr` in the case there hasn't been any `get` operations yet(the other part of the program were `lowestIncr` is updated).

`length` only increases when it's less than `cap`. When `length` reaches the same value as `cap` it will **never** be lower as we have no "delete" method. The only way to delete an item is by cache eviction and that is already handled above.
```go
this.cache[key] = value
this.incr++
this.keyByIncr[this.incr] = key
if this.incr < this.lowestIncr {
    this.lowestIncr = this.incr
}

this.incrByKey[key] = this.incr
if this.length < this.cap {
    this.length++
}

```





