# Problem

https://leetcode.com/problems/insert-delete-getrandom-o1/description/

Implement the RandomizedSet class:

- `RandomizedSet()` Initializes the RandomizedSet object.
- `bool insert(int val)` Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
- `bool remove(int val)` Removes an item val from the set if present. Returns true if the item was present, false otherwise.
- `int getRandom()` Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.

You must implement the functions of the class such that each function works in average O(1) time complexity.



### Example 1:

    Input
    ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
    [[], [1], [2], [2], [], [1], [2], []]
    Output
    [null, true, false, true, 2, true, false, 2]
    
    Explanation
    RandomizedSet randomizedSet = new RandomizedSet();
    randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
    randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
    randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
    randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
    randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
    randomizedSet.insert(2); // 2 was already in the set, so return false.
    randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.


### Constraints:

    -231 <= val <= 231 - 1
    At most 2 * 105 calls will be made to insert, remove, and getRandom.
    There will be at least one element in the data structure when getRandom is called.

# Solution
Use two maps, one to store the values of the set(main map) and one to store the indices of said values(index map). The keys of the main map will emulate indices of an array with the purpose of simplifying the process of getting a random element.

## Main actors
The `RandomizedSet` struct has the following fields:
- `m`: The main map that keeps the items of the set as its values. The key of `m` is the incremental index that emulates an array index: `i -> value`.
- `i`: The incremental index associated with each `m` entry. It will allow us to get a random value from the set in `O(1)`, by using a random function between `0` and `i`.
- `indexByVal`: The opposite of `m`. Holds the indices in which the items of the set are stored on `m`. This map allow us to know the indices that are empty(this happens after a removal) or not. The existence of "free spots" on `m` is very important because we can **reutilize** the spaces of `m`. Reutilizing slots translates to lowering the amount of times we have to regenerate random indices that don't have values. The more times we regenerate a random index, the farther we get from `O(1)` time for the `GetRandom()` function. 
- `lastDelIdx`: Holds the index to the last deleted value of `m`. In other words, indicates a free spot of the map, a place where the next inserted item will be located. The `lastDelIdx` along with `indexByVal` is what allows the reutilization of spaces in `m`.


## Implementation
### Insert()
Whether there is a free spot in `m` or not is what dictates the insertion logic. A positive value of `lastDelIdx` indicates that there is, so we insert the new value in that position and update `indexByVal` accordingly.
We then set `lastDelIdx` to a negative value to indicate there are no free spots available anymore. 
```go
   ...
    if this.lastDelIdx >= 0 {
        this.m[this.lastDelIdx] = val
        this.indexByVal[val] = this.lastDelIdx
        this.lastDelIdx = -1
    } ...
    ...
```
---
If there are no free spots, we "create" a new one by inserting the new item on `i`, updating `indexByVal` accordingly and incrementing `i` to create a new spot if needed in the future.
```go
    ... else {
        this.m[this.i] = val
        this.indexByVal[val] = this.i
        this.i++
    }
    return true
```
### Remove()
We first get the index of the item we want to delete, then delete the item from both maps. We need said index to update `lastDelIdx` and to know where we'll insert the new item in order to reutilize the deleted slot.
```go
func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.indexByVal[val]
	if ok {
		delete(this.m, i)
		this.lastDelIdx = i
		delete(this.indexByVal, val)
		return true
	}

	return false
}
```
### GetRandom()
This is where the magic happens. This is why we emulate the indices of an array to store the values of the set: to easily get a random number. By just getting a random number between `0` and `i`, we can indirectly get a random number from `m` in `O(1)` time. However, there is the possibility that the random number corresponds to a index that has no value, i.e. a non-existing key in `m`, in which case the function would return 0. To solve that we generate indices until we find one that exists in `m`.

The reason why we need to reutilize spots on `m` is because of this loop, to avoid, as much as possible, generating non-existing indices that would increase the amount of times we have to iterate. If we only use the incrementing value of `i` to insert items, and have 1,500 insertions and 1,498 removals, we would end up with `i = 1500` and a random-number generating function that would be wrong 1498 out of 1500 times... Yikes! On the other hand, by using the implemented approach we would only have to look **2 times** at the maximum!
```go
func (this *RandomizedSet) GetRandom() int {
    var val, index int
    var ok bool

    for !ok {
        index = rand.Intn(this.i)
        val, ok = this.m[index]
    }

    return val
}
```







